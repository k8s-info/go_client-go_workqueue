package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/sample-controller/pkg/signals"
	"time"
)

/* 控制器 */
type Controller struct {
	// 此控制器使用的客户端
	clientset kubernetes.Interface
	// 此控制器使用的工作队列
	queue workqueue.RateLimitingInterface
	// 此控制器使用的共享Informer， SharedIndexInformer可以维护缓存中对象的索引
	informer cache.SharedIndexInformer
}

/* 主函数 */

var (
	// 参数变量
	masterURL string
	kubeconfig string
)

// 启动控制器
func (c *Controller) Run(stopCh <-chan struct{}) {
	// 捕获应用程序崩溃并打印日志
	defer utilruntime.HandleCrash()
	// 关闭队列， 从而导致 Worker 结束
	defer c.queue.ShutDown()

	glog.Info("启动控制器...")

	// 运行 Informer
	go c.informer.Run(stopCh)

	// 在启动 Worker 之前， 等待缓存同步完成
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		utilruntime.HandleError(fmt.Errorf("同步缓存超时"))
		return
	}

	glog.Info("缓存已经同步， 准别启动Worker")

	// 循环执行Worker， 直到TERM
	wait.Until(c.runWorker, time.Second, stopCh)
}

// 启动Worker
func (c *Controller) runWorker() {
	for c.processNextItem() {

	}
}

// Worker 的逻辑框架
func (c *Controller) processNextItem() bool {
	// 最大重试次数
	maxRetries := 3

	// 获取下一个元素， 第二个出参提示队列是否已经关闭
	key, quit := c.queue.Get()
	if quit {
		return false
	}

	// 总是移除 Key
	defer c.queue.Done(key)

	//
	err := c.processItem(key.(string))

	if err == nil {
		// 处理成功，提示队列不再跟踪事件历史
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < maxRetries {
		glog.Errorf("处理%s事件失败， 准别重试: %v", key, err)
		c.queue.AddRateLimited(key)
	} else {
		glog.Errorf("处理%s事件失败， 放弃： %v", key, err)
		c.queue.Forget(key)
		utilruntime.HandleError(err)
	}

	return true
}

// Worker 核心逻辑
func (c *Controller) processItem(key string) error {
	glog.Infof("开始处理事件%s", key)
	// 根据 Key 获取对象
	obj, exists, err := c.informer.GetIndexer().GetByKey(key)
	if err != nil {
		return fmt.Errorf("获取对象%s失败：%v", key, err)
	}
	fmt.Print(obj)
	if !exists {
		// 在这里处理对象删除事件
	} else {
		// 在这里处理对象创建事件
	}

	// 因为不进行 Resync，不会有更新事件
	return nil
}

func main() {
	// 解析参数，存入上述变量
	flag.Parse()
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("构建kubeconfig失败： %s", err.Error())
	}

	// 创建客户端，Clientset 是一系列 K8S API 的集合
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("构建 clientset 失败: %s", err.Error())
	}

	// 信号处理通道， 当进程接收到信号后， 此通道可读
	stopCh := signals.SetupSignalHandler()

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				// 仅仅列出所有命名空间的Pod
				return clientset.CoreV1().Pods(metav1.NamespaceAll).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return clientset.CoreV1().Pods(metav1.NamespaceAll).Watch(options)
			},
		},
		&apiv1.Pod{},
		0,              // 不进行relist
		cache.Indexers{}, // map[string]IndexFunc
	)

	// 添加事件处理回调，仅仅是简单的入队
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		// 此结构是实现 ResourceEventHandler
		AddFunc: func(obj interface{}) {
			// 从对象中抽取 Key
			key, err := cache.MetaNamespaceIndexFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	})

	// 构建控制器对象
	ctrl := Controller{
		clientset: clientset,
		queue: queue,
		informer: informer,
	}

	// 启动
	ctrl.Run(stopCh)
}