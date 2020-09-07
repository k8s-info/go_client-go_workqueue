module github.com/k8s-info/go_client-go_workqueue

go 1.13

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.1 // indirect
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200506145744-7e3656a0809f // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.0.0-20200509044756-6aff5f38e54f // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/sample-controller v0.18.2
	k8s.io/utils v0.0.0-20200414100711-2df71ebbae66 // indirect
)

replace (
	k8s.io/api v0.18.2 => k8s.io/api v0.17.4
	k8s.io/apimachinery v0.18.2 => k8s.io/apimachinery v0.17.4
	k8s.io/client-go v11.0.0+incompatible => k8s.io/client-go v0.17.0
)
