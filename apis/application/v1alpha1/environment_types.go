	/*
Copyright 2020 The Crossplane Authors.

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
		metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
		"k8s.io/apimachinery/pkg/runtime/schema"
		"k8s.io/apimachinery/pkg/types"
		"reflect"

		xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	)

// EnvironmentParameters are the configurable fields of an Environment.
type EnvironmentParameters struct {
	ApplicationName string `json:"applicationName,omitempty"`
	TenantName string `json:"tenantName,omitempty"`
	EnvironmentType string `json:"environmentType"`
	ExpirationTime metav1.Time `json:"expirationTime,omitempty"`
	Owner string `json:"owner"`
}

// EnvironmentObservation are the observable fields of an Environment.
type EnvironmentObservation struct {
	Instances []EnvironmentXRC `json:"instances,omitempty"`

}

type EnvironmentXRC struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	UID types.UID `json:"uid,omitempty"`
	Status bool `json:"status"`
}

// A EnvironmentSpec defines the desired state of an Environment.
type EnvironmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EnvironmentParameters `json:"forProvider"`
}

// A EnvironmentStatus represents the observed state of an Environment.
type EnvironmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A MyType is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,orchestrator}
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvironmentSpec   `json:"spec"`
	Status EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environments
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Environment type metadata.
var (
	EnvironmentKind             = reflect.TypeOf(Environment{}).Name()
	EnvironmentGroupKind        = schema.GroupKind{Group: Group, Kind: EnvironmentKind}.String()
	EnvironmentKindAPIVersion   = EnvironmentKind + "." + SchemeGroupVersion.String()
	EnvironmentGroupVersionKind = SchemeGroupVersion.WithKind(EnvironmentKind)
)

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
