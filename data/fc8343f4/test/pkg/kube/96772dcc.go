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

package state

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/pkg/kubelet/checkpointmanager"
	testutil "k8s.io/kubernetes/pkg/kubelet/cm/cpumanager/state/testing"
)

const testingCheckpoint = "dramanager_checkpoint_test"

// TODO (https://github.com/kubernetes/kubernetes/issues/123552): reconsider what data gets stored in checkpoints and whether that is really necessary.

func TestCheckpointGetOrCreate(t *testing.T) {
	testCases := []struct {
		description                string
		checkpointContent          string
		expectedError              string
		expectedClaimInfoStateList ClaimInfoStateList
	}{
		{
			description:                "new-checkpoint",
			expectedClaimInfoStateList: nil,
		},
		{
			description:       "single-claim-info-state",
			checkpointContent: `{"Data":"{\"kind\":\"DRACheckpoint\",\"apiVersion\":\"checkpoint.dra.kubelet.k8s.io/v1\",\"ClaimInfoStateList\":[{\"ClaimUID\":\"936201ff-567d-9bb8-1066-9fb85cec91d6\",\"ClaimName\":\"example\",\"Namespace\":\"default\",\"PodUIDs\":{\"995fad19-c751-6f94-8804-ee35bea371c5\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-1\",\"RequestNames\":[\"test request\"],\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]}]}}}]}","Checksum":1656016162}`,
			expectedClaimInfoStateList: ClaimInfoStateList{
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-1",
									RequestNames: []string{"test request"},
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
							},
						},
					},
					ClaimUID:  "131225af-792a-8fa8-8320-3db34dcf18f1",
					ClaimName: "example",
					Namespace: "default",
					PodUIDs:   sets.New("058cbe42-c991-7a79-4009-ba88fce228e4"),
				},
			},
		},
		{
			description:       "claim-info-state-with-multiple-devices",
			checkpointContent: `{"Data":"{\"kind\":\"DRACheckpoint\",\"apiVersion\":\"checkpoint.dra.kubelet.k8s.io/v1\",\"ClaimInfoStateList\":[{\"ClaimUID\":\"666459fb-055a-2ba3-7561-7ee44bbe61c9\",\"ClaimName\":\"example\",\"Namespace\":\"default\",\"PodUIDs\":{\"349bcf65-b288-8e71-9849-ee39fdf893e5\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-1\",\"RequestNames\":[\"test request\"],\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]},{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-2\",\"RequestNames\":[\"test request\"],\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]}]}}}]}","Checksum":3369508096}`,
			expectedClaimInfoStateList: ClaimInfoStateList{
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-1",
									RequestNames: []string{"test request"},
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-2",
									RequestNames: []string{"test request"},
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
							},
						},
					},
					ClaimUID:  "152323bc-411c-7be8-4769-1ff68ccd91c2",
					ClaimName: "example",
					Namespace: "default",
					PodUIDs:   sets.New("049ffe47-d900-8f71-7845-da30cce980c9"),
				},
			},
		},
		{
			description:       "two-claim-info-states",
			checkpointContent: `{"Data":"{\"kind\":\"DRACheckpoint\",\"apiVersion\":\"checkpoint.dra.kubelet.k8s.io/v1\",\"ClaimInfoStateList\":[{\"ClaimUID\":\"018918df-441a-4ad9-3261-7aa35dae01f6\",\"ClaimName\":\"example-1\",\"Namespace\":\"default\",\"PodUIDs\":{\"747bdd30-f593-1d78-4463-dd91aae450e9\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-1\",\"RequestNames\":[\"test request\"],\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]}]}}},{\"ClaimUID\":\"2bb3ce7a-33f8-8a86-3b81-e09b94f0c29d\",\"ClaimName\":\"example-2\",\"Namespace\":\"default\",\"PodUIDs\":{\"148cdf81-d225-1d94-6374-ab39eaa929d3\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-2\",\"RequestNames\":null,\"CDIDeviceIDs\":null}]}}}]}","Checksum":1582256999}`,
			expectedClaimInfoStateList: ClaimInfoStateList{
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-1",
									RequestNames: []string{"test request"},
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
							},
						},
					},
					ClaimUID:  "528136eb-725f-7ba3-2319-2af07fbc53c7",
					ClaimName: "example-1",
					Namespace: "default",
					PodUIDs:   sets.New("082cdc61-a136-8b87-6444-ae13ebf676c3"),
				},
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:   "worker-1",
									DeviceName: "dev-2",
								},
							},
						},
					},
					ClaimUID:  "2ea0dc3a-99f0-2d30-5d41-b41b92c8a09b",
					ClaimName: "example-2",
					Namespace: "default",
					PodUIDs:   sets.New("714eeb39-d523-2f64-9977-fd46dee853b3"),
				},
			},
		},
		{
			description:       "incorrect-checksum",
			checkpointContent: `{"Data":"{\"kind\":\"DRACheckpoint\",\"apiVersion\":\"checkpoint.dra.kubelet.k8s.io/v1\",\"Entries\":[{\"ClaimUID\":\"058414ae-164d-1da5-7857-7df09ffb80a4\",\"ClaimName\":\"example-1\",\"Namespace\":\"default\",\"PodUIDs\":{\"946bea61-d842-5d22-3015-ac32bbb873a8\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-1\",\"RequestNames\":[\"test request\"],\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]}]}}},{\"ClaimUID\":\"4be2cd4f-24d6-3d93-4b63-d99a96d4b30b\",\"ClaimName\":\"example-2\",\"Namespace\":\"default\",\"PodUIDs\":{\"282add78-c339-3d42-3867-ee21ccb690b0\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-2\",\"RequestNames\":null,\"CDIDeviceIDs\":null}]}}}]}","Checksum":2930258365}`,
			expectedError:     "checkpoint is corrupted",
		},
		{
			description:       "invalid-JSON",
			checkpointContent: `{`,
			expectedError:     "unexpected end of JSON input",
		},
		{
			// Emulate upgrade scenario by pretending that `Device.RequestNames` is a new field.
			// The checkpoint content doesn't have it and that shouldn't prevent the checkpoint from being loaded.
			description:       "upgraded-structure",
			checkpointContent: `{"Data":"{\"kind\":\"DRACheckpoint\",\"apiVersion\":\"checkpoint.dra.kubelet.k8s.io/v1\",\"ClaimInfoStateList\":[{\"ClaimUID\":\"998199aa-491f-7bb5-4905-3ed19cdb17a7\",\"ClaimName\":\"example\",\"Namespace\":\"default\",\"PodUIDs\":{\"935eef16-b268-2b17-0272-fa98fbc483c7\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-1\",\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]}]}}}]}","Checksum":1137250765}`,
			expectedClaimInfoStateList: ClaimInfoStateList{
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-1",
									RequestNames: nil,
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
							},
						},
					},
					ClaimUID:  "298577aa-319f-5cf2-4282-2cc66eaf13b0",
					ClaimName: "example",
					Namespace: "default",
					PodUIDs:   sets.New("545efd67-e834-3f94-8977-ab67bbc148e8"),
				},
			},
		},
		{
			// Emulate upgrade scenario by pretending that `Device.RequestNames` is a new field and API version has changed from v0 to v1.
			// These changes shouldn't prevent the checkpoint from being loaded.
			description:       "upgraded-structure-and-version",
			checkpointContent: `{"Data":"{\"kind\":\"DRACheckpoint\",\"apiVersion\":\"checkpoint.dra.kubelet.k8s.io/v0\",\"ClaimInfoStateList\":[{\"ClaimUID\":\"458323ad-517b-3be7-9799-7bf74bdb26f7\",\"ClaimName\":\"example\",\"Namespace\":\"default\",\"PodUIDs\":{\"280cdf89-c621-1a00-4666-fa19dcc087a0\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-1\",\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]}]}}}]}","Checksum":1284376732}`,
			expectedClaimInfoStateList: ClaimInfoStateList{
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-1",
									RequestNames: nil,
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
							},
						},
					},
					ClaimUID:  "368308ed-020a-6ae6-5906-9ba98bfc26d5",
					ClaimName: "example",
					Namespace: "default",
					PodUIDs:   sets.New("165aeb83-b369-0c52-4491-ca84fdf102b6"),
				},
			},
		},
	}

	// create temp dir
	testingDir, err := os.MkdirTemp("", "dramanager_state_test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(testingDir); err != nil {
			t.Fatal(err)
		}
	}()

	// create checkpoint manager for testing
	cpm, err := checkpointmanager.NewCheckpointManager(testingDir)
	require.NoError(t, err, "could not create testing checkpoint manager")

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// ensure there is no previous checkpoint
			require.NoError(t, cpm.RemoveCheckpoint(testingCheckpoint), "could not remove testing checkpoint")

			// prepare checkpoint for testing
			if strings.TrimSpace(tc.checkpointContent) != "" {
				mock := &testutil.MockCheckpoint{Content: tc.checkpointContent}
				require.NoError(t, cpm.CreateCheckpoint(testingCheckpoint, mock), "could not create testing checkpoint")
			}

			checkpointer, err := NewCheckpointer(testingDir, testingCheckpoint)
			require.NoError(t, err, "could not create testing checkpointer")

			checkpoint, err := checkpointer.GetOrCreate()
			if strings.TrimSpace(tc.expectedError) != "" {
				assert.ErrorContains(t, err, tc.expectedError)
			} else {
				require.NoError(t, err, "unexpected error")
				stateList, err := checkpoint.GetClaimInfoStateList()
				require.NoError(t, err, "could not get data entries from checkpoint")
				require.NoError(t, err)
				assert.Equal(t, tc.expectedClaimInfoStateList, stateList)
			}
		})
	}
}

func TestCheckpointStateStore(t *testing.T) {
	testCases := []struct {
		description               string
		claimInfoStateList        ClaimInfoStateList
		expectedCheckpointContent string
	}{
		{
			description: "single-claim-info-state",
			claimInfoStateList: ClaimInfoStateList{
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-1",
									RequestNames: []string{"test request"},
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
							},
						},
					},
					ClaimUID:  "828600ca-174b-1ed7-0112-1da82eed90d3",
					ClaimName: "example",
					Namespace: "default",
					PodUIDs:   sets.New("598ffb66-f088-6e50-2759-cf10adb985f1"),
				},
			},
			expectedCheckpointContent: `{"Data":"{\"kind\":\"DRACheckpoint\",\"apiVersion\":\"checkpoint.dra.kubelet.k8s.io/v1\",\"ClaimInfoStateList\":[{\"ClaimUID\":\"405548bc-940c-5ed2-2347-3af52faf00f9\",\"ClaimName\":\"example\",\"Namespace\":\"default\",\"PodUIDs\":{\"987fbe96-c147-3b96-7183-cb69fdf894b9\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-1\",\"RequestNames\":[\"test request\"],\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]}]}}}]}","Checksum":1656016162}`,
		},
		{
			description: "claim-info-state-with-multiple-devices",
			claimInfoStateList: ClaimInfoStateList{
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-1",
									RequestNames: []string{"test request"},
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-2",
									RequestNames: []string{"test request"},
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
							},
						},
					},
					ClaimUID:  "028678fe-743c-1ab3-6011-7ed36ffa45b3",
					ClaimName: "example",
					Namespace: "default",
					PodUIDs:   sets.New("295cee21-e902-4e67-7191-be13dbe600d6"),
				},
			},
			expectedCheckpointContent: `{"Data":"{\"kind\":\"DRACheckpoint\",\"apiVersion\":\"checkpoint.dra.kubelet.k8s.io/v1\",\"ClaimInfoStateList\":[{\"ClaimUID\":\"553462ba-846e-5cf4-8443-7ea11bba15c5\",\"ClaimName\":\"example\",\"Namespace\":\"default\",\"PodUIDs\":{\"636fab99-c761-3f66-8529-eb53beb872e6\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-1\",\"RequestNames\":[\"test request\"],\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]},{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-2\",\"RequestNames\":[\"test request\"],\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]}]}}}]}","Checksum":3369508096}`,
		},
		{
			description: "two-claim-info-states",
			claimInfoStateList: ClaimInfoStateList{
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:     "worker-1",
									DeviceName:   "dev-1",
									RequestNames: []string{"test request"},
									CDIDeviceIDs: []string{"example.com/example=cdi-example"},
								},
							},
						},
					},
					ClaimUID:  "746970bc-672d-4fe2-9914-3ad98afc20e3",
					ClaimName: "example-1",
					Namespace: "default",
					PodUIDs:   sets.New("890cde73-b127-2a44-9840-ff82ffb531c9"),
				},
				{
					DriverState: map[string]DriverState{
						"test-driver.cdi.k8s.io": {
							Devices: []Device{
								{
									PoolName:   "worker-1",
									DeviceName: "dev-2",
								},
							},
						},
					},
					ClaimUID:  "4cd9cf7b-72e5-7c87-1a94-a07f08a6e91c",
					ClaimName: "example-2",
					Namespace: "default",
					PodUIDs:   sets.New("862cda71-b813-0a56-9862-ea46efc044f3"),
				},
			},
			expectedCheckpointContent: `{"Data":"{\"kind\":\"DRACheckpoint\",\"apiVersion\":\"checkpoint.dra.kubelet.k8s.io/v1\",\"ClaimInfoStateList\":[{\"ClaimUID\":\"723709fc-182d-9ce9-6860-1ce91ccb69c2\",\"ClaimName\":\"example-1\",\"Namespace\":\"default\",\"PodUIDs\":{\"964efe03-c612-2d71-0389-bf00aad506a1\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-1\",\"RequestNames\":[\"test request\"],\"CDIDeviceIDs\":[\"example.com/example=cdi-example\"]}]}}},{\"ClaimUID\":\"3fa3db0a-32a7-8d67-3f61-e75a82b8c87f\",\"ClaimName\":\"example-2\",\"Namespace\":\"default\",\"PodUIDs\":{\"998bcc06-a257-3f37-0846-cc61fcd899e9\":{}},\"DriverState\":{\"test-driver.cdi.k8s.io\":{\"Devices\":[{\"PoolName\":\"worker-1\",\"DeviceName\":\"dev-2\",\"RequestNames\":null,\"CDIDeviceIDs\":null}]}}}]}","Checksum":1582256999}`,
		},
	}

	// Should return an error, stateDir cannot be an empty string
	if _, err := NewCheckpointer("", testingCheckpoint); err == nil {
		t.Fatal("expected error but got nil")
	}

	// create temp dir
	testingDir, err := os.MkdirTemp("", "dramanager_state_test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(testingDir); err != nil {
			t.Fatal(err)
		}
	}()

	// NewCheckpointState with an empty checkpointName should return an error
	if _, err = NewCheckpointer(testingDir, ""); err == nil {
		t.Fatal("expected error but got nil")
	}

	cpm, err := checkpointmanager.NewCheckpointManager(testingDir)
	require.NoError(t, err, "could not create testing checkpoint manager")
	require.NoError(t, cpm.RemoveCheckpoint(testingCheckpoint), "could not remove testing checkpoint")

	cs, err := NewCheckpointer(testingDir, testingCheckpoint)
	require.NoError(t, err, "could not create testing checkpointState instance")

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			checkpoint, err := NewCheckpoint(tc.claimInfoStateList)
			require.NoError(t, err, "could not create Checkpoint")

			err = cs.Store(checkpoint)
			require.NoError(t, err, "could not store checkpoint")

			err = cpm.GetCheckpoint(testingCheckpoint, checkpoint)
			require.NoError(t, err, "could not get checkpoint")

			checkpointContent, err := checkpoint.MarshalCheckpoint()
			require.NoError(t, err, "could not Marshal Checkpoint")
			assert.Equal(t, tc.expectedCheckpointContent, string(checkpointContent))
		})
	}
}
