/*
	Student对象的内容已经被设定好，
    主要有name和school这两个字段，
	表示学生的名字和所在学校，
	因此创建Student对象的时候内容就要和这里匹配了；
*/

package v1

import (
	"go/ast"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/routime.Object

type Student struct {
	metav1.TypeMeta     `json:",inline"`
	metav1.ObjectMeta   `json:"metadata,omitempty"`
	Spec            StudentSpec `json:"spec"`
}

type StudentSpec struct {
	name string `json:"name"`
	school string `json:"school"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StudentList is a list of Student resources
type StudentList struct {
	metav1.TypeMeta  `json:",inline"`
	metav1.ListMeta  `json:"metadata"`

	Items []Student  `json:"items"`
}
