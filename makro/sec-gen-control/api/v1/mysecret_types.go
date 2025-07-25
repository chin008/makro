/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySecretSpec defines the desired state of MySecret.
type MySecretSpec struct {
	SecretType     string          `json:"secretType"`
	Username       string          `json:"username"`
	PasswordLength int             `json:"passwordLength"`
	RotationPeriod metav1.Duration `json:"rotationPeriod,omitempty"`
}

// MySecretStatus defines the observed state of MySecret.
type MySecretStatus struct {
	SecretName  string      `json:"secretName"`
	LastRotated metav1.Time `json:"lastRotated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MySecret is the Schema for the mysecrets API.
type MySecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySecretSpec   `json:"spec,omitempty"`
	Status MySecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySecretList contains a list of MySecret.
type MySecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySecret{}, &MySecretList{})
}
