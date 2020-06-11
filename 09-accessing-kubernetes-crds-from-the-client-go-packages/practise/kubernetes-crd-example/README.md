[Accessing Kubernetes CRDs from the client-go package](https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html)

[kubernetes-crd-example](https://github.com/martin-helmich/kubernetes-crd-example)

done:
```bash
go/bin/go build -o /private/var/folders/k8/t39hjy910bx2gh5k1ztk7fbc0000gn/T/___go_build_github_com_ycliu912_kubernetes_crd_example github.com/ycliu912/kubernetes-crd-example #gosetup
/private/var/folders/k8/t39hjy910bx2gh5k1ztk7fbc0000gn/T/___go_build_github_com_ycliu912_kubernetes_crd_example #gosetup
projects found: &{TypeMeta:{Kind: APIVersion:} ListMeta:{SelfLink:/apis/example.ycliu912.com/v1alpha1/namespaces/default/projects ResourceVersion:560221 Continue: RemainingItemCount:<nil>} Items:[{TypeMeta:{Kind:Project APIVersion:example.ycliu912.com/v1alpha1} ObjectMeta:{Name:example-project GenerateName: Namespace:default SelfLink:/apis/example.ycliu912.com/v1alpha1/namespaces/default/projects/example-project UID:1ddb269f-8edf-40fe-8184-c13ae7eeea14 ResourceVersion:524148 Generation:1 CreationTimestamp:2020-06-09 11:05:56 +0800 CST DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"example.ycliu912.com/v1alpha1","kind":"Project","metadata":{"annotations":{},"name":"example-project","namespace":"default"},"spec":{"replicas":1}}
] OwnerReferences:[] Finalizers:[] ClusterName: ManagedFields:[]} Spec:{Replicas:1}}]}
project in store: 0
project in store: 0
project in store: 0
```