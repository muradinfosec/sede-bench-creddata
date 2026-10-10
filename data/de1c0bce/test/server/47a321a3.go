// +build !js

package webrtc

import (
	"testing"

	"github.com/pion/ice/v2"
	"github.com/pion/webrtc/v3/pkg/rtcerr"
	"github.com/stretchr/testify/assert"
)

func TestICEServer_validate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		testCases := []struct {
			iceServer        ICEServer
			expectedValidate bool
		}{
			{ICEServer{
				URLs:           []string{"turn:192.158.29.39?transport=udp"},
				Username:       "unittest",
				Credential:     "placeholder",
				CredentialType: ICECredentialTypePassword,
			}, true},
			{ICEServer{
				URLs:           []string{"turn:[2001:db8:1234:5678::1]?transport=udp"},
				Username:       "unittest",
				Credential:     "placeholder",
				CredentialType: ICECredentialTypePassword,
			}, true},
			{ICEServer{
				URLs:     []string{"turn:192.158.29.39?transport=udp"},
				Username: "unittest",
				Credential: OAuthCredential{
					MACKey:      "CwqgzgA4WC9qtTeagt45MuAqLI7=",
					AccessToken: "KMiv7uPBZKwnf9iGZT791hWiupcaTUWhB0VsXBhtBV1+RkF60abs870uE3FfeFk5ku46KGZOPeQufvklAkEDmK==",
				},
				CredentialType: ICECredentialTypeOauth,
			}, true},
		}

		for i, testCase := range testCases {
			_, err := testCase.iceServer.urls()
			assert.Nil(t, err, "testCase: %d %v", i, testCase)
		}
	})
	t.Run("Failure", func(t *testing.T) {
		testCases := []struct {
			iceServer   ICEServer
			expectedErr error
		}{
			{ICEServer{
				URLs: []string{"turn:192.158.29.39?transport=udp"},
			}, &rtcerr.InvalidAccessError{Err: ErrNoTurnCredentials}},
			{ICEServer{
				URLs:           []string{"turn:192.158.29.39?transport=udp"},
				Username:       "unittest",
				Credential:     false,
				CredentialType: ICECredentialTypePassword,
			}, &rtcerr.InvalidAccessError{Err: ErrTurnCredentials}},
			{ICEServer{
				URLs:           []string{"turn:192.158.29.39?transport=udp"},
				Username:       "unittest",
				Credential:     false,
				CredentialType: ICECredentialTypeOauth,
			}, &rtcerr.InvalidAccessError{Err: ErrTurnCredentials}},
			{ICEServer{
				URLs:           []string{"turn:192.158.29.39?transport=udp"},
				Username:       "unittest",
				Credential:     false,
				CredentialType: Unknown,
			}, &rtcerr.InvalidAccessError{Err: ErrTurnCredentials}},
			{ICEServer{
				URLs:           []string{"stun:google.de?transport=udp"},
				Username:       "unittest",
				Credential:     false,
				CredentialType: ICECredentialTypeOauth,
			}, &rtcerr.InvalidAccessError{Err: ice.ErrSTUNQuery}},
		}

		for i, testCase := range testCases {
			_, err := testCase.iceServer.urls()
			assert.EqualError(t,
				err,
				testCase.expectedErr.Error(),
				"testCase: %d %v", i, testCase,
			)
		}
	})
}
