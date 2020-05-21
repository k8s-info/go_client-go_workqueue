module github.com/golang-info/go_client-go_workqueue

go 1.13

replace (
	k8s.io/api v0.18.2 => k8s.io/api v0.17.4
	k8s.io/client-go v11.0.0+incompatible => k8s.io/client-go v0.17.0
)

require k8s.io/client-go v11.0.0+incompatible

require (
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.17.4
	k8s.io/klog v1.0.0
)
