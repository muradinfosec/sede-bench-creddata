package httptest

import (
	"crypto/tls"
	"net"
)

// copied from net/http/httptest/internal

// LocalhostCert is a PEM-encoded TLS cert with SAN IPs
// "127.0.0.1" and "[::1]", expiring at Jan 29 16:00:00 2084 GMT.
// generated from src/crypto/tls:
// go run generate_cert.go  --rsa-bits 1024 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
// note: these are not the net/http/httptest/internal contents but doesn't matter.
var LocalhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDAzCCAeugAwIBAgIJAP0pWSuIYyQCMA0GCSqGSIb3DQEBBQUAMBgxFjAUBgNV
BAMMDWxvY2FsaG9zdDozMzEwHhcNMTYxMjI1MDk1OTI3WhcNMjYxMjIzMDk1OTI3
WjAYMRYwFAYDVQQDDA1sb2NhbGhvc3Q6MzMxMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEA5vETjLa+8W856rWXO1xMF/CLss9vn5xZhPXKhgz+D7ogSAXm
mWP53eeBUGC2r26J++CYfVqwOmfJEu9kkGUVi8cGMY9dHeIFPfxD31MYX175jJQe
tu0WeUII7ciNsSUDyBMqsl7yi1IgN7iLONM++1+QfbbmNiEbghRV6icEH6M+bWlz
3YSAMEdpK3mg2gsugfLKMwJkaBKEehUNMySRlIhyLITqt1exYGaggRd1zjqUpqpD
sL2sRVHJ3qHGkSh8nVy8MvG8BXiFdYQJP3mCQDZzruCyMWj5/19KAyu7Cto3Bcvu
PgujnwRtU+itt8WhZUVtU1n7Ivf6lMJTBcc4OQIDAQABo1AwTjAdBgNVHQ4EFgQU
MXrBvbILQmiwjUj19aecF2N+6IkwHwYDVR0jBBgwFoAUMXrBvbILQmiwjUj19aec
F2N+6IkwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQUFAAOCAQEA4zbFml1t9KXJ
OijAV8gALePR8v04DQwJP+jsRxXw5zzhc8Wqzdd2hjUd07mfRWAvmyywrmhCV6zq
OHznR+aqIqHtm0vV8OpKxLoIQXavfBd6axEXt3859RDM4xJNwIlxs3+LWGPgINud
wjJqjyzSlhJpQpx4YZ5Da+VMiqAp8N1UeaZ5lBvmSDvoGh6HLODSqtPlWMrldRW9
AfsXVxenq81MIMeKW2fSOoPnWZ4Vjf1+dSlbJE/DD4zzcfbyfgY6Ep/RrUltJ3ag
FQbuNTQlgKabe21dSL9zJ2PengVKXl4Trl+4t/Kina9N9Jw535IRCSwinD6a/2Ca
m7DnVXFiVA==
-----END CERTIFICATE-----
`)

// LocalhostKey is the private key for localhostCert.
var LocalhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA5vETjLa+8W856rWXO1xMF/CLss9vn5xZhPXKhgz+D7ogSAXm
gOU26DWgZzU4J25n++ArMneIYgGIiy5ZCMxpF2jdUE6cDGispQlw69NRn841SKuo
DO8hETue4LBoZovnCQYOBj0rm3PTx1hOHCY++8+DLnWABnEOOlEQ0yGUu3r+EIPS
4ZixXtWsO9Hl0GGIHaHYkDghSkKXuqOEiLhLXLyubEbBe0krAvLbGnu9iMlhDeyO
Cy2SMrGM4xylxpt4HMV7eMo2eKNXhcMvc6EalAnYGmmodxs8/29OOxp5zVW2oyLv
rGUsfIiqH+yIw6KmQGSAa4p0WAn4pmTECTD9KORUQzgvxLHHxkuwMC5CMWiSmOe3
+wSDJmK2YmTf8pUVn1wAadxBRB92wn3h+hO68d9bikshGA41s6uzK/2O7jx6HhQ3
CiSZZLoSUYJ/rlpETQB89PO+cQCICm0nPJPtrgW6skfgfObQfVksb0fpbtYGdtQL
KpE3tmPgx7GAcQgXjBGgdJnwMOM2bPqMZNGYDfsK3+8d5LXSWiL1tH43j6eI8906
bOrwVIx2sSDub/F/IZQux/rnzE2/0TlBsWO+UXvNUTo53UH3Rctq0AQuQzj6O8KR
VYrTaWaftd1xK02+9iU25xXtXNfHs8Qj8C9qJp+K6DkO+55L7lOOkjB9xCqeMtu7
SLy05t/UTqBUwRyG0EF3eu/itWuAXqZYeK6Mfm97wXoXuosJyq9lVp/vuHMl7Ru+
aHJT8B1UrPjjyi6ulR2eR0NUNThUW60gVfpyxM0KUV/8gtwNK0iFwxr4oQO76PCF
V+fr8DaKazbMZqgOxfaYJ5hnWamkTlquHomtQXpOvfH5vWaYYs+S52TTppmdziWU
MZs4RE2eSGbcwEetLI9Uy6NRYMf4OmKCwPLbmqvK3b5ADNRqoeUWyi6tyOnEy41D
GbI4Tr6dok4wEQ0sGop8ckCmmFgAFjXVTGrt4ohGS5IdtuB8P/ktOnBfF6LbkqPc
r4sh8CvrGqJVW86qTmwjjHz5OED+vIaswl9GY7PtWtSoaETlOisw6jDr8I2jB54s
VtVogiiMHT565B+drX+5/rks+Ms3dyd0ZoFFdknmeS/ZhGUM8gMN50s3ky773Zcm
nfSJmQsVlPHb8LyLW/0HDbZd4ogU5jlKTkSacfIDYb4pGILTkPkF2TgD3vhhrbQM
okTbwlxymSoVN+YcWQrBl1lniLSXSuQ9beTdPW0IXuNBb9ODSQ0sZasl4iUvk04N
SmWMAv5x6hJ10NyUf8Zsj2Rd0mcqYbLBSnrbnYWxB4wR2N4QgDT8Gvy+LXtnJ2dx
XGTxGmXrP3uVsNtuJzUA8d8bgsSmvQNfHz9VYEV8x+9NPwjv/GPyZ1LVsij9xGf+
7dUyuvxLrLEZbAXfkzz5VNYJnxAdPd/SODBYUPKh0R9j57x0TcsRVDy0IDWRlbrE
1hqgEUqg1WhBwBJMe5lIog3p59OtmeqDBEzSKGKGZkh63Yu9bZuL3saO20gRe8wU
h0N8OePPROe7YF9SpTe/iP9xZVSoK6CeAp5qzvMy+Jfxr1DIio5mKY==
-----END RSA PRIVATE KEY-----
`)

// NewLocalListener returns a new ipv4 "127.0.0.1:0"
// or tcp6 "[::1]:0" tcp listener.
func NewLocalListener() net.Listener {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			panic(err)
		}
	}
	return l
}

// NewLocalTLSListener returns a new tls listener
// based on the "tcpListener", if "tcpListener" is nil
// it make use of the `NewLocalListener`.
// Cert and Key are `LocalhostCert` and `LocalhostKey` respectfully.
func NewLocalTLSListener(tcpListener net.Listener) net.Listener {
	if tcpListener == nil {
		tcpListener = NewLocalListener()
	}

	cert, err := tls.X509KeyPair(LocalhostCert, LocalhostKey)
	if err != nil {
		panic(err)
	}

	cfg := new(tls.Config)
	cfg.NextProtos = []string{"http/1.1"}
	cfg.Certificates = []tls.Certificate{cert}
	cfg.InsecureSkipVerify = true
	return tls.NewListener(tcpListener, cfg)
}
