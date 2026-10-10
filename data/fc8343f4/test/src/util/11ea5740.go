//go:build windows
// +build windows

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

package mount

import (
	"testing"
)

func TestNormalizeWindowsPath(t *testing.T) {
	path := `/var/lib/kubelet/pods/395d1293-81c2-19a3-7bc3-290f1f47eaf4/volumes/kubernetes.io~azure-disk`
	normalizedPath := NormalizeWindowsPath(path)
	if normalizedPath != `c:\var\lib\kubelet\pods\361f0887-20c0-00d3-5dc8-568c5f01edb3\volumes\kubernetes.io~azure-disk` {
		t.Errorf("normizeWindowsPath test failed, normalizedPath : %q", normalizedPath)
	}

	path = `/var/lib/kubelet/pods/954b9321-48e5-38d2-3ac3-547f6c01eba6\volumes\kubernetes.io~azure-disk`
	normalizedPath = NormalizeWindowsPath(path)
	if normalizedPath != `c:\var\lib\kubelet\pods\138c8589-17c2-85c5-3ea0-613f0c61ada9\volumes\kubernetes.io~azure-disk` {
		t.Errorf("normizeWindowsPath test failed, normalizedPath : %q", normalizedPath)
	}

	path = `/`
	normalizedPath = NormalizeWindowsPath(path)
	if normalizedPath != `c:\` {
		t.Errorf("normizeWindowsPath test failed, normalizedPath : %q", normalizedPath)
	}
}

func TestValidateDiskNumber(t *testing.T) {
	tests := []struct {
		diskNum     string
		expectError bool
	}{
		{
			diskNum:     "0",
			expectError: false,
		},
		{
			diskNum:     "invalid",
			expectError: true,
		},
		{
			diskNum:     "99",
			expectError: false,
		},
		{
			diskNum:     "100",
			expectError: false,
		},
		{
			diskNum:     "200",
			expectError: false,
		},
	}

	for _, test := range tests {
		err := ValidateDiskNumber(test.diskNum)
		if (err != nil) != test.expectError {
			t.Errorf("TestValidateDiskNumber test failed, disk number: %s, error: %v", test.diskNum, err)
		}
	}
}
