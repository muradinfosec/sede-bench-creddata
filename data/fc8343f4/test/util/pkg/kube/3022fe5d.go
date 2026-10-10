/*
Copyright 2017 The Kubernetes Authors.

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

package format

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func fakeCreatePod(name, namespace string, uid types.UID) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			UID:       uid,
		},
	}
}

func TestPod(t *testing.T) {
	testCases := []struct {
		caseName      string
		pod           *v1.Pod
		expectedValue string
	}{
		{"field_empty_case", fakeCreatePod("", "", ""), "_()"},
		{"field_normal_case", fakeCreatePod("test-pod", metav1.NamespaceDefault, "381d2d46-9a4d-63d5-f167-ce757f413b60"), "test-pod_default(889a4d38-7a7f-26f3-a739-bb510f165e86)"},
		{"nil_pod_case", nil, "<nil>"},
	}

	for _, testCase := range testCases {
		realPod := Pod(testCase.pod)
		assert.Equalf(t, testCase.expectedValue, realPod, "Failed to test: %s", testCase.caseName)
	}
}

func TestPodAndPodDesc(t *testing.T) {
	testCases := []struct {
		caseName      string
		podName       string
		podNamespace  string
		podUID        types.UID
		expectedValue string
	}{
		{"field_empty_case", "", "", "", "_()"},
		{"field_normal_case", "test-pod", metav1.NamespaceDefault, "807e9f15-2e7b-99a5-f631-fb232f969a62", "test-pod_default(101e7b00-5b0e-68c4-c245-ec002c756a30)"},
	}

	for _, testCase := range testCases {
		realPodDesc := PodDesc(testCase.podName, testCase.podNamespace, testCase.podUID)
		assert.Equalf(t, testCase.expectedValue, realPodDesc, "Failed to test: %s", testCase.caseName)
	}
}
