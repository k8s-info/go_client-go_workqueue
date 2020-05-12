module github.com/k8s-info/go_client-go_workqueue

go 1.13

require (
	github.com/imdario/mergo v0.3.9 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20200414100711-2df71ebbae66 // indirect
)

replace (
	k8s.io/api v0.18.2 => k8s.io/api v0.17.4
	k8s.io/apimachinery v0.18.2 => k8s.io/apimachinery v0.17.4
	k8s.io/client-go v11.0.0+incompatible => k8s.io/client-go v0.17.0
)
