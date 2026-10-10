/*
Copyright 2015 The Kubernetes Authors.

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

package transport

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"testing"
)

const (
	rootCACert = `-----BEGIN CERTIFICATE-----
MIIC4DCCAcqgAwIBAgIBATALBgkqhkiG9w0BAQswIzEhMB8GA1UEAwwYMTAuMTMu
MTI5LjEwNkAxNDIxMzU5MDU4MB4XDTE1MDExNTIxNTczN1oXDTE2MDExNTIxNTcz
OFowIzEhMB8GA1UEAwwYMTAuMTMuMTI5LjEwNkAxNDIxMzU5MDU4MIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAunDRXGwsiYWGFDlWH6kjGun+PshDGeZX
xtx9lUnL8pIRWH3wX6f13PO9sktaOWW0T0mlo6k2bMlSLlSZgG9H6og0W6gLS3vq
s4VavZ6DbXIwemZG2vbRwsvR+t4G6Nbwelm6F8RFnA1Fwt428pavmNQ/wgYzo+T1
1eS+HiN4ACnSoDSx3QRWcgBkB1g6VReofVjx63i0J+w8Q/41L9GUuLqquFxu6ZnH
60vTB55lHgFiDLjA1FkEz2dGvGh/wtnFlRvjaPC54JH2K1mPYAUXTreoeJtLJKX0
ycoiyB24+zGCniUmgIsmQWRPaOPircexCp1BOeze82BT1LCZNTVaxQIDAQABoyMw
ITAOBgNVHQ8BAf8EBAMCAKQwDwYDVR0TAQH/BAUwAwEB/zALBgkqhkiG9w0BAQsD
ggEBADMxsUuAFlsYDpF4fRCzXXwrhbtj4oQwcHpbu+rnOPHCZupiafzZpDu+rw4x
YGPnCb594bRTQn4pAu3Ac18NbLD5pV3uioAkv8oPkgr8aUhXqiv7KdDiaWm6sbAL
EHiXVBBAFvQws10HMqMoKtO8f1XDNAUkWduakR/U6yMgvOPwS7xl0eUTqyRB6zGb
K55q2dejiFWaFqB/y78txzvz6UlOZKE44g2JAVoJVM6kGaxh33q8/FmrL4kuN3ut
W+MmJCVDvd4eEqPwbp7146ZWTqpIJ8lvA6wuChtqV8lhAPka2hD/LMqY8iXNmfXD
uml0obOEy+ON91k+SWTJ3ggmF/U=
-----END CERTIFICATE-----`

	certData = `-----BEGIN CERTIFICATE-----
MIIC6jCCAdSgAwIBAgIBCzALBgkqhkiG9w0BAQswIzEhMB8GA1UEAwwYMTAuMTMu
MTI5LjEwNkAxNDIxMzU5MDU4MB4XDTE1MDExNTIyMDEzMVoXDTE2MDExNTIyMDEz
MlowGzEZMBcGA1UEAxMQb3BlbnNoaWZ0LWNsaWVudDCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAKtdhz0+uCLXw5cSYns9rU/XifFSpb/x24WDdrm72S/v
b9BPYsAStiP148buylr1SOuNi8sTAZmlVDDIpIVwMLff+o2rKYDicn9fjbrTxTOj
lI4pHJBH+JU3AJ0tbajupioh70jwFS0oYpwtneg2zcnE2Z4l6mhrj2okrc5Q1/X2
I2HChtIU4JYTisObtin10QKJX01CLfYXJLa8upWzKZ4/GOcHG+eAV3jXWoXidtjb
1Usw70amoTZ6mIVCkiu1QwCoa8+ycojGfZhvqMsAp1536ZcCul+Na+AbCv4zKS7F
kQQaImVrXdUiFansIoofGlw/JNuoKK6ssVpS5Ic3pgcCAwEAAaM1MDMwDgYDVR0P
AQH/BAQDAgCgMBMGA1UdJQQMMAoGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwCwYJ
KoZIhvcNAQELA4IBAQCKLREH7bXtXtZ+8vI6cjD7W3QikiArGqbl36bAhhWsJLp/
p/ndKz39iFNaiZ3GlwIURWOOKx3y3GA0x9m8FR+Llthf0EQ8sUjnwaknWs0Y6DQ3
jjPFZOpV3KPCFrdMJ3++E3MgwFC/Ih/N2ebFX9EcV9Vcc6oVWMdwT0fsrhu683rq
6GSR/3iVX1G/pmOiuaR0fNUaCyCfYrnI4zHBDgSfnlm3vIvN2lrsR/DQBakNL8DJ
HBgKxMGeUPoneBv+c8DMXIL0EhaFXRlBv9QW45/GiAIOuyFJ0i6hCtGZpJjq4OpQ
BRjCI+izPzFTjsxD4aORE+WOkyWFCGPWKfNejfw0
-----END CERTIFICATE-----`

	keyData = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAq12HPT64ItfDlxJiez2tT9eJ8VKlv/HbhYN2ubvZL+9v0E9i
KNA2b/IWBg6jXoHo067wYcXHmdpzkNAzAdYxD2/8aPBogjbPr5+ZWWbIo3CgvxQG
glv4nIEpnv7hEj5WpLbfFUpqJALNeu7r0SPBTICQswQKCliUEexhnsCg7vzKpGwn
4pfApsWfS6S4tljATgqQYQzX9QKkCDs5Zjmllr2I3WoJ55GDOsUvpVd86pgALCyk
gFKTtCODeccwB1gTDHwiU9yljkO3tN+bUVCpfyGGPWA4B16W3vQy/LgBPUlbfDGE
NUgq3bmvbbLYwp4wUS3k94ruLThrtjrIMSBXfwhPUhMcJdUvyX8kGnb6ylIugcVn
ESFXmVkvOeLS3GEPIDgXLrV1mEX1Zlkj6668YtdnNVZbTSPUUHQ+wTySlSaUOCjM
QN28roc4p5amha4mwRPoBZmO8j2lQ7rSN3ss6jkUqCSa9O2A8m0rbN1z59aeWAzE
9qzXWPBH2GR20iJpotRCw8APAcukU3kVW5iJ6JWeJNGrEu2peDQx0IGYMT7XAq9S
UIMoZNK7J3rF6FJMqmaLM5/1Es4nMfWq3XnseEgmX+OaJetAoGby2PvRCJK5BgQ9
CRnkyukaHsn6kEFD6+rI5AsEwDVEplW8AsSADor10lEECAK+5lpDazJSnoxnKrQg
2l3ZSgzMgcVC4EnTjSbrQADSUkhYBD4TpLfayl09altqj2L2nz6cvlcg9zW9LWgq
xg24Wa6hjvS96jZDwCGCV7LJ8FOb/KJ87P8LoASCjXHWBixjsUkebXrutgugynDD
tsSXKRduA6PadtVcUeOrYsa1uAx+m/FOEhEYHIXH0IaL2JO+AEaDDf1sKIfIHmW/
FmdRM9sisSmWUqQVceKTZOh14l1tYLAdzkLWSunWHwVKM5mmnUF2IvAV67DFcGB8
4mOsQwGdLFmh+RCCSkmzNzuqIaf3N25urffajc/wZ1f8ksNyCwr0B8k1Kv/tvRNi
ApaVSd9DsZkNE7WLZ8LcX+CydDPiJa2lEOqMn/nYaot7/sYIpaD6Q7/VqbeZZ54C
7yEmQ+t3T+KTn9WxaBMRRXWamPWIRcGnJrTxu0CFxiED7UjboFM/oSNWmS2c/TtK
hLtVMI/tlGiapwXtucT6ENNiVfqU33rgGleX/7mVdi7lp/TMh3eHalnmFxdpMaL+
MMblQp0XA0Qngkp8XOwYxdcqVVt5vvZLfH8rZp9SfMMptIL/ISOkUDDu6qRMH60h
8Lb1So2RnjpGeE0Ul5SBpI85CiyNt4pqDIuwqqysh2L+tbi7YqhwDKi+kSPd4SFi
50RZfZuwhG6pyfCAHm2MCMtu7HoNCs+m9d8E59KBPBvUY7KynPbiYCQY3AvExyf/
yjJZzDUylCOqlsENKL+2llCXdTHOgquedVj1aWNeWOz7af3JZW+8IfuIV0JDF5h/
Y6LYbqWdk7jmI/CLETxQXA1jyQckGTmv0pMtCsdr2lGa5FRwwiQjAj4y0TxbR2ot
KoK9sfG1CYP8EE/IZGT6ZQo0IjH62TDU+kVJxd6jqtuaCHq8MtkE
-----END RSA PRIVATE KEY-----`
)

func TestNew(t *testing.T) {
	globalGetCert := &GetCertHolder{
		GetCert: func() (*tls.Certificate, error) { return nil, nil },
	}
	globalDial := &DialHolder{
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) { return nil, nil },
	}

	testCases := map[string]struct {
		Config       *Config
		Err          bool
		TLS          bool
		TLSCert      bool
		TLSErr       bool
		Default      bool
		Insecure     bool
		DefaultRoots bool
	}{
		"default transport": {
			Default: true,
			Config:  &Config{},
		},

		"insecure": {
			TLS:          true,
			Insecure:     true,
			DefaultRoots: true,
			Config: &Config{TLS: TLSConfig{
				Insecure: true,
			}},
		},

		"server name": {
			TLS:          true,
			DefaultRoots: true,
			Config: &Config{TLS: TLSConfig{
				ServerName: "foo",
			}},
		},

		"ca transport": {
			TLS: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData: []byte(rootCACert),
				},
			},
		},
		"bad ca file transport": {
			Err: true,
			Config: &Config{
				TLS: TLSConfig{
					CAFile: "invalid file",
				},
			},
		},
		"bad ca data transport": {
			Err: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData: []byte(rootCACert + "this is not valid"),
				},
			},
		},
		"ca data overriding bad ca file transport": {
			TLS: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData: []byte(rootCACert),
					CAFile: "invalid file",
				},
			},
		},

		"cert transport": {
			TLS:     true,
			TLSCert: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData:   []byte(rootCACert),
					CertData: []byte(certData),
					KeyData:  []byte(keyData),
				},
			},
		},
		"bad cert data transport": {
			Err: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData:   []byte(rootCACert),
					CertData: []byte(certData),
					KeyData:  []byte("bad key data"),
				},
			},
		},
		"bad file cert transport": {
			Err: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData:   []byte(rootCACert),
					CertData: []byte(certData),
					KeyFile:  "invalid file",
				},
			},
		},
		"key data overriding bad file cert transport": {
			TLS:     true,
			TLSCert: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData:   []byte(rootCACert),
					CertData: []byte(certData),
					KeyData:  []byte(keyData),
					KeyFile:  "invalid file",
				},
			},
		},
		"callback cert and key": {
			TLS:     true,
			TLSCert: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData: []byte(rootCACert),
					GetCertHolder: &GetCertHolder{
						GetCert: func() (*tls.Certificate, error) {
							crt, err := tls.X509KeyPair([]byte(certData), []byte(keyData))
							return &crt, err
						},
					},
				},
			},
		},
		"cert callback error": {
			TLS:     true,
			TLSCert: true,
			TLSErr:  true,
			Config: &Config{
				TLS: TLSConfig{
					CAData: []byte(rootCACert),
					GetCertHolder: &GetCertHolder{
						GetCert: func() (*tls.Certificate, error) {
							return nil, errors.New("GetCert failure")
						},
					},
				},
			},
		},
		"cert data overrides empty callback result": {
			TLS:     true,
			TLSCert: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData: []byte(rootCACert),
					GetCertHolder: &GetCertHolder{
						GetCert: func() (*tls.Certificate, error) {
							return nil, nil
						},
					},
					CertData: []byte(certData),
					KeyData:  []byte(keyData),
				},
			},
		},
		"callback returns nothing": {
			TLS:     true,
			TLSCert: true,
			Config: &Config{
				TLS: TLSConfig{
					CAData: []byte(rootCACert),
					GetCertHolder: &GetCertHolder{
						GetCert: func() (*tls.Certificate, error) {
							return nil, nil
						},
					},
				},
			},
		},
		"nil holders": {
			Config: &Config{
				TLS: TLSConfig{
					GetCertHolder: nil,
				},
				DialHolder: nil,
			},
			Err:          false,
			TLS:          false,
			TLSCert:      false,
			TLSErr:       false,
			Default:      true,
			Insecure:     false,
			DefaultRoots: false,
		},
		"non-nil dial holder and nil internal": {
			Config: &Config{
				TLS: TLSConfig{
					GetCertHolder: nil,
				},
				DialHolder: &DialHolder{},
			},
			Err: true,
		},
		"non-nil cert holder and nil internal": {
			Config: &Config{
				TLS: TLSConfig{
					GetCertHolder: &GetCertHolder{},
				},
				DialHolder: nil,
			},
			Err: true,
		},
		"non-nil dial holder+internal": {
			Config: &Config{
				TLS: TLSConfig{
					GetCertHolder: nil,
				},
				DialHolder: &DialHolder{
					Dial: func(ctx context.Context, network, address string) (net.Conn, error) { return nil, nil },
				},
			},
			Err:          false,
			TLS:          true,
			TLSCert:      false,
			TLSErr:       false,
			Default:      false,
			Insecure:     false,
			DefaultRoots: true,
		},
		"non-nil cert holder+internal": {
			Config: &Config{
				TLS: TLSConfig{
					GetCertHolder: &GetCertHolder{
						GetCert: func() (*tls.Certificate, error) { return nil, nil },
					},
				},
				DialHolder: nil,
			},
			Err:          false,
			TLS:          true,
			TLSCert:      true,
			TLSErr:       false,
			Default:      false,
			Insecure:     false,
			DefaultRoots: true,
		},
		"non-nil holders+internal with global address": {
			Config: &Config{
				TLS: TLSConfig{
					GetCertHolder: globalGetCert,
				},
				DialHolder: globalDial,
			},
			Err:          false,
			TLS:          true,
			TLSCert:      true,
			TLSErr:       false,
			Default:      false,
			Insecure:     false,
			DefaultRoots: true,
		},
	}
	for k, testCase := range testCases {
		t.Run(k, func(t *testing.T) {
			rt, err := New(testCase.Config)
			switch {
			case testCase.Err && err == nil:
				t.Fatal("unexpected non-error")
			case !testCase.Err && err != nil:
				t.Fatalf("unexpected error: %v", err)
			}
			if testCase.Err {
				return
			}

			switch {
			case testCase.Default && rt != http.DefaultTransport:
				t.Fatalf("got %#v, expected the default transport", rt)
			case !testCase.Default && rt == http.DefaultTransport:
				t.Fatalf("got %#v, expected non-default transport", rt)
			}

			// We only know how to check TLSConfig on http.Transports
			transport := rt.(*http.Transport)
			switch {
			case testCase.TLS && transport.TLSClientConfig == nil:
				t.Fatalf("got %#v, expected TLSClientConfig", transport)
			case !testCase.TLS && transport.TLSClientConfig != nil:
				t.Fatalf("got %#v, expected no TLSClientConfig", transport)
			}
			if !testCase.TLS {
				return
			}

			switch {
			case testCase.DefaultRoots && transport.TLSClientConfig.RootCAs != nil:
				t.Fatalf("got %#v, expected nil root CAs", transport.TLSClientConfig.RootCAs)
			case !testCase.DefaultRoots && transport.TLSClientConfig.RootCAs == nil:
				t.Fatalf("got %#v, expected non-nil root CAs", transport.TLSClientConfig.RootCAs)
			}

			switch {
			case testCase.Insecure != transport.TLSClientConfig.InsecureSkipVerify:
				t.Fatalf("got %#v, expected %#v", transport.TLSClientConfig.InsecureSkipVerify, testCase.Insecure)
			}

			switch {
			case testCase.TLSCert && transport.TLSClientConfig.GetClientCertificate == nil:
				t.Fatalf("got %#v, expected TLSClientConfig.GetClientCertificate", transport.TLSClientConfig)
			case !testCase.TLSCert && transport.TLSClientConfig.GetClientCertificate != nil:
				t.Fatalf("got %#v, expected no TLSClientConfig.GetClientCertificate", transport.TLSClientConfig)
			}
			if !testCase.TLSCert {
				return
			}

			_, err = transport.TLSClientConfig.GetClientCertificate(nil)
			switch {
			case testCase.TLSErr && err == nil:
				t.Error("got nil error from GetClientCertificate, expected non-nil")
			case !testCase.TLSErr && err != nil:
				t.Errorf("got error from GetClientCertificate: %q, expected nil", err)
			}
		})
	}
}

type fakeRoundTripper struct {
	Req  *http.Request
	Resp *http.Response
	Err  error
}

func (rt *fakeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.Req = req
	return rt.Resp, rt.Err
}

type chainRoundTripper struct {
	rt    http.RoundTripper
	value string
}

func testChain(value string) WrapperFunc {
	return func(rt http.RoundTripper) http.RoundTripper {
		return &chainRoundTripper{rt: rt, value: value}
	}
}

func (rt *chainRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := rt.rt.RoundTrip(req)
	if resp != nil {
		if resp.Header == nil {
			resp.Header = make(http.Header)
		}
		resp.Header.Set("Value", resp.Header.Get("Value")+rt.value)
	}
	return resp, err
}

func TestWrappers(t *testing.T) {
	resp1 := &http.Response{}
	wrapperResp1 := func(rt http.RoundTripper) http.RoundTripper {
		return &fakeRoundTripper{Resp: resp1}
	}
	resp2 := &http.Response{}
	wrapperResp2 := func(rt http.RoundTripper) http.RoundTripper {
		return &fakeRoundTripper{Resp: resp2}
	}

	tests := []struct {
		name    string
		fns     []WrapperFunc
		wantNil bool
		want    func(*http.Response) bool
	}{
		{fns: []WrapperFunc{}, wantNil: true},
		{fns: []WrapperFunc{nil, nil}, wantNil: true},
		{fns: []WrapperFunc{nil}, wantNil: false},

		{fns: []WrapperFunc{nil, wrapperResp1}, want: func(resp *http.Response) bool { return resp == resp1 }},
		{fns: []WrapperFunc{wrapperResp1, nil}, want: func(resp *http.Response) bool { return resp == resp1 }},
		{fns: []WrapperFunc{nil, wrapperResp1, nil}, want: func(resp *http.Response) bool { return resp == resp1 }},
		{fns: []WrapperFunc{nil, wrapperResp1, wrapperResp2}, want: func(resp *http.Response) bool { return resp == resp2 }},
		{fns: []WrapperFunc{wrapperResp1, wrapperResp2}, want: func(resp *http.Response) bool { return resp == resp2 }},
		{fns: []WrapperFunc{wrapperResp2, wrapperResp1}, want: func(resp *http.Response) bool { return resp == resp1 }},

		{fns: []WrapperFunc{testChain("1")}, want: func(resp *http.Response) bool { return resp.Header.Get("Value") == "1" }},
		{fns: []WrapperFunc{testChain("1"), testChain("2")}, want: func(resp *http.Response) bool { return resp.Header.Get("Value") == "12" }},
		{fns: []WrapperFunc{testChain("2"), testChain("1")}, want: func(resp *http.Response) bool { return resp.Header.Get("Value") == "21" }},
		{fns: []WrapperFunc{testChain("1"), testChain("2"), testChain("3")}, want: func(resp *http.Response) bool { return resp.Header.Get("Value") == "123" }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Wrappers(tt.fns...)
			if got == nil != tt.wantNil {
				t.Errorf("Wrappers() = %v", got)
				return
			}
			if got == nil {
				return
			}

			rt := &fakeRoundTripper{Resp: &http.Response{}}
			nested := got(rt)
			req := &http.Request{}
			resp, _ := nested.RoundTrip(req)
			if tt.want != nil && !tt.want(resp) {
				t.Errorf("unexpected response: %#v", resp)
			}
		})
	}
}

func Test_contextCanceller_RoundTrip(t *testing.T) {
	tests := []struct {
		name string
		open bool
		want bool
	}{
		{name: "open context should call nested round tripper", open: true, want: true},
		{name: "closed context should return a known error", open: false, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{}
			rt := &fakeRoundTripper{Resp: &http.Response{}}
			ctx := context.Background()
			if !tt.open {
				c, fn := context.WithCancel(ctx)
				fn()
				ctx = c
			}
			errTesting := fmt.Errorf("testing")
			b := &contextCanceller{
				rt:  rt,
				ctx: ctx,
				err: errTesting,
			}
			got, err := b.RoundTrip(req)
			if tt.want {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if got != rt.Resp {
					t.Errorf("wanted response")
				}
				if req != rt.Req {
					t.Errorf("expect nested call")
				}
			} else {
				if err != errTesting {
					t.Errorf("unexpected error: %v", err)
				}
				if got != nil {
					t.Errorf("wanted no response")
				}
				if rt.Req != nil {
					t.Errorf("want no nested call")
				}
			}
		})
	}
}
