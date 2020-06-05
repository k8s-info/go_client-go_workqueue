[Building your own kubernetes CRDs](https://itnext.io/building-your-own-kubernetes-crds-701de1c9a161)

[kubebuilder](https://github.com/kubernetes-sigs/kubebuilder/releases)

[complete code](https://gitlab.com/pongsatt/githook) 


```bash
  568  cd gitbook-tutorial/
  569  ls
  570  ~/go_k8s_kubebuilder/kubebuilder_2.3.1_darwin_amd64/bin/kubebuilder init --domain my.domain
  571  ls
  573  ~/go_k8s_kubebuilder/kubebuilder_2.3.1_darwin_amd64/bin/kubebuilder create api --group tools --version v1alpha1 --kind GitHook
  574  tree
  575  ls
  576  go mod download -x
```

// TODO: 3. Implement controller

// error:
```bash
~/go_client-go_workqueue/08-kubebuilder/gitbook-tutorial$ make generate
/Users/liuyanchao/go/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
../../../go/pkg/mod/k8s.io/kube-openapi@v0.0.0-20200427153329-656914f816f9/pkg/util/proto/document.go:24:2: case-insensitive import collision: "github.com/googleapis/gnostic/openapiv2" and "github.com/googleapis/gnostic/OpenAPIv2"
Error: not all generators ran successfully
run `controller-gen object:headerFile=hack/boilerplate.go.txt paths=./... -w` to see all available markers, or `controller-gen object:headerFile=hack/boilerplate.go.txt paths=./... -h` for usage
make: *** [generate] Error 1
```