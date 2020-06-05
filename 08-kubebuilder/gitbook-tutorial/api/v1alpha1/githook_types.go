/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	tektonv1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SecretValueFromSource represents the source of a secret value
type SecretValueFromSource struct {
	// The secret key to select from.
	secretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// +kubebuilder:validation:Enum=gitlab;github;gogs

// GitProvider providers name of git provider

type GitProvider string

var (
	// Gitlab gitlab.com compatible
	Gitlab GitProvider = "gitlab"

	// Github github.com compatible
	Github GitProvider = "github"

	//Gogs gogs compatible
	Gogs GitProvider = "gogs"
)

// +kubebuilder:validation:Enum=create;delete;fork;push;issues;issue_comment;pull_request;release
type gitEvent string


// GitHookSpec defines the desired state of GitHook
type GitHookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ServiceAccountName holds the name of the kubernetes service account
	// as which the underlying k8s resources should be run. If unspecified
	// this will default to the "default" service account for the namespace
	// in which the GitHook exists.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// ProjectUrl is the url of the git project for which we are interested
	// to receive events from.
	// Example:
	// https://gitlab.com/pongsatt/githook
	// +kubebuilder:validation:MinLength=1
	ProjectURL string `json:"projectUrl"`

	// GitProvider is the name of the git source in which we would like register webhook
	GitProvider GitProvider `json:"gitProvider"`

	// EventType is the type of event to receive from Gogs. These
	// correspond to supported events to the add project hook
	// +kubebuilder:validation:MinItems=1
	EventTypes []gitEvent `json:"eventType"`

	// AccessToken is the Kubernetes secret containing the Gogs
	// access token
	AccessToken SecretValueFromSource `json:"accessToken"`

	// SecretToken is the Kubernetes secret containing the Gogs
	// secret token
	SecretToken SecretValueFromSource `json:"secretToken"`

	// SslVerify if true configure webhook so the ssl verification is done when triggering the hook
	// +optional
	SslVerify bool `json:"sslverify,omitempty"`

	// Runspec is a tekton pipelinerun spec to be run when events triggered
	RunSpec tektonv1alpha1.PipelineRunSpec `json:"runspec"`

	// Foo is an example field of GitHook. Edit GitHook_types.go to remove/update
	// Foo string `json:"foo,omitempty"`
}

// GitHookStatus defines the observed state of GitHook
type GitHookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ID of the project hook registered with Gogs
	ID string `json:"id,omitempty"`
}

// +kubebuilder:object:root=true

// GitHook is the Schema for the githooks API
type GitHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitHookSpec   `json:"spec,omitempty"`
	Status GitHookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GitHookList contains a list of GitHook
type GitHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitHook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitHook{}, &GitHookList{})
}
