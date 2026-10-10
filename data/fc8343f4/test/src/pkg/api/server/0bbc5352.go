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

package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"golang.org/x/net/http2"
)

var (
	// startServerShutdown is a signal from the backend after receiving all (25) requests
	// after which the test shuts down the HTTP server
	startServerShutdown = make(chan struct{})

	// handlerLock used in the backendHTTPHandler to count the number of requests and signal the test that the termination can start
	handlerLock = sync.Mutex{}
)

var backendCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIDTjCCAjagAwIBAgIJANYWBFaLyBC/MA0GCSqGSIb3DQEBCwUAMFAxCzAJBgNV
BAYTAlBMMQ8wDQYDVQQIDAZQb2xhbmQxDzANBgNVBAcMBkdkYW5zazELMAkGA1UE
CgwCU0sxEjAQBgNVBAMMCTEyNy4wLjAuMTAeFw0yMDEyMTExMDI0MzBaFw0zMDEy
MDkxMDI0MzBaMFAxCzAJBgNVBAYTAlBMMQ8wDQYDVQQIDAZQb2xhbmQxDzANBgNV
BAcMBkdkYW5zazELMAkGA1UECgwCU0sxEjAQBgNVBAMMCTEyNy4wLjAuMTCCASIw
DQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMYax2q/m/N237UFMFKZsox4EyKq
De+mbaRGeKqnI7Gi9Ai3b7BPCIa7RFJ2ntpGUd5GyL+HCQHG8/f6DjsbUuhZnmn7
F7ZJeih2DP2acKkODdGbXA52kABCMdDs2DMYhR2UwECY2t+DLpxqJqE2ab8pI9Xd
BZ3pCNodS03yHXzfeJV44lCjxoDOi9ynXLjd3w3+FowomHMEBunTepiqnbgoYtnn
RW9tQyQQK5g6+/j/O1M8o71s/0loBT3vKSqNSrdlMOEGrj4yyL/Cw1NmQf1V1sGf
w1QAW5xk7Br5oh8h1D+oflGWV3Y3zluuZQnA9D+vFpjL0969oFedsgr4UU8CAwEA
AaMrMCkwCQYDVR0TBAIwADALBgNVHQ8EBAMCBaAwDwYDVR0RBAgwBocEfwAAATAN
BgkqhkiG9w0BAQsFAAOCAQEAWbOF7TOfGiC59S50okfcS7M4gwz2kcbqOftWzcA1
lT1qX6TWj7A4bVIOMAFK2tWNd4Omk6bnIAxTJdHB7b1hrBjkpt2krEGH1S8xeRRz
Gs62KQwehM3fMhLvYSEqOQMETZn9AjEigYm6ohCO5obG9Gkfz7uvuv9rbIetbAmm
YE9HdDv6qhCqtynpP2yad3v53idlrDnCIe9e4eKUD5uR/MIp9mEFgnMXR1m43/ya
DnmddSsjtzamVvI/+2Cqjb8qT8dMHZrCBK64UwSaJsUKzSeF6yNvZKQ1yfA/NrfV
P6gNULDOqtPgXFP4j+Z402gjYox1bGHjeDHh1OVSnr9jVw==
-----END CERTIFICATE-----`)

var backendKey = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDGGsdqv5vzdt+1
udwQUjmYFXZevv6prK7OWfDRulnqwCQvL0+TUYlbW1YGEm4pajeFDnw/fpSTKpt6
+p56L3fbLM3k+JA2nqzpGuU1stZRGJ7BX5SlcWGeWXjc1DzDVPibhecOVgYwDJ5K
YWgJqoH/GrUH8Hhh8FoIHZqH4K3813ADeAbNX8CoDdGnA1x01Y9p/sQIctdCWUUz
85qHsJ06PrvT36VAcmqlnjhpSLI9/NPjclo9gP0kdHB32DlJgsH4DIGaiY1+SEi/
qRkcHAh9oojpg1cijrOcfbAQ+LOtgLm/wR0IRBy5I84gXRgHeTN/iIyWj5utfYZc
GoyV+uNXvWiGzfFEkQWTMlCRXHt5jgVvyn7mTWwvXGk6mqBUsdZ5/lUVD3/qf/vG
poTpk+Kn7lfFb0mjd3hb0u79SYXpc3iTpcq7XdjrarRMR/EMHsGzTOPaCeFdZHSE
JHSy33+sMBIVyvKqMliK80eSxHjpdY8cZ1YcXmKkigG9YnwW4U3lVB+RciOJT/hH
aq13V8LoVy1giF4Rwf90npOfAoSFrF47oW8y6/aGyq+HUo40otZiQIbdPKGnj/J3
Ngpxg79RgyyVpDnbp45t+YfidBcLcs/15pO0W31r73QflhFNrVPfeWF3aj5beZGX
Y5b3SnTXVnO1JSVu/WJKAzhs4k4d63J53Sac+YbUEEcUFkfm1Gj1+om57C4WjJXS
g46hNT+2tbxmSuWnAbcHe9tIIExGUeJ04mKgmHLTGFCJd/dCJEi77XulYxbpnueG
3YnVXk5CTwjo175pom9DVP/O0tC0rhKnZddYqp//5fIFHwoQ2+uvJv+1HzBl9ltK
LR4AHOcWeeFb2nd6CZ+TDilpgDWWnljvwASBYyHJHMWKeIjge6RS+Fwu0y45Ru8s
XICnc6am06OChImFi3tRkznH8OQmY5IK7PyoianhUG9VdPACWmvUuUTS084YROSv
PX/wZI2T0jsitB7jZEwGlzce2POZtccMwAXgTHXFUMPDMMfQIe8FpZubM3+5D3+w
UahIQm5/1waRJglOVGTxGCMyCyNR+l+uYMSKU6IxUCbUCdpOV6ejqpWjdLOSZ0+e
GS5v06Ju4Fz/5djeCHsPFz9PyVRBLgHsc4/Lic+D6Nni2DWW4FJrh7kEQMtUYibZ
1JY5CwTdK+eGQTl1Iq/hcRCgW1ceqNVN6lCvc6MhEmmzVwyfg0IgQRDK9XCRyP/1
1mGL5RNIFYnH32ApN5zF/sjGYT3uVl0CGttA45ITI92zrXD3TstbFG/PDalnaLK9
DeWr70r82g3N05HGclJE+ik9F1wlV7uPuooxyydGxQEgwhX9zodrifeptO61yFvA
R+PP84R9Vj/z4MdOG+/ULjkO2kLyehOqy5fCx+ugSjGRCXfVd9uLRTkWgbjC6Gpo
2dcYoxvDca+rl6gFftdEa2PE7WPuTNc/JoKquywr0VnMzF+lfaQR5M97/Du6QAgW
6Mo2KcTBaBaNKbJHJy7x1Xn9SP5yF0mcFrk0N+38rRKg+iHbg+DmzOYKCFAgz7Mf
I+U2qNKpJT88AOW2YUBrMB6R
-----END PRIVATE KEY-----`)

// TestGracefulShutdownForActiveHTTP2Streams checks if graceful shut down of HTTP2 server works.
// It expects that all active connections will be finished (without any errors) before the server exits.
//
// The test sends 25 requests to the target server in parallel. Each request is held by the target server for 60s.
// As soon as the target server receives the last request the test calls backendServer.Config.Shutdown which gracefully shuts down the server without interrupting any active connections.
//
// See more at: https://github.com/golang/go/issues/39776
//
// Note this test will fail on upstream golang 1.15
func TestGracefulShutdownForActiveHTTP2Streams(t *testing.T) {
	// set up the backend server
	backendHandler := &backendHTTPHandler{}
	backendServer := httptest.NewUnstartedServer(backendHandler)
	backendCert, err := tls.X509KeyPair(backendCrt, backendKey)
	if err != nil {
		t.Fatalf("backend: invalid x509/key pair: %v", err)
	}
	backendServer.TLS = &tls.Config{
		Certificates: []tls.Certificate{backendCert},
		NextProtos:   []string{http2.NextProtoTLS},
	}
	backendServer.StartTLS()
	defer backendServer.Close()

	// set up the client
	clientCACertPool := x509.NewCertPool()
	clientCACertPool.AppendCertsFromPEM(backendCrt)
	clientTLSConfig := &tls.Config{
		RootCAs:    clientCACertPool,
		NextProtos: []string{http2.NextProtoTLS},
	}
	client := &http.Client{}
	client.Transport = &http2.Transport{
		TLSClientConfig: clientTLSConfig,
	}

	// client request
	sendRequest := func(wg *sync.WaitGroup) {
		defer func() {
			wg.Done()
		}()

		// act
		resp, err := client.Get(fmt.Sprintf("https://127.0.0.1:%d", backendServer.Listener.Addr().(*net.TCPAddr).Port))
		if err != nil {
			t.Errorf("%v", err)
			return
		}

		// validate
		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%v", err)
		}
		if resp.StatusCode != 200 {
			t.Errorf("unexpected HTTP staus: %v, expected: 200", resp.StatusCode)
		}
		expectedProto := "HTTP/2.0"
		if resp.Proto != expectedProto {
			t.Errorf("unexpected response proto: %v, expected: %v", resp.Proto, expectedProto)
		}
	}

	// this function starts the graceful shutdown
	go func() {
		<-startServerShutdown // signal from the backend after receiving all (25) requests

		backendServer.Config.Shutdown(context.Background())
	}()

	wg := sync.WaitGroup{}
	wg.Add(25)
	for i := 0; i < 25; i++ {
		go sendRequest(&wg)
	}
	wg.Wait()

	// validate backendHandler
	if backendHandler.counter != 25 {
		t.Errorf("the target server haven't received all expected requests, expected 25, it got %d", backendHandler.counter)
	}
}

type backendHTTPHandler struct {
	counter int
}

func (b *backendHTTPHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	handlerLock.Lock()
	b.counter++
	if b.counter == 25 {
		startServerShutdown <- struct{}{}
	}
	handlerLock.Unlock()

	time.Sleep(60 * time.Second)

	w.Write([]byte("hello from the backend"))
	w.WriteHeader(http.StatusOK)
}
