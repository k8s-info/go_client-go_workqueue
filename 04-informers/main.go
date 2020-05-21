/*
	Extend Kubernetes via a Shared Informer
	https://gianarb.it/blog/kubernetes-shared-informer
*/

package main

import (
	"flag"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"os"
	"path/filepath"
	"time"
)

func main() {

	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	// Create the client configuration
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		klog.Fatal(err.Error())
	}

	// Create the client
	clientset := kubernetes.NewForConfigOrDie(config)

	// Create the shared informer factory and use the client to connect to Kubernetes
	factory := informers.NewSharedInformerFactory(clientset, time.Minute * 1)

	// Get the informer for the right resource, in this case a Pod
	informer := factory.Core().V1().Pods().Informer()

	// Create a channel to stops the shared informer gracefully
	stopper := make(chan struct{})
	defer close(stopper)

	// Kubernetes serves an utility to handle API crashes
	defer runtime.HandleCrash()

	// This is the part where you custom code gets triggered based on the
	// event that the shared informer catches
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		// When a new pod gets created
		//AddFunc:    func(obj interface{}) { panic("not implemented") },
		AddFunc: func(obj interface{}) {
			// Cast the obj as Pod
			pod := obj.(*corev1.Pod)
			klog.Info("Added: ", pod)
		},

		// When a pod gets updated
		//UpdateFunc: func(NewObj interface{}, OldObj interface{}) { panic("net implemented") },
		UpdateFunc: func(oldObj, newObj interface{}) {
			newPod := newObj.(*corev1.Pod)
			klog.Info("Update: ", newPod)
		},

		// When a pod get deleted
		// DeleteFunc: func(obj interface{}) { panic("not implemented") },
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*corev1.Pod)
			klog.Info("Delete: ", pod)
		},
	})

	// You need to start the informer, in my case, it runs in the background
	go informer.Run(stopper)

	select {
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows
}