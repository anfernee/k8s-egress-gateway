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

// TCPRouteSpec defines the desired state of TCPRoute
// The route should be supported on CNI side
type TCPRouteSpec struct {
	// Rules ...
	Rules []TCPRouteRule `json:"rules" protobuf:"rules,1,rep,name=rules"`
}

// TCPRouteStatus defines the observed state of TCPRoute
type TCPRouteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// TCPRoute is the Schema for the tcproutes API
type TCPRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TCPRouteSpec   `json:"spec,omitempty"`
	Status TCPRouteStatus `json:"status,omitempty"`
}

// TCPRouteRule is the configuration for a given rule.
type TCPRouteRule struct {
	// Match defines which connections match this rule.
	//
	// +optional
	Match *TCPRouteMatch `json:"match" protobuf:"bytes,1,opt,name=match"`

	// Action defines what happens to the connection.
	//
	// +optional
	Action *TCPRouteAction `json:"action" protobuf:"bytes,2,opt,name=action"`
}

// TCPRouteAction is the action for a given rule.
type TCPRouteAction struct {
	// ForwardTo sends requests to the referenced object(s).  The
	// resource may be "services" (omit or use the empty string for the
	// group), or an implementation may support other resources (for
	// example, resource "myroutetargets" in group "networking.acme.io").
	// Omitting or specifying the empty string for both the resource and
	// group indicates that the resource is "services".  If the referent
	// cannot be found, the "InvalidRoutes" status condition on any Gateway
	// that includes the HTTPRoute will be true.
	//
	// Support: core
	//
	// +kubebuilder:validation:MinItems=1
	ForwardTo []ForwardToTarget `json:"forwardTo" protobuf:"bytes,1,rep,name=forwardTo"`
}

// TCPRouteMatch defines the predicate used to match connections to a
// given action.
type TCPRouteMatch struct {

	// RouteNamespaces indicates in which namespaces Routes should be selected
	// for this Gateway. This is restricted to the namespace of this Gateway by
	// default.
	//
	// Support: Core
	//
	// +optional
	RouteNamespaces RouteNamespaces `json:"routeNamespaces,omitempty" protobuf:"bytes,1,opt,name=routeNamespaces"`

	// RouteSelector specifies a set of route labels used for selecting
	// routes to associate with the Gateway. If RouteSelector is defined,
	// only routes matching the RouteSelector are associated with the Gateway.
	// An empty RouteSelector matches all routes.
	//
	// Support: Core
	//
	// +optional
	RouteSelector metav1.LabelSelector `json:"routeSelector,omitempty" protobuf:"bytes,2,opt,name=routeSelector"`

	// ExtensionRef is an optional, implementation-specific extension to the
	// "match" behavior.  The resource may be "configmap" (use the empty
	// string for the group) or an implementation-defined resource (for
	// example, resource "myroutematchers" in group "networking.acme.io").
	// Omitting or specifying the empty string for both the resource and
	// group indicates that the resource is "configmaps".  If the referent
	// cannot be found, the "InvalidRoutes" status condition on any Gateway
	// that includes the TCPRoute will be true.
	//
	// Support: custom
	//
	// +optional
	// ExtensionRef *RouteMatchExtensionObjectReference `json:"extensionRef" protobuf:"bytes,1,opt,name=extensionRef"`
}

// +kubebuilder:object:root=true

// TCPRouteList contains a list of TCPRoute
type TCPRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TCPRoute{}, &TCPRouteList{})
}
