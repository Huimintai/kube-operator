/*
Copyright 2023.

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
	api "github.com/Huimintai/kube-operator/pkg/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope="Cluster"
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=`.metadata.creationTimestamp`,name="Age",type=date

// ClusterOperationSpec defines the desired state of ClusterOperation
type ClusterOperationSpec struct {
	// Cluster is the name of Cluster
	// +required
	Cluster string `json:"cluster"`
	// HostsConfRef will be filled by operator when it performs backup
	// +optional
	HostsConfRef *api.ConfigMapRef `json:"hostsConfRef"`
	// VarsConfRef will be filled by operator when it performs backup
	// +optional
	VarsConfRef *api.ConfigMapRef `json:"varsConfRef,omitempty"`
	// SSHAuthRef will be filled by operator when it performs backup
	// +optional
	SSHAuthRef *api.SecretRef `json:"sshAuthRef,omitempty"`
	// EntrypointSHRef will be filled by operator when it renders entrypoint.sh
	// +optional
	EntrypointSHRef *api.ConfigMapRef `json:"entrypointSHRef,omitempty"`
	// Action is the kubespray yaml file to execute for example cluster.yml reset.yml
	// +required
	Action string `json:"action"`
	// ActionType is the file type to execute supported playbook and shell
	// +required
	ActionType ActionType `json:"actionType"`
	// +optional
	// +kubebuilder:default="builtin"
	ActionSource *ActionSource `json:"actionSource"`
	// +optional
	ActionSourceRef *api.ConfigMapRef `json:"actionSourceRef,omitempty"`
	// +optional
	ExtraArgs string `json:"extraArgs"`
	// +required
	BackoffLimit int `json:"BackoffLimit"`
	// Image is the kubespray base image to create k8s clusters
	// +required
	Image string `json:"image"`
	// +optional
	PreHook []HookAction `json:"preHook"`
	// +optional
	PostHook []HookAction `json:"postHook"`
	// +optional
	Resource corev1.ResourceRequirements `json:"resource"`
	// +optional
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds"`
}

type (
	ActionType   string
	ActionSource string
)

type HookAction struct {
	// +required
	ActionType ActionType `json:"actionType"`
	// +required
	Action string `json:"action"`
	// +optional
	// +kubebuilder:default="builtin"
	ActionSource *ActionSource `json:"actionSource"`
	// +optional
	ActionSourceRef *api.ConfigMapRef `json:"actionSourceRef,omitempty"`
	// +optional
	ExtraArgs string `json:"extraArgs"`
}

// ClusterOperationStatus defines the observed state of ClusterOperation
type ClusterOperationStatus struct {
	// +optional
	Action string `json:"action"`
	// +optional
	JobRef *api.JobRef `json:"jobRef,omitempty"`
	// +optional
	Status OpsStatus `json:"status"`
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// +optional
	EndTime *metav1.Time `json:"endTime,omitempty"`
	// Digest it used to avoid the change of clusterOps by others
	// It will be filled by operator
	// +optional
	Digest string `json:"digest,omitempty"`
	// HasModified indicates the spec has been modified by others after created.
	// +optional
	HasModified bool `json:"hasModified,omitempty"`
}

type OpsStatus string

const (
	RunningStatus   OpsStatus = "Running"
	SucceededStatus OpsStatus = "Succeeded"
	FailedStatus    OpsStatus = "Failed"
	BlockStatus     OpsStatus = "Blocked"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClusterOperation is the Schema for the clusteroperations API
type ClusterOperation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterOperationSpec   `json:"spec,omitempty"`
	Status ClusterOperationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterOperationList contains a list of ClusterOperation
type ClusterOperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterOperation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterOperation{}, &ClusterOperationList{})
}
