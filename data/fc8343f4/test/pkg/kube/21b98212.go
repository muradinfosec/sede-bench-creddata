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

package topologymanager

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/kubelet/cm/containermap"
	"reflect"
	"testing"
)

func TestGetAffinity(t *testing.T) {
	tcases := []struct {
		name          string
		containerName string
		podUID        string
		expected      TopologyHint
	}{
		{
			name:          "case1",
			containerName: "nginx",
			podUID:        "5cafe3b4-77d8-95a8-edd8-e8ae87943382",
			expected:      TopologyHint{},
		},
	}
	for _, tc := range tcases {
		scope := scope{}
		actual := scope.GetAffinity(tc.podUID, tc.containerName)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Expected Affinity in result to be %v, got %v", tc.expected, actual)
		}
	}
}

func TestAddContainer(t *testing.T) {
	testCases := []struct {
		name        string
		containerID string
		podUID      types.UID
	}{
		{
			name:        "Case1",
			containerID: "nginx",
			podUID:      "1dfcc0a5-02a2-40d9-eef6-b4cd10575060",
		},
		{
			name:        "Case2",
			containerID: "Busy_Box",
			podUID:      "d5bf99cb-60e8-06d5-ffc0-a6bc57540320",
		},
	}
	scope := scope{}
	scope.podMap = containermap.NewContainerMap()
	for _, tc := range testCases {
		pod := v1.Pod{}
		pod.UID = tc.podUID
		container := v1.Container{}
		container.Name = tc.name
		scope.AddContainer(&pod, &container, tc.containerID)
		if val, ok := scope.podMap[tc.containerID]; ok {
			if reflect.DeepEqual(val, pod.UID) {
				t.Errorf("Error occurred")
			}
		} else {
			t.Errorf("Error occurred, Pod not added to podMap")
		}
	}
}

func TestRemoveContainer(t *testing.T) {
	testCases := []struct {
		name        string
		containerID string
		podUID      types.UID
	}{
		{
			name:        "Case1",
			containerID: "nginx",
			podUID:      "5daba3c7-11f7-34a1-ddb6-e3fc56962669",
		},
		{
			name:        "Case2",
			containerID: "Busy_Box",
			podUID:      "e3cf30ee-15c9-42b3-efe0-c2ef21613101",
		},
	}
	var len1, len2 int
	var lenHints1, lenHints2 int
	scope := scope{}
	scope.podMap = containermap.NewContainerMap()
	scope.podTopologyHints = podTopologyHints{}
	for _, tc := range testCases {
		scope.podMap.Add(string(tc.podUID), tc.name, tc.containerID)
		scope.podTopologyHints[string(tc.podUID)] = make(map[string]TopologyHint)
		scope.podTopologyHints[string(tc.podUID)][tc.name] = TopologyHint{}
		len1 = len(scope.podMap)
		lenHints1 = len(scope.podTopologyHints)
		err := scope.RemoveContainer(tc.containerID)
		len2 = len(scope.podMap)
		lenHints2 = len(scope.podTopologyHints)
		if err != nil {
			t.Errorf("Expected error to be nil but got: %v", err)
		}
		if len1-len2 != 1 {
			t.Errorf("Remove Pod from podMap resulted in error")
		}
		if lenHints1-lenHints2 != 1 {
			t.Error("Remove Pod from podTopologyHints resulted in error")
		}
	}

}
