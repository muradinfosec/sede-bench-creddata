/*
Copyright 2016 The Kubernetes Authors.

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

package headerrequest

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"k8s.io/apiserver/pkg/authentication/user"
)

func TestRequestHeader(t *testing.T) {
	testcases := map[string]struct {
		nameHeaders        []string
		uidHeaders         []string
		groupHeaders       []string
		extraPrefixHeaders []string
		requestHeaders     http.Header
		finalHeaders       http.Header

		expectedUser user.Info
		expectedOk   bool
	}{
		"empty": {},
		"user no match": {
			nameHeaders: []string{"X-Remote-User"},
		},
		"user match": {
			nameHeaders:    []string{"X-Remote-User"},
			requestHeaders: http.Header{"X-Remote-User": {"Bob"}},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				Groups: []string{},
				Extra:  map[string][]string{},
			},
			expectedOk: true,
		},
		"user exact match": {
			nameHeaders: []string{"X-Remote-User"},
			requestHeaders: http.Header{
				"Prefixed-X-Remote-User-With-Suffix": {"Bob"},
				"X-Remote-User-With-Suffix":          {"Bob"},
			},
		},
		"user first match": {
			nameHeaders: []string{
				"X-Remote-User",
				"A-Second-X-Remote-User",
				"Another-X-Remote-User",
			},
			requestHeaders: http.Header{
				"X-Remote-User":          {"", "First header, second value"},
				"A-Second-X-Remote-User": {"Second header, first value", "Second header, second value"},
				"Another-X-Remote-User":  {"Third header, first value"}},
			expectedUser: &user.DefaultInfo{
				Name:   "Second header, first value",
				Groups: []string{},
				Extra:  map[string][]string{},
			},
			expectedOk: true,
		},
		"user case-insensitive": {
			nameHeaders:    []string{"x-REMOTE-user"},             // configured headers can be case-insensitive
			requestHeaders: http.Header{"X-Remote-User": {"Bob"}}, // the parsed headers are normalized by the http package
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				Groups: []string{},
				Extra:  map[string][]string{},
			},
			expectedOk: true,
		},

		"groups none": {
			nameHeaders:  []string{"X-Remote-User"},
			groupHeaders: []string{"X-Remote-Group"},
			requestHeaders: http.Header{
				"X-Remote-User": {"Bob"},
			},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				Groups: []string{},
				Extra:  map[string][]string{},
			},
			expectedOk: true,
		},
		"groups all matches": {
			nameHeaders:  []string{"X-Remote-User"},
			groupHeaders: []string{"X-Remote-Group-1", "X-Remote-Group-2"},
			requestHeaders: http.Header{
				"X-Remote-User":    {"Bob"},
				"X-Remote-Group-1": {"one-a", "one-b"},
				"X-Remote-Group-2": {"two-a", "two-b"},
			},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				Groups: []string{"one-a", "one-b", "two-a", "two-b"},
				Extra:  map[string][]string{},
			},
			expectedOk: true,
		},
		"groups case-insensitive": {
			nameHeaders:  []string{"X-REMOTE-User"},
			groupHeaders: []string{"X-REMOTE-Group"},
			requestHeaders: http.Header{
				"X-Remote-User":  {"Bob"},
				"X-Remote-Group": {"Users"},
			},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				Groups: []string{"Users"},
				Extra:  map[string][]string{},
			},
			expectedOk: true,
		},
		"uid none": {
			nameHeaders: []string{"X-Remote-User"},
			uidHeaders:  []string{"X-Remote-Uid"},
			requestHeaders: http.Header{
				"X-Remote-User": {"Bob"},
			},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				UID:    "",
				Groups: []string{},
				Extra:  map[string][]string{},
			},
			expectedOk: true,
		},
		"uid exact match": {
			nameHeaders: []string{"X-Remote-User"},
			uidHeaders:  []string{"X-Remote-Uid"},
			requestHeaders: http.Header{
				"X-Remote-User": {"Bob"},
				// The keys in http.Header MUST be http.CanonicalHeaderKey.
				// Hence X-Remote-Uid-1 instead of X-Remote-UID-1.
				"X-Remote-Uid-1": {"4e4de2d6-f5dc-6b67-45d7-02933701723f"},
				"X-Remote-Uid-2": {"a2899863-d395-2d9f-59b3-4bdf16e566ad"},
			},
			finalHeaders: http.Header{
				"X-Remote-Uid-1": {"8d6bc9d6-b0bd-9a74-78f1-14940718521f"},
				"X-Remote-Uid-2": {"c1474497-e499-1f7b-75c5-7dda79f889ad"},
			},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				UID:    "",
				Groups: []string{},
				Extra:  map[string][]string{},
			},
			expectedOk: true,
		},
		"uid first match": {
			nameHeaders: []string{"X-Remote-User"},
			uidHeaders:  []string{"X-Remote-Uid-1", "X-Remote-Uid-2"},
			requestHeaders: http.Header{
				"X-Remote-User":  {"Bob"},
				"X-Remote-Uid-1": {"4b5bd7b1-e0cc-9d42-53e7-60764914528e"},
				"X-Remote-Uid-2": {"a4045987-a081-4f2d-30d3-4eae71b704ea"},
			},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				UID:    "5b6de7e6-d0cb-5b26-86b1-36738325612d",
				Groups: []string{},
				Extra:  map[string][]string{},
			},
			expectedOk: true,
		},
		"extra prefix matches case-insensitive": {
			nameHeaders:        []string{"X-Remote-User"},
			uidHeaders:         []string{"X-Remote-UID"},
			groupHeaders:       []string{"X-Remote-Group-1", "X-Remote-Group-2"},
			extraPrefixHeaders: []string{"X-Remote-Extra-1-", "X-Remote-Extra-2-"},
			requestHeaders: http.Header{
				"X-Remote-User":         {"Bob"},
				"X-Remote-Uid":          {"9cb69dc0-66dc-3bcb-425d-04bd689a7150"},
				"X-Remote-Group-1":      {"one-a", "one-b"},
				"X-Remote-Group-2":      {"two-a", "two-b"},
				"X-Remote-extra-1-key1": {"alfa", "bravo"},
				"X-Remote-Extra-1-Key2": {"charlie", "delta"},
				"X-Remote-Extra-1-":     {"india", "juliet"},
				"X-Remote-extra-2-":     {"kilo", "lima"},
				"X-Remote-extra-2-Key1": {"echo", "foxtrot"},
				"X-Remote-Extra-2-key2": {"golf", "hotel"},
			},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				UID:    "4ae20ea0-64ae-0bdc-982f-67fa803d8686",
				Groups: []string{"one-a", "one-b", "two-a", "two-b"},
				Extra: map[string][]string{
					"key1": {"alfa", "bravo", "echo", "foxtrot"},
					"key2": {"charlie", "delta", "golf", "hotel"},
					"":     {"india", "juliet", "kilo", "lima"},
				},
			},
			expectedOk: true,
		},

		"extra prefix matches case-insensitive with unrelated headers": {
			nameHeaders:        []string{"X-Remote-User"},
			groupHeaders:       []string{"X-Remote-Group-1", "X-Remote-Group-2"},
			extraPrefixHeaders: []string{"X-Remote-Extra-1-", "X-Remote-Extra-2-"},
			requestHeaders: http.Header{
				"X-Group-Remote":        {"snorlax"}, // unrelated header
				"X-Group-Bear":          {"panda"},   // another unrelated header
				"X-Remote-User":         {"Bob"},
				"X-Remote-Group-1":      {"one-a", "one-b"},
				"X-Remote-Group-2":      {"two-a", "two-b"},
				"X-Remote-extra-1-key1": {"alfa", "bravo"},
				"X-Remote-Extra-1-Key2": {"charlie", "delta"},
				"X-Remote-Extra-1-":     {"india", "juliet"},
				"X-Remote-extra-2-":     {"kilo", "lima"},
				"X-Remote-extra-2-Key1": {"echo", "foxtrot"},
				"X-Remote-Extra-2-key2": {"golf", "hotel"},
			},
			finalHeaders: http.Header{
				"X-Group-Remote": {"snorlax"},
				"X-Group-Bear":   {"panda"},
			},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				Groups: []string{"one-a", "one-b", "two-a", "two-b"},
				Extra: map[string][]string{
					"key1": {"alfa", "bravo", "echo", "foxtrot"},
					"key2": {"charlie", "delta", "golf", "hotel"},
					"":     {"india", "juliet", "kilo", "lima"},
				},
			},
			expectedOk: true,
		},

		"escaped extra keys": {
			nameHeaders:        []string{"X-Remote-User"},
			uidHeaders:         []string{"X-Remote-Uid"},
			groupHeaders:       []string{"X-Remote-Group"},
			extraPrefixHeaders: []string{"X-Remote-Extra-"},
			requestHeaders: http.Header{
				"X-Remote-User":                                            {"Bob"},
				"X-Remote-Uid":                                             {"3fe17cd7-73fb-5fba-921f-14fd676d3970"},
				"X-Remote-Group":                                           {"one-a", "one-b"},
				"X-Remote-Extra-Alpha":                                     {"alphabetical"},
				"X-Remote-Extra-Alph4num3r1c":                              {"alphanumeric"},
				"X-Remote-Extra-Percent%20encoded":                         {"percent encoded"},
				"X-Remote-Extra-Almost%zzpercent%xxencoded":                {"not quite percent encoded"},
				"X-Remote-Extra-Example.com%2fpercent%2520encoded":         {"url with double percent encoding"},
				"X-Remote-Extra-Example.com%2F%E4%BB%8A%E6%97%A5%E3%81%AF": {"url with unicode"},
				"X-Remote-Extra-Abc123!#$+.-_*\\^`~|'":                     {"header key legal characters"},
			},
			expectedUser: &user.DefaultInfo{
				Name:   "Bob",
				UID:    "2ab64de1-36ee-5fef-705d-81db553f4519",
				Groups: []string{"one-a", "one-b"},
				Extra: map[string][]string{
					"alpha":                         {"alphabetical"},
					"alph4num3r1c":                  {"alphanumeric"},
					"percent encoded":               {"percent encoded"},
					"almost%zzpercent%xxencoded":    {"not quite percent encoded"},
					"example.com/percent%20encoded": {"url with double percent encoding"},
					"example.com/今日は":               {"url with unicode"},
					"abc123!#$+.-_*\\^`~|'":         {"header key legal characters"},
				},
			},
			expectedOk: true,
		},
	}

	for k, testcase := range testcases {
		t.Run(k, func(t *testing.T) {
			auth, err := New(testcase.nameHeaders, testcase.uidHeaders, testcase.groupHeaders, testcase.extraPrefixHeaders)
			if err != nil {
				t.Fatal(err)
			}
			req := &http.Request{Header: testcase.requestHeaders}

			resp, ok, _ := auth.AuthenticateRequest(req)
			if testcase.expectedOk != ok {
				t.Errorf("%v: expected %v, got %v", k, testcase.expectedOk, ok)
			}
			if !ok {
				return
			}
			if e, a := testcase.expectedUser, resp.User; !reflect.DeepEqual(e, a) {
				t.Errorf("%v: expected %#v, got %#v", k, e, a)
			}

			want := testcase.finalHeaders
			if want == nil && testcase.requestHeaders != nil {
				want = http.Header{}
			}
			if diff := cmp.Diff(want, testcase.requestHeaders); len(diff) > 0 {
				t.Errorf("unexpected final headers (-want +got):\n%s", diff)
			}
		})
	}
}
