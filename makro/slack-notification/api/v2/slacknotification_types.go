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

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SlackNotificationSpec defines the desired state of SlackNotification.
type SlackNotificationSpec struct {
	Namespace       string   `json:"namespace"`
	EventTypes      []string `json:"eventTypes,omitempty"`
	ReasonCodes     []string `json:"reasonCodes,omitempty"`
	SlackWebhookURL string   `json:"slackWebhookURL,omitempty"`
}

// SlackNotificationStatus defines the observed state of SlackNotification.
type SlackNotificationStatus struct {
	LastNotificationTime metav1.Time `json:"lastNotificationTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SlackNotification is the Schema for the slacknotifications API.
type SlackNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SlackNotificationSpec   `json:"spec,omitempty"`
	Status SlackNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SlackNotificationList contains a list of SlackNotification.
type SlackNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SlackNotification `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SlackNotification{}, &SlackNotificationList{})
}
