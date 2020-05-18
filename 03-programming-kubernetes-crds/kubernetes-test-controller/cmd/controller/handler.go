package main

import (
	"k8s.io/client-go/tools/cache"
	"time"

	corev1 "k8s.io/api/core/v1"
	utilrentime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog"
)

// Run first sync client-go cache by calling cache.WaitforCacheSync
// then block the main thread forever
// Event handler will run in another Goroutine,
// generated in c.NewControoler() function
func (c *Controller) Run() {
	defer utilrentime.HandleCrash()
	defer c.workqueue.ShutDown()

	klog.Infoln("waiting cache to be synced.")
	// Handle timeout for syncing.
	timeout := time.NewTimer(time.Second * 30)
	timeoutCh := make(chan struct{})

	go func() {
		<-timeout.C
		timeoutCh <- struct{}{}
	}()

	if ok := cache.WaitForCacheSync(timeoutCh, c.informer.HasSynced); !ok {
		klog.Fatalln("Timeout expired duiing waiting for caches to sync.")
	}

	klog.Infoln("Starting sustom controller.")

	select {}
}

const statusMessage = "HANDLED"

func (c *Controller) objectAddedCallBack(object interface{}) {
	klog.Info("Added: %v", object)
	// You can get Testresource type as this
	// resource := object.(*testresourcev1beta1.TestResource)

	// or this

	key, _ := cache.MetaNamespaceKeyFunc(object)
	namespace, name, _ := cache.SplitMetaNamespaceKey(key)
	resource, _ := c.lister.TestResources(namespace).Get(name)

	// If the object is in the desired state, end callback.
	if resource.Status == statusMessage {
		return
	}

	// If the object is not handled yet, handle it by modifying in status.
	copy := resource.DeepCopy()
	copy.Status = statusMessage
	_, err := c.testresourceclientset.Ycliu912V1beta1().TestResources(corev1.NamespaceDefault).Update(copy)
	if err != nil {
		klog.Error(err.Error())
		return
	}
	c.recorder.Event(copy, corev1.EventTypeNormal, "ObjectHandled", "Object is handled by custom controller.")
}
