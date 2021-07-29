/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha2

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ReferencePolicy identifies kinds of resources in other namespaces that are
// trusted to reference the specified kinds of resources in the local namespace.
// Each ReferencePolicy can be used to represent a unique trust relationship.
// Additional Reference Policies can be used to add to the set of trusted
// sources of inbound references for the namespace they are defined within.
type ReferencePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of ReferencePolicy.
	Spec ReferencePolicySpec `json:"spec,omitempty"`
}

// ReferencePolicySpec identifies a cross namespace relationship that is trusted
// for Gateway API.
type ReferencePolicySpec struct {
	// From describes the trusted namespaces and kinds that can reference the
	// resources described in "To". Each entry in this list must be considered
	// to be an additional place that references can be valid from, or to put
	// this another way, entries must be combined using OR.
	//
	// Support: Core
	//
	// +kubebuilder:validation:MinItems=1
	From []ReferencePolicyFrom `json:"from"`

	// To describes the resources that may be referenced by the resources
	// described in "From". Each entry in this list must be considered to be an
	// additional place that references can be valid to, or to put this another
	// way, entries must be combined using OR.
	//
	// Support: Core
	//
	// +kubebuilder:validation:MinItems=1
	To []ReferencePolicyTo `json:"to"`
}

// ReferencePolicyFrom describes trusted namespaces and kinds.
type ReferencePolicyFrom struct {
	// Group is the group of the referent.
	//
	// Support: Core
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Group string `json:"group"`

	// Kind is the kind of the referent. Although implementations may support
	// additional resources, the following Route types are part of the "Core"
	// support level for this field:
	//
	// * HTTPRoute
	// * TCPRoute
	// * TLSRoute
	// * UDPRoute
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Kind string `json:"kind"`

	// Namespace is the namespace of the referent.
	//
	// Support: Core
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Namespace string `json:"namespace,omitempty"`
}

// ReferencePolicyTo describes what Kinds are allowed as targets of the
// references.
type ReferencePolicyTo struct {
	// Group is the group of the referent.
	//
	// Support: Core
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Group string `json:"group"`

	// Kind is the kind of the referent. Although implementations may support
	// additional resources, the following types are part of the "Core"
	// support level for this field:
	//
	// * Service
	// * HTTPRoute
	// * TCPRoute
	// * TLSRoute
	// * UDPRoute
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Kind string `json:"kind"`
}
