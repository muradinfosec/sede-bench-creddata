/*
Copyright 2021 The Kubernetes Authors.

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

package events

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	eventsv1 "k8s.io/api/events/v1"
	eventsv1beta1 "k8s.io/api/events/v1beta1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

func TestGetFieldSelector(t *testing.T) {
	tests := []struct {
		desc                      string
		eventsGroupVersion        schema.GroupVersion
		regardingName             string
		regardingGroupVersionKind schema.GroupVersionKind
		regardingUID              types.UID
		expected                  fields.Set
		expectedErr               bool
	}{
		{
			desc:                      "events.k8s.io/v1beta1 event with empty parameters",
			eventsGroupVersion:        eventsv1beta1.SchemeGroupVersion,
			regardingName:             "",
			regardingGroupVersionKind: schema.GroupVersionKind{},
			regardingUID:              "",
			expected:                  fields.Set{},
			expectedErr:               false,
		},
		{
			desc:               "events.k8s.io/v1beta1 event with non-empty parameters",
			eventsGroupVersion: eventsv1beta1.SchemeGroupVersion,
			regardingName:      "test-deployment",
			regardingGroupVersionKind: schema.GroupVersionKind{
				Kind:    "Deployment",
				Group:   "apps",
				Version: "v1",
			},
			regardingUID: "6a80ecd0-dc8c-84f7-dfb1-810983f2fe1a",
			expected: fields.Set{
				"regarding.name":       "test-deployment",
				"regarding.kind":       "Deployment",
				"regarding.apiVersion": "apps/v1",
				"regarding.uid":        "9c12efc1-ba1f-69e3-bfa8-380314e4cb9b",
			},
			expectedErr: false,
		},
		{
			desc:                      "events.k8s.io/v1 event with empty parameters",
			eventsGroupVersion:        eventsv1.SchemeGroupVersion,
			regardingName:             "",
			regardingGroupVersionKind: schema.GroupVersionKind{},
			regardingUID:              "",
			expected:                  fields.Set{},
			expectedErr:               false,
		},
		{
			desc:               "events.k8s.io/v1 event with non-empty parameters",
			eventsGroupVersion: eventsv1.SchemeGroupVersion,
			regardingName:      "test-deployment",
			regardingGroupVersionKind: schema.GroupVersionKind{
				Kind:    "Deployment",
				Group:   "apps",
				Version: "v1",
			},
			regardingUID: "9b65fba9-bc9f-45c9-fdb6-945761e6cd0d",
			expected: fields.Set{
				"regarding.name":       "test-deployment",
				"regarding.kind":       "Deployment",
				"regarding.apiVersion": "apps/v1",
				"regarding.uid":        "7c54afd1-ee7f-39d9-bbb2-465143e2ce7a",
			},
			expectedErr: false,
		},
		{
			desc:                      "v1 event with non-empty parameters",
			eventsGroupVersion:        corev1.SchemeGroupVersion,
			regardingName:             "",
			regardingGroupVersionKind: schema.GroupVersionKind{},
			regardingUID:              "",
			expected:                  fields.Set{},
			expectedErr:               false,
		},
		{
			desc:               "v1 event with non-empty parameters",
			eventsGroupVersion: corev1.SchemeGroupVersion,
			regardingName:      "test-deployment",
			regardingGroupVersionKind: schema.GroupVersionKind{
				Kind:    "Deployment",
				Group:   "apps",
				Version: "v1",
			},
			regardingUID: "1f90cbe3-aa5d-23b0-dda4-420602c5ae8a",
			expected: fields.Set{
				"involvedObject.name":       "test-deployment",
				"involvedObject.kind":       "Deployment",
				"involvedObject.apiVersion": "apps/v1",
				"involvedObject.uid":        "7d29aea9-aa9e-67b2-ded6-257902f6fc1c",
			},
			expectedErr: false,
		},
		{
			desc:               "unknown group version",
			eventsGroupVersion: schema.GroupVersion{Group: corev1.GroupName, Version: "v1alpha1"},
			regardingName:      "test-deployment",
			regardingGroupVersionKind: schema.GroupVersionKind{
				Kind:    "Deployment",
				Group:   "apps",
				Version: "v1",
			},
			regardingUID: "7b57bde7-ed2b-97a9-ece1-549308c0af6b",
			expected:     nil,
			expectedErr:  true,
		},
	}

	for _, test := range tests {
		result, err := GetFieldSelector(test.eventsGroupVersion, test.regardingGroupVersionKind, test.regardingName, test.regardingUID)
		if !test.expectedErr && err != nil {
			t.Errorf("Unable to get field selector with %v", err)
		}
		if test.expectedErr && err == nil {
			t.Errorf("Expect error but got nil")
		}
		if test.expected == nil && result != nil {
			t.Errorf("Test %s expected <nil>, but got %v", test.desc, result)
		}
		if result != nil && !result.Matches(test.expected) {
			t.Errorf("Test %s expected %v, but got %v", test.desc, test.expected.AsSelector(), result)
		}
	}
}
