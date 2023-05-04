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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	// HostsConfRef represent a configMap which actually store the hosts.yml in ansible
	// +required
	HostsConfRef *api.ConfigMapRef `json:"hostsConfRef"`
	// VarsConfRef represent a configMap which actually store the group_vars.yml in ansible
	// +required
	VarsConfRef *api.ConfigMapRef `json:"varsConfRef"`
	// KubeConfRef represent a configMap which actually store kubeconfig
	// +optional
	KubeConfRef *api.ConfigMapRef `json:"kubeConfRef"`
	// SSHAuthRef store sshkey.If it is empty ansible will use ssh password
	// +optional
	SSHAuthRef *api.SecretRef `json:"SSHAuthRef"`
	// +optional
	PreCheckRef *api.ConfigMapRef `json:"preCheckRef"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	Conditions []ClusterCondition `json:"conditions"`
}

type ClusterCondition struct {
	// ClusterOps refers to the name of ClusterOperation
	// +required
	ClusterOps string `json:"clusterOps"`
	// +optional
	Status ClusterConditionType `json:"clusterConditionType"`
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// +optional
	EndTime *metav1.Time `json:"endTime,omitempty"`
}

type ClusterConditionType string

const (
	ClusterCreating ClusterConditionType = "Running"

	ClusterRunning ClusterConditionType = "Succeeded"

	ClusterUpdating ClusterConditionType = "Failed"

	BlockedStatus ClusterConditionType = "Blocked"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
