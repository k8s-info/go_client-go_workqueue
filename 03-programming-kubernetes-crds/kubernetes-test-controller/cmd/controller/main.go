package main

import "k8s.io/klog"

func main() {
	controller := NewController()
	controller.CreateCRD()
	if err := controller.CreateObject(); err != nil {
		klog.Info(err)
	}

	controller.Run()
}
