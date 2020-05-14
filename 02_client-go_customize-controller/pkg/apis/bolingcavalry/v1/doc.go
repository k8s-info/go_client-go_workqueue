/*
	都是代码生成工具会用到的，
	一个是声明为整个v1包下的类型定义生成DeepCopy方法，
	另一个声明了这个包对应的API的组名，和CRD中的组名一致；
*/

// +k8s:deepcopy-gen=package

// +groupName=bolingcavalry.k8s.io
package v1

