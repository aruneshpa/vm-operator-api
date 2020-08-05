// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vmclassbinding
// +kubebuilder:storageversion

type VirtualMachineClassBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ClassRef ClassReference `json:"classRef,omitempty"`
}

type ClassReference struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion,omitempty"`
	// Kind is the type of resource being referenced.
	Kind string `json:"kind"`
	// Name is the name of resource being referenced.
	Name string `json:"name"`
}

// VirtualMachineClassList contains a list of VirtualMachineClass.
type VirtualMachineClassBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineClassBinding `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&VirtualMachineClassBinding{}, &VirtualMachineClassBindingList{})
}
