/*
Copyright 2023 The Kubernetes Authors.

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

package remotecommand

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/httpstream"
	utilnettesting "k8s.io/apimachinery/pkg/util/net/testing"
	"k8s.io/apimachinery/pkg/util/remotecommand"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
)

func TestFallbackClient_WebSocketPrimarySucceeds(t *testing.T) {
	// Create fake WebSocket server. Copy received STDIN data back onto STDOUT stream.
	websocketServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		conns, err := webSocketServerStreams(req, w, streamOptionsFromRequest(req))
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		defer conns.conn.Close()
		// Loopback the STDIN stream onto the STDOUT stream.
		_, err = io.Copy(conns.stdoutStream, conns.stdinStream)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
	}))
	defer websocketServer.Close()

	// Now create the fallback client (executor), and point it to the "websocketServer".
	// Must add STDIN and STDOUT query params for the client request.
	websocketServer.URL = websocketServer.URL + "?" + "stdin=true" + "&" + "stdout=true"
	websocketLocation, err := url.Parse(websocketServer.URL)
	require.NoError(t, err)
	websocketExecutor, err := NewWebSocketExecutor(&rest.Config{Host: websocketLocation.Host}, "GET", websocketServer.URL)
	require.NoError(t, err)
	spdyExecutor, err := NewSPDYExecutor(&rest.Config{Host: websocketLocation.Host}, "POST", websocketLocation)
	require.NoError(t, err)
	// Never fallback, so always use the websocketExecutor, which succeeds against websocket server.
	exec, err := NewFallbackExecutor(websocketExecutor, spdyExecutor, func(error) bool { return false })
	require.NoError(t, err)
	// Generate random data, and set it up to stream on STDIN. The data will be
	// returned on the STDOUT buffer.
	randomSize := 1024 * 1024
	randomData := make([]byte, randomSize)
	if _, err := rand.Read(randomData); err != nil {
		t.Errorf("unexpected error reading random data: %v", err)
	}
	var stdout bytes.Buffer
	options := &StreamOptions{
		Stdin:  bytes.NewReader(randomData),
		Stdout: &stdout,
	}
	errorChan := make(chan error)
	go func() {
		// Start the streaming on the WebSocket "exec" client.
		errorChan <- exec.StreamWithContext(context.Background(), *options)
	}()

	select {
	case <-time.After(wait.ForeverTestTimeout):
		t.Fatalf("expect stream to be closed after connection is closed.")
	case err := <-errorChan:
		if err != nil {
			t.Errorf("unexpected error")
		}
	}

	data, err := io.ReadAll(bytes.NewReader(stdout.Bytes()))
	if err != nil {
		t.Errorf("error reading the stream: %v", err)
		return
	}
	// Check the random data sent on STDIN was the same returned on STDOUT.
	if !bytes.Equal(randomData, data) {
		t.Errorf("unexpected data received: %d sent: %d", len(data), len(randomData))
	}
}

func TestFallbackClient_SPDYSecondarySucceeds(t *testing.T) {
	// Create fake SPDY server. Copy received STDIN data back onto STDOUT stream.
	spdyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var stdin, stdout bytes.Buffer
		ctx, err := createHTTPStreams(w, req, &StreamOptions{
			Stdin:  &stdin,
			Stdout: &stdout,
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		defer ctx.conn.Close()
		_, err = io.Copy(ctx.stdoutStream, ctx.stdinStream)
		if err != nil {
			t.Fatalf("error copying STDIN to STDOUT: %v", err)
		}
	}))
	defer spdyServer.Close()

	spdyLocation, err := url.Parse(spdyServer.URL)
	require.NoError(t, err)
	websocketExecutor, err := NewWebSocketExecutor(&rest.Config{Host: spdyLocation.Host}, "GET", spdyServer.URL)
	require.NoError(t, err)
	spdyExecutor, err := NewSPDYExecutor(&rest.Config{Host: spdyLocation.Host}, "POST", spdyLocation)
	require.NoError(t, err)
	// Always fallback to spdyExecutor, and spdyExecutor succeeds against fake spdy server.
	exec, err := NewFallbackExecutor(websocketExecutor, spdyExecutor, func(error) bool { return true })
	require.NoError(t, err)
	// Generate random data, and set it up to stream on STDIN. The data will be
	// returned on the STDOUT buffer.
	randomSize := 1024 * 1024
	randomData := make([]byte, randomSize)
	if _, err := rand.Read(randomData); err != nil {
		t.Errorf("unexpected error reading random data: %v", err)
	}
	var stdout bytes.Buffer
	options := &StreamOptions{
		Stdin:  bytes.NewReader(randomData),
		Stdout: &stdout,
	}
	errorChan := make(chan error)
	go func() {
		errorChan <- exec.StreamWithContext(context.Background(), *options)
	}()

	select {
	case <-time.After(wait.ForeverTestTimeout):
		t.Fatalf("expect stream to be closed after connection is closed.")
	case err := <-errorChan:
		if err != nil {
			t.Errorf("unexpected error")
		}
	}

	data, err := io.ReadAll(bytes.NewReader(stdout.Bytes()))
	if err != nil {
		t.Errorf("error reading the stream: %v", err)
		return
	}
	// Check the random data sent on STDIN was the same returned on STDOUT.
	if !bytes.Equal(randomData, data) {
		t.Errorf("unexpected data received: %d sent: %d", len(data), len(randomData))
	}
}

func TestFallbackClient_PrimaryAndSecondaryFail(t *testing.T) {
	// Create fake WebSocket server. Copy received STDIN data back onto STDOUT stream.
	websocketServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		conns, err := webSocketServerStreams(req, w, streamOptionsFromRequest(req))
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		defer conns.conn.Close()
		// Loopback the STDIN stream onto the STDOUT stream.
		_, err = io.Copy(conns.stdoutStream, conns.stdinStream)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
	}))
	defer websocketServer.Close()

	// Now create the fallback client (executor), and point it to the "websocketServer".
	// Must add STDIN and STDOUT query params for the client request.
	websocketServer.URL = websocketServer.URL + "?" + "stdin=true" + "&" + "stdout=true"
	websocketLocation, err := url.Parse(websocketServer.URL)
	require.NoError(t, err)
	websocketExecutor, err := NewWebSocketExecutor(&rest.Config{Host: websocketLocation.Host}, "GET", websocketServer.URL)
	require.NoError(t, err)
	spdyExecutor, err := NewSPDYExecutor(&rest.Config{Host: websocketLocation.Host}, "POST", websocketLocation)
	require.NoError(t, err)
	// Always fallback to spdyExecutor, but spdyExecutor fails against websocket server.
	exec, err := NewFallbackExecutor(websocketExecutor, spdyExecutor, func(error) bool { return true })
	require.NoError(t, err)
	// Update the websocket executor to request remote command v4, which is unsupported.
	fallbackExec, ok := exec.(*FallbackExecutor)
	assert.True(t, ok, "error casting executor as FallbackExecutor")
	websocketExec, ok := fallbackExec.primary.(*wsStreamExecutor)
	assert.True(t, ok, "error casting executor as websocket executor")
	// Set the attempted subprotocol version to V4; websocket server only accepts V5.
	websocketExec.protocols = []string{remotecommand.StreamProtocolV4Name}

	// Generate random data, and set it up to stream on STDIN. The data will be
	// returned on the STDOUT buffer.
	randomSize := 1024 * 1024
	randomData := make([]byte, randomSize)
	if _, err := rand.Read(randomData); err != nil {
		t.Errorf("unexpected error reading random data: %v", err)
	}
	var stdout bytes.Buffer
	options := &StreamOptions{
		Stdin:  bytes.NewReader(randomData),
		Stdout: &stdout,
	}
	errorChan := make(chan error)
	go func() {
		errorChan <- exec.StreamWithContext(context.Background(), *options)
	}()

	select {
	case <-time.After(wait.ForeverTestTimeout):
		t.Fatalf("expect stream to be closed after connection is closed.")
	case err := <-errorChan:
		// Ensure secondary executor returned an error.
		require.Error(t, err)
	}
}

// localhostCert was generated from crypto/tls/generate_cert.go with the following command:
//
//	go run generate_cert.go  --rsa-bits 2048 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var localhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDGTCCAgGgAwIBAgIRALL5AZcefF4kkYV1SEG6YrMwDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAgFw03MDAxMDEwMDAwMDBaGA8yMDg0MDEyOTE2
MDAwMFowEjEQMA4GA1UEChMHQWNtZSBDbzCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBALQ/FHcyVwdFHxARbbD2KBtDUT7Eni+8ioNdjtGcmtXqBv45EC1C
JOqqGJTroFGJ6Q9kQIZ9FqH5IJR2fOOJD9kOTueG4Vt1JY1rj1Kbpjefu8XleZ5L
SBwIWVnN/lEsEbuKmj7N2gLt5AH3zMZiBI1mg1u9Z5ZZHYbCiTpBrwsq6cTlvR9g
dyo1YkM5hRESCzsrL0aUByoo0qRMD8ZsgANJwgsiO0/M6idbxDwv1BnGwGmRYvOE
Hxpy3v0Jg7GJYrvnpnifJTs4nw91N5X9pXxR7FFzi/6HTYDWRljvTb0w6XciKYAz
bWZ0+cJr5F7wB7ovlbm7HrQIR7z7EIIu2d8CAwEAAaNoMGYwDgYDVR0PAQH/BAQD
AgKkMBMGA1UdJQQMMAoGCCsGAQUFBwMBMA8GA1UdEwEB/wQFMAMBAf8wLgYDVR0R
BCcwJYILZXhhbXBsZS5jb22HBH8AAAGHEAAAAAAAAAAAAAAAAAAAAAEwDQYJKoZI
hvcNAQELBQADggEBAFPPWopNEJtIA2VFAQcqN6uJK+JVFOnjGRoCrM6Xgzdm0wxY
XCGjsxY5dl+V7KzdGqu858rCaq5osEBqypBpYAnS9C38VyCDA1vPS1PsN8SYv48z
DyBwj+7R2qar0ADBhnhWxvYO9M72lN/wuCqFKYMeFSnJdQLv3AsrrHe9lYqOa36s
8wxSwVTFTYXBzljPEnSaaJMPqFD8JXaZK1ryJPkO5OsCNQNGtatNiWAf3DcmwHAT
MGYMzP0u4nw47aRz9shB8w+taPKHx2BVwE1m/yp3nHVioOjXqA1fwRQVGclCJSH1
D2iq3hWVHRENgjTjANBPICLo9AZ4JfN6PH19mnU=
-----END CERTIFICATE-----`)

// localhostKey is the private key for localhostCert.
var localhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAtD8UdzJXB0UfEBFtsPYoG0NRPsSeL7yKg12O0Zya1eoG/jkQ
piFI7QEqUKuNtRKkq3qhiq4OGnMDPBE844OY1t2T50uPC3ViiIuYtFYJk3+5XkN6
YrONtMbqdz4+JVccU2QRiC7WwI2hOCkgmPzelvxhY58rSWWonZppkJMBeUsbALg8
v0h9iqGttOQOgbHhYfgKxiraCLMxyaGKfLndZ1kYwRH0a9kBY5yWLP/MfnfpoMnG
74sjHTRe/oQIgSBHl+eZnw7lSmVTi8Q0ZO0JeHWJUhgb/BymsTBzKW8RIaUtBfKT
rQAUGBk9QRTTdpRIQC+zIcdiVnbivLtcgC9H8VZFiNpGYZKnytdx09aDpFoz1XTL
KmqKiREAZKEXqupzkrM1thKuQSRIU0f861lReosqYxwWsMHUCuEUGa5B+ym+ThJj
vlpu0ZmWs0Gzin4yqjBKEq5LdsU7euqaxjHo6f38v0drj31zpAqpPmKIsVsCUVy2
ebisfbmnqIu7UJv1fkCZchBE7qV7PES70IxuhZ+o+JbbI2p5TIueHHagcSHL/ufs
ItZfnp/STFw4GYSX0Xl2e6ZQWBjt0GJnN8xFiQMS5bohh4YV9dwZ9OSU3+a2Wb71
Yu+b1JL3+O+0elQs9JxFiVW/LBgTPiIXFLm4LJDFVRNGrv2k/7xhpHkPPiZFXdH6
mdfa99YooHOIjTDXG51lkme5UAjZoQqHI3RBrP8c31Xunm6KaCijijcZFpWBc712
vYCf3+m8yHu3xJiqVWDdkTSOn7QHFRHpeM/EpcRwdp4abGhcvE4YXuLuLMhyTpXC
GOSGjB1CewxAJZ20wGvYyK4vq2dbEKTaUeUvHsG7g31SleHxJ8H3JIMHSZEi3wIn
1uXtKnp2ney0s9pV9HCN7Np4EjNLa1/2kFwVd3FO3LrsKuMfnN6Yy+8EHBUEPG3+
6zrVeONwq3xxzH/OK1Qnjs1Cb/W98gymCmtCfWWtpWrMK00G+pZjVoEhINEgjTQP
Zi37cwtLIGSy8BMrXHUHPuVwIxlOtALo4TCjAm7mtYueoCkJqJMd6RyTevJ3by6/
K8e4JnjlIUrGImSU6mZ1VA5WMDzusowbL067eAgzTdVZljonBumiRUeaJLR0IfxZ
ctrd6ytQjVsEyv3i6+63bRbc3tjIlxT5j11VLssz+Kxw5uJUWaieIvywNNraGkmO
TFLII1CbFSWm8uwI/XsDUvTkNr8MhgFut6WQslhZHvzXctXZ+AmE4LULkZIQOxML
rMkgWlkAP2O69hpiJIUjsQC88BYUBtyckkZcn6p38mQG1HDxkKSjdJChf+NNM6Y7
uNYSNfR1Xpn5Jp/MLNnWRlSlHGbM66fxONmcC2FaY4xHLt89EzHOyE6h93uXkeQP
9p7SEJuRUO8Kp1TH9XdmvEAbxAjNUgVzYv4Yo9/ucb13jBCfQJ4K+uNx0IeNlBJF
1LogU6Y3x/oxwU9DAcKEqORXRFEuZLCk2a3SWeIhotwq9k304Dzg1YVxd3uZcyu0
Nf5XnFdA88vB/zRc9HvheX85TNcA6oU1rTfSgL1l3gNTSzvkU9I=
-----END RSA PRIVATE KEY-----
`)

// See (https://github.com/kubernetes/kubernetes/issues/126134).
func TestFallbackClient_WebSocketHTTPSProxyNoFallback(t *testing.T) {
	cert, err := tls.X509KeyPair(localhostCert, localhostKey)
	if err != nil {
		t.Errorf("https (valid hostname): proxy_test: %v", err)
	}

	var proxyCalled atomic.Int64
	proxyHandler := utilnettesting.NewHTTPProxyHandler(t, func(req *http.Request) bool {
		proxyCalled.Add(1)
		return true
	})
	defer proxyHandler.Wait()

	proxyServer := httptest.NewUnstartedServer(proxyHandler)
	proxyServer.TLS = &tls.Config{Certificates: []tls.Certificate{cert}}
	proxyServer.StartTLS()
	defer proxyServer.Close() //nolint:errcheck

	proxyLocation, err := url.Parse(proxyServer.URL)
	require.NoError(t, err)

	// Create fake WebSocket server. Copy received STDIN data back onto STDOUT stream.
	websocketServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		conns, err := webSocketServerStreams(req, w, streamOptionsFromRequest(req))
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		defer conns.conn.Close() //nolint:errcheck
		// Loopback the STDIN stream onto the STDOUT stream.
		_, err = io.Copy(conns.stdoutStream, conns.stdinStream)
		if err != nil {
			t.Fatalf("websocket copy error: %v", err)
		}
	}))
	defer websocketServer.Close() //nolint:errcheck

	// Now create the WebSocket client (executor), and point it to the TLS proxy server.
	// The proxy server should open a websocket connection to the fake websocket server.
	websocketServer.URL = websocketServer.URL + "?" + "stdin=true" + "&" + "stdout=true"
	websocketLocation, err := url.Parse(websocketServer.URL)
	require.NoError(t, err)
	clientConfig := &rest.Config{
		Host:            websocketLocation.Host,
		TLSClientConfig: rest.TLSClientConfig{CAData: localhostCert},
		Proxy: func(req *http.Request) (*url.URL, error) {
			return proxyLocation, nil
		},
	}
	websocketExecutor, err := NewWebSocketExecutor(clientConfig, "GET", websocketServer.URL)
	require.NoError(t, err)
	emptyURL, _ := url.Parse("")
	spdyExecutor, err := NewSPDYExecutor(clientConfig, "POST", emptyURL)
	require.NoError(t, err)
	// No fallback to spdyExecutor with websocket.
	sawHTTPSProxyError := false
	exec, err := NewFallbackExecutor(websocketExecutor, spdyExecutor, func(err error) bool {
		if httpstream.IsUpgradeFailure(err) {
			t.Errorf("saw upgrade failure: %v", err)
			return true
		}
		if httpstream.IsHTTPSProxyError(err) {
			sawHTTPSProxyError = true
			t.Logf("saw https proxy error: %v", err)
			return true
		}
		return false
	})
	require.NoError(t, err)

	// Generate random data, and set it up to stream on STDIN. The data will be
	// returned on the STDOUT buffer.
	randomSize := 1024 * 1024
	randomData := make([]byte, randomSize)
	if _, err := rand.Read(randomData); err != nil {
		t.Errorf("unexpected error reading random data: %v", err)
	}
	var stdout bytes.Buffer
	options := &StreamOptions{
		Stdin:  bytes.NewReader(randomData),
		Stdout: &stdout,
	}
	errorChan := make(chan error)
	go func() {
		errorChan <- exec.StreamWithContext(context.Background(), *options)
	}()

	select {
	case <-time.After(wait.ForeverTestTimeout):
		t.Fatalf("expect stream to be closed after connection is closed.")
	case err := <-errorChan:
		if err != nil {
			t.Errorf("unexpected error")
		}
	}

	data, err := io.ReadAll(bytes.NewReader(stdout.Bytes()))
	if err != nil {
		t.Errorf("error reading the stream: %v", err)
		return
	}
	// Check the random data sent on STDIN was the same returned on STDOUT.
	if !bytes.Equal(randomData, data) {
		t.Errorf("unexpected data received: %d sent: %d", len(data), len(randomData))
	}

	// Ensure the https proxy error was *not* observed
	if sawHTTPSProxyError {
		t.Errorf("expected to *not* see https proxy error")
	}
	// Ensure the proxy was called once
	if e, a := int64(1), proxyCalled.Load(); e != a {
		t.Errorf("expected %d proxy call, got %d", e, a)
	}
}
