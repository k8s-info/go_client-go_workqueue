```bash
bash ./generate-groups.sh all ~/go_client-go_workqueue/02_client-go_customize-controller/pkg/client ~/go_client-go_workqueue/02_client-go_customize-controller/pkg/apis bolingcavalry:v1
```

- 查看etcd数据直接在k8s的etcd容器中使用etcdctl命令即可
```bash
docker exec -it $ETCD_CONTAINER_NAME sh
ETCDCTL_API=3 etcdctl --endpoints=https://[127.0.0.1]:2379 --cacert=/run/config/pki/etcd/ca.crt --cert=/run/config/pki/etcd/healthcheck-client.crt --key=/run/config/pki/etcd/hea
lthcheck-client.key get /registry/bolingcavalry.k8s.io/students/default/object-student --print-value-only
```

doing：
bug：
generate-groups.sh all
Generating deepcopy funcs
F0514 16:04:16.425623   12096 main.go:82] Error: Failed making a parser: unable to add directory "/Users/liuyanchao/go_client-go_workqueue/02_client-go_customize-controller/pkg/apis/bolingcavalry/v1": unable to import "/Users/liuyanchao/go_client-go_workqueue/02_client-go_customize-controller/pkg/apis/bolingcavalry/v1": import "/Users/liuyanchao/go_client-go_workqueue/02_client-go_customize-controller/pkg/apis/bolingcavalry/v1": cannot import absolute path

liuyanchao@ycliu912-mac:~/go/pkg/mod/k8s.io/code-generator@v0.18.2$ bash ./generate-groups.sh all ~/go_client-go_workqueue/02_client-go_customize-controller/pkg/client ~/go_client-go_workqueue/02_client-go_customize-controller/pkg/apis bolingcavalry:v1
Generating deepcopy funcs
F0514 16:08:21.454265   12305 main.go:82] Error: Failed making a parser: unable to add directory "/Users/liuyanchao/go_client-go_workqueue/02_client-go_customize-controller/pkg/apis/bolingcavalry/v1": unable to import "/Users/liuyanchao/go_client-go_workqueue/02_client-go_customize-controller/pkg/apis/bolingcavalry/v1": import "/Users/liuyanchao/go_client-go_workqueue/02_client-go_customize-controller/pkg/apis/bolingcavalry/v1": cannot import absolute path