- [client-go workqueue demo](https://www.cnblogs.com/wangjq19920210/p/11551825.html)


```
go build -o /private/var/folders/k8/t39hjy910bx2gh5k1ztk7fbc0000gn/T/___go_build_github_com_k8s_info_go_client_go_workqueue github.com/k8s-info/go_client-go_workqueue #gosetup
/private/var/folders/k8/t39hjy910bx2gh5k1ztk7fbc0000gn/T/___go_build_github_com_k8s_info_go_client_go_workqueue #gosetup
W0511 17:55:08.228701   40019 client_config.go:543] Neither --kubeconfig nor --master was specified.  Using the inClusterConfig.  This might not work.
W0511 17:55:08.228954   40019 client_config.go:548] error creating inClusterConfig, falling back to default config: unable to load in-cluster configuration, KUBERNETES_SERVICE_HOST and KUBERNETES_SERVICE_PORT must be defined
F0511 17:55:08.229831   40019 main.go:126] 构建kubeconfig失败： invalid configuration: no configuration has been provided
goroutine 1 [running]:
github.com/golang/glog.stacks(0xc0000ac200, 0xc000404000, 0x7f, 0xd3)
        /Users/liuyanchao/go/pkg/mod/github.com/golang/glog@v0.0.0-20160126235308-23def4e6c14b/glog.go:769 +0xb8
github.com/golang/glog.(*loggingT).output(0x2b55e80, 0xc000000003, 0xc000382d90, 0x2ac51cc, 0x7, 0x7e, 0x0)
        /Users/liuyanchao/go/pkg/mod/github.com/golang/glog@v0.0.0-20160126235308-23def4e6c14b/glog.go:720 +0x372
github.com/golang/glog.(*loggingT).printf(0x2b55e80, 0x3, 0x2012e8c, 0x1c, 0xc00010def8, 0x1, 0x1)
        /Users/liuyanchao/go/pkg/mod/github.com/golang/glog@v0.0.0-20160126235308-23def4e6c14b/glog.go:655 +0x14b
github.com/golang/glog.Fatalf(...)
        /Users/liuyanchao/go/pkg/mod/github.com/golang/glog@v0.0.0-20160126235308-23def4e6c14b/glog.go:1148
main.main()
        /Users/liuyanchao/go_client-go_workqueue/01/main.go:126 +0x879

Process finished with exit code 255
```