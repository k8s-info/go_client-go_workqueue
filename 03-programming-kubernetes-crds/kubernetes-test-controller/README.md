需要注意的地方：

留意导入的包，少一不小心就会导入无用的包，编译报错

代码编写错误


go run cmd/controller/main.go
报错：
```bash
# ycliu912.github.io/kubernetes-test-controller/lib/testresource/v1beta1
../../lib/testresource/v1beta1/register.go:29:3: cannot use &TestResourceList literal (type *TestResourceList) as type runtime.Object in argument to scheme.AddKnownTypes:
	*TestResourceList does not implement runtime.Object (missing DeepCopyObject method)
note: module requires Go 1.14
```