[Programming Kubernetes CRDs](https://insujang.github.io/2020-02-13/programming-kubernetes-crd/)

```bash
~/go_client-go_workqueue/03-programming-kubernetes-crds/code-generator$ ./generate-groups.sh all ycliu912.github.io/kubernetes-test-controller/lib/testresource/generated ycliu912.github.io/kubernetes-test-controller/lib testresource:v1beta1 --go-header-file ./hack/boilerplate.go.txt --output-base ..
Generating deepcopy funcs
Generating clientset for testresource:v1beta1 at ycliu912.github.io/kubernetes-test-controller/lib/testresource/generated/clientset
Generating listers for testresource:v1beta1 at ycliu912.github.io/kubernetes-test-controller/lib/testresource/generated/listers
Generating informers for testresource:v1beta1 at ycliu912.github.io/kubernetes-test-controller/lib/testresource/generated/informers
liuyanchao@ycliu912-mac:~/go_client-go_workqueue/03-programming-kubernetes-crds/code-generator$ cd ..
(reverse-i-search)`cp': cp -r ./ycliu912.github.io/kubernetes-test-controller/lib/testresource/* kubernetes-test-controller/lib/testresource/
```

需要注意的地方：

留意导入的包，稍一不小心就会导入无用的包，编译报错

代码编写错误


go run cmd/controller/main.go
报错：
```bash
# ycliu912.github.io/kubernetes-test-controller/lib/testresource/v1beta1
../../lib/testresource/v1beta1/register.go:29:3: cannot use &TestResourceList literal (type *TestResourceList) as type runtime.Object in argument to scheme.AddKnownTypes:
	*TestResourceList does not implement runtime.Object (missing DeepCopyObject method)
note: module requires Go 1.14

done：
// +k8s:deepcopy-gen:interface=k8s.io/apimachinery/pkg/runtime.Object 
=>
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Objec
```

---
```bash
1.13.1/bin/go build -o /private/var/folders/k8/t39hjy910bx2gh5k1ztk7fbc0000gn/T/___go_build_main_go /Users/liuyanchao/go_client-go_workqueue/03-programming-kubernetes-crds/kubernetes-test-controller/cmd/controller/main.go #gosetup
/private/var/folders/k8/t39hjy910bx2gh5k1ztk7fbc0000gn/T/___go_build_main_go #gosetup
I0515 18:20:45.254982   27264 main.go:98] Waiting cache to be synced.
E0515 18:20:45.282751   27264 reflector.go:156] pkg/mod/k8s.io/client-go@v0.17.0/tools/cache/reflector.go:108: Failed to list *v1beta1.TestResource: the server could not find the requested resource (get testresources.ycliu912.github.io)

```

