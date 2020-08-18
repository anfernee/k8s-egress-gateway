/*
Copyright 2020 The Kubernetes authors.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EgressGatewayClassSpec defines the desired state of EgressGatewayClass
type EgressGatewayClassSpec struct {

	// Controller is a string that indicates the controller that manages the gateway class
	// +required
	Controller string `json:"controller" protobuf:"bytes,1,opt,name=controller"`
}

// EgressGatewayClassStatus defines the observed state of EgressGatewayClass
type EgressGatewayClassStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// EgressGatewayClass is the Schema for the egressgatewayclasses API
type EgressGatewayClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EgressGatewayClassSpec   `json:"spec,omitempty"`
	Status EgressGatewayClassStatus `json:"status,omitempty"`
}

// RouteNamespaces is used by Gateway and GatewayClass to indicate which
// namespaces Routes should be selected from.
type RouteNamespaces struct {
	// NamespaceSelector is a selector of namespaces that Routes should be
	// selected from. This is a standard Kubernetes LabelSelector, a label query
	// over a set of resources. The result of matchLabels and matchExpressions
	// are ANDed. Controllers must not support Routes in namespaces outside this
	// selector.
	//
	// An empty selector (default) indicates that Routes in any namespace can be
	// selected.
	//
	// The OnlySameNamespace field takes precedence over this field. This
	// selector will only take effect when OnlySameNamespace is false.
	//
	// Support: Core
	//
	// +optional
	NamespaceSelector metav1.LabelSelector `json:"namespaceSelector" protobuf:"bytes,1,opt,name=namespaceSelector"`

	// OnlySameNamespace is a boolean used to indicate if Route references are
	// limited to the same Namespace as the Gateway. When true, only Routes
	// within the same Namespace as the Gateway should be selected.
	//
	// This field takes precedence over the NamespaceSelector field. That
	// selector should only take effect when this field is set to false.
	//
	// Support: Core
	//
	// +optional
	OnlySameNamespace bool `json:"onlySameNamespace" protobuf:"bytes,2,opt,name=onlySameNamespace"`
}

// +kubebuilder:object:root=true

// EgressGatewayClassList contains a list of EgressGatewayClass
type EgressGatewayClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EgressGatewayClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EgressGatewayClass{}, &EgressGatewayClassList{})
}
