package main

import (
	"os"
	"time"

	corev1 "k8s.io/api/core/v1"
	apiextenssionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"

	testresourceclienteset "ycliu912.github.io/kubernetes-test-controller/lib/testresource/generated/clientset/versioned"
	testresourcescheme "ycliu912.github.io/kubernetes-test-controller/lib/testresource/generated/clientset/versioned/scheme"
	testresorceinformers "ycliu912.github.io/kubernetes-test-controller/lib/testresource/generated/informers/externalversions"
	testresoucelisers   "ycliu912.github.io/kubernetes-test-controller/lib/testresource/generated/listers/testresource/v1beta1"
	testresourcev1beta1 "ycliu912.github.io/kubernetes-test-controller/lib/testresource/v1beta1"
)

type Controller struct {
	kubeclientset   kubernetes.Interface
	apiextenssionsclientset  apiextenssionsclientset.Interface
	testresourceclientset  testresourceclienteset.Interface
	informer    cache.SharedIndexInformer
	lister  testresoucelisers.TestResourceLister
	recorder record.EventRecorder
	workqueue   workqueue.RateLimitingInterface
}

func NewController() *Controller {
	kubeconfig := os.Getenv("KUBECONFIG")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		klog.Fatal(err.Error())
	}

	kubeClient := kubernetes.NewForConfigOrDie(config)
	apiextensionsClient := apiextenssionsclientset.NewForConfigOrDie(config)
	testClient := testresourceclienteset.NewForConfigOrDie(config)

	informerFactory := testresorceinformers.NewSharedInformerFactory(testClient, time.Minute * 1)
	informer := informerFactory.Ycliu912().V1beta1().TestResources()
	informer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(object interface{}) {
				klog.Info("Added: %v", object)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				klog.Infof("Updated: %v", newObj)
			},
			DeleteFunc: func(object interface{}) {
				klog.Info("Deleted: %v", object)
			},
		})
	informerFactory.Start(wait.NeverStop)

	utilruntime.Must(testresourcev1beta1.AddToScheme(testresourcescheme.Scheme))
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(testresourcescheme.Scheme, corev1.EventSource{Component: "testresource-controller"})

	workqueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	return &Controller{
		kubeclientset: kubeClient,
		apiextenssionsclientset: apiextensionsClient,
		testresourceclientset: testClient,
		informer: informer.Informer(),
		lister: informer.Lister(),
		recorder: recorder,
		workqueue: workqueue,
	}
}

func (c *Controller) Run() {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	klog.Info("Waiting cache to be synced.")
	// Handle timeout for syncing.
	timeout := time.NewTicker(time.Second * 30)
	timeoutCh := make(chan struct{})

	go func() {
		<-timeout.C
		timeoutCh <- struct{}{}
	}()
	if ok := cache.WaitForCacheSync(timeoutCh, c.informer.HasSynced); !ok {
		klog.Fatal("Timeout expired during waiting for caches to sync.")
	}

	klog.Infoln("Starting custom controller.")
	select {
	}
}

func main() {
	Controller := NewController()
	Controller.Run()
}