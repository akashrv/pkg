/*
Copyright 2018 The Knative Authors

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
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/knative/pkg/apis"
	"github.com/knative/pkg/apis/duck"
)

// Addressable provides a generic mechanism for a custom resource
// definition to indicate a destination for message delivery.
// (Currently, only hostname, port and path are supported, and HTTP is implied. In the
// future, additional schemes may be supported
// ala UI may also be supported.)

// Addressable is the schema for the destination information. This is
// typically stored in the object's `status`, as this information may
// be generated by the controller.
type Addressable struct {
	Hostname string `json:"hostname,omitempty"`
	Port     string `json:"port,omitempty"`
	Path     string `json:"path,omitempty"`
}

// Addressable is an Implementable "duck type".
var _ duck.Implementable = (*Addressable)(nil)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AddressableType is a skeleton type wrapping Addressable in the manner we expect
// resource writers defining compatible resources to embed it.  We will
// typically use this type to deserialize Addressable ObjectReferences and
// access the Addressable data.  This is not a real resource.
type AddressableType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status AddressStatus `json:"status"`
}

// AddressStatus shows how we expect folks to embed Addressable in
// their Status field.
type AddressStatus struct {
	Address *Addressable `json:"address,omitempty"`
}

// Verify AddressableType resources meet duck contracts.
var _ duck.Populatable = (*AddressableType)(nil)
var _ apis.Listable = (*AddressableType)(nil)

// GetFullType implements duck.Implementable
func (_ *Addressable) GetFullType() duck.Populatable {
	return &AddressableType{}
}

// Populate implements duck.Populatable
func (t *AddressableType) Populate() {
	t.Status = AddressStatus{
		&Addressable{
			// Populate ALL fields
			Hostname: "this is not empty",
			Path:     "this is not empty",
		},
	}
}

// GetListType implements apis.Listable
func (r *AddressableType) GetListType() runtime.Object {
	return &AddressableTypeList{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AddressableTypeList is a list of AddressableType resources
type AddressableTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AddressableType `json:"items"`
}
