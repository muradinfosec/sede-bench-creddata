package models

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/util"
)

func TestDataSourceCache(t *testing.T) {
	Convey("When caching a datasource proxy", t, func() {
		clearCache()
		ds := DataSource{
			Id:   1,
			Url:  "http://k8s:8001",
			Type: "Kubernetes",
		}

		t1, err := ds.GetHttpTransport()
		So(err, ShouldBeNil)

		t2, err := ds.GetHttpTransport()
		So(err, ShouldBeNil)

		Convey("Should be using the cached proxy", func() {
			So(t2, ShouldEqual, t1)
		})
		Convey("Should verify TLS by default", func() {
			So(t1.TLSClientConfig.InsecureSkipVerify, ShouldEqual, false)
		})
		Convey("Should have no TLS client certificate configured", func() {
			So(len(t1.TLSClientConfig.Certificates), ShouldEqual, 0)
		})
		Convey("Should have no user-supplied TLS CA onfigured", func() {
			So(t1.TLSClientConfig.RootCAs, ShouldBeNil)
		})
	})

	Convey("When caching a datasource proxy then updating it", t, func() {
		clearCache()
		setting.SecretKey = "password"

		json := simplejson.New()
		json.Set("tlsAuthWithCACert", true)

		tlsCaCert, err := util.Encrypt([]byte(caCert), "password")
		So(err, ShouldBeNil)
		ds := DataSource{
			Id:             1,
			Url:            "http://k8s:8001",
			Type:           "Kubernetes",
			SecureJsonData: map[string][]byte{"tlsCACert": tlsCaCert},
			Updated:        time.Now().Add(-2 * time.Minute),
		}

		t1, err := ds.GetHttpTransport()
		So(err, ShouldBeNil)

		Convey("Should verify TLS by default", func() {
			So(t1.TLSClientConfig.InsecureSkipVerify, ShouldEqual, false)
		})
		Convey("Should have no TLS client certificate configured", func() {
			So(len(t1.TLSClientConfig.Certificates), ShouldEqual, 0)
		})
		Convey("Should have no user-supplied TLS CA configured", func() {
			So(t1.TLSClientConfig.RootCAs, ShouldBeNil)
		})

		ds.JsonData = nil
		ds.SecureJsonData = map[string][]byte{}
		ds.Updated = time.Now()

		t2, err := ds.GetHttpTransport()
		So(err, ShouldBeNil)

		Convey("Should have no user-supplied TLS CA configured after the update", func() {
			So(t2.TLSClientConfig.RootCAs, ShouldBeNil)
		})
	})

	Convey("When caching a datasource proxy with TLS client authentication enabled", t, func() {
		clearCache()
		setting.SecretKey = "password"

		json := simplejson.New()
		json.Set("tlsAuth", true)

		tlsClientCert, err := util.Encrypt([]byte(clientCert), "password")
		So(err, ShouldBeNil)
		tlsClientKey, err := util.Encrypt([]byte(clientKey), "password")
		So(err, ShouldBeNil)

		ds := DataSource{
			Id:       1,
			Url:      "http://k8s:8001",
			Type:     "Kubernetes",
			JsonData: json,
			SecureJsonData: map[string][]byte{
				"tlsClientCert": tlsClientCert,
				"tlsClientKey":  tlsClientKey,
			},
		}

		tr, err := ds.GetHttpTransport()
		So(err, ShouldBeNil)

		Convey("Should verify TLS by default", func() {
			So(tr.TLSClientConfig.InsecureSkipVerify, ShouldEqual, false)
		})
		Convey("Should have a TLS client certificate configured", func() {
			So(len(tr.TLSClientConfig.Certificates), ShouldEqual, 1)
		})
	})

	Convey("When caching a datasource proxy with a user-supplied TLS CA", t, func() {
		clearCache()
		setting.SecretKey = "password"

		json := simplejson.New()
		json.Set("tlsAuthWithCACert", true)

		tlsCaCert, err := util.Encrypt([]byte(caCert), "password")
		So(err, ShouldBeNil)

		ds := DataSource{
			Id:             1,
			Url:            "http://k8s:8001",
			Type:           "Kubernetes",
			JsonData:       json,
			SecureJsonData: map[string][]byte{"tlsCACert": tlsCaCert},
		}

		tr, err := ds.GetHttpTransport()
		So(err, ShouldBeNil)

		Convey("Should verify TLS by default", func() {
			So(tr.TLSClientConfig.InsecureSkipVerify, ShouldEqual, false)
		})
		Convey("Should have a TLS CA configured", func() {
			So(len(tr.TLSClientConfig.RootCAs.Subjects()), ShouldEqual, 1)
		})
	})

	Convey("When caching a datasource proxy when user skips TLS verification", t, func() {
		clearCache()

		json := simplejson.New()
		json.Set("tlsSkipVerify", true)

		ds := DataSource{
			Id:       1,
			Url:      "http://k8s:8001",
			Type:     "Kubernetes",
			JsonData: json,
		}

		tr, err := ds.GetHttpTransport()
		So(err, ShouldBeNil)

		Convey("Should skip TLS verification", func() {
			So(tr.TLSClientConfig.InsecureSkipVerify, ShouldEqual, true)
		})
	})
}

func clearCache() {
	ptc.Lock()
	defer ptc.Unlock()

	ptc.cache = make(map[int64]cachedTransport)
}

const caCert string = `-----BEGIN CERTIFICATE-----
MIIDATCCAemgAwIBAgIJAMQ5hC3CPDTeMA0GCSqGSIb3DQEBCwUAMBcxFTATBgNV
BAMMDGNhLWs4cy1zdGhsbTAeFw0xNjEwMjcwODQyMjdaFw00NDAzMTQwODQyMjda
MBcxFTATBgNVBAMMDGNhLWs4cy1zdGhsbTCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBAMLe2AmJ6IleeUt69vgNchOjjmxIIxz5sp1vFu94m1vUip7CqnOg
QkpUsHeBPrGYv8UGloARCL1xEWS+9FVZeXWQoDmbC0SxXhFwRIESNCET7Q8KMi/4
4YPvnMLGZi3Fjwxa8BdUBCN1cx4WEooMVTWXm7RFMtZgDfuOAn3TNXla732sfT/d
1HNFrh48b0wA+HhmA3nXoBnBEblA665hCeo7lIAdRr0zJxJpnFnWXkyTClsAUTMN
iL905LdBiiIRenojipfKXvMz88XSaWTI7JjZYU3BvhyXndkT6f12cef3I96NY3WJ
0uIK4k04WrbzdYXMU3rN6NqlvbHqnI+E7aMCAwEAAaNQME4wHQYDVR0OBBYEFHHx
2+vSPw9bECHj3O51KNo5VdWOMB8GA1UdIwQYMBaAFHHx2+vSPw9bECHj3O51KNo5
VdWOMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAH2eV5NcV3LBJHs9
I+adbiTPg2vyumrGWwy73T0X8Dtchgt8wU7Q9b9Ucg2fOTmSSyS0iMqEu1Yb2ORB
CknM9mixHC9PwEBbkGCom3VVkqdLwSP6gdILZgyLoH4i8sTUz+S1yGPepi+Vzhs7
adOXtryjcGnwft6HdfKPNklMOHFnjw6uqpho54oj/z55jUpicY/8glDHdrr1bh3k
MHuiWLGewHXPvxfG6UoUx1te65IhifVcJGFZDQwfEmhBflfCmtAJlZEsgTLlBBCh
FHoXIyGOdq1chmRVocdGBCF8fUoGIbuF14r53rpvcbEKtKnnP8+96luKAZLq0a4n
3lb92xM=
-----END CERTIFICATE-----`

const clientCert string = `
-----BEGIN CERTIFICATE-----
MIICsjCCAZoCCQCcd8sOfstQLzANBgkqhkiG9w0BAQsFADAXMRUwEwYDVQQDDAxj
YS1rOHMtc3RobG0wHhcNMTYxMTAyMDkyNTE1WhcNMTcxMTAyMDkyNTE1WjAfMR0w
GwYDVQQDDBRhZG0tZGFuaWVsLWs4cy1zdGhsbTCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBAOMliaWyNEUJKM37vWCl5bGub3lMicyRAqGQyY/qxD9yKKM2
FbucVcmWmg5vvTqQVl5rlQ+c7GI8OD6ptmFl8a26coEki7bFr8bkpSyBSEc5p27b
Z0ORFSqBHWHQbr9PkxPLYW6T3gZYUtRYv3OQgGxLXlvUh85n/mQfuR3N1FgmShHo
GtAFi/ht6leXa0Ms+jNSDLCmXpJm1GIEqgyKX7K3+g3vzo9coYqXq4XTa8Efs2v8
SCwqWfBC3rHfgs/5DLB8WT4Kul8QzxkytzcaBQfRfzhSV6bkgm7oTzt2/1eRRsf4
YnXzLE9YkCC9sAn+Owzqf+TYC1KRluWDfqqBTJUCAwEAATANBgkqhkiG9w0BAQsF
AAOCAQEAdMsZg6edWGC+xngizn0uamrUg1ViaDqUsz0vpzY5NWLA4MsBc4EtxWRP
ueQvjUimZ3U3+AX0YWNLIrH1FCVos2jdij/xkTUmHcwzr8rQy+B17cFi+a8jtpgw
AU6WWoaAIEhhbWQfth/Diz3mivl1ARB+YqiWca2mjRPLTPcKJEURDVddQ423el0Q
4JNxS5icu7T2zYTYHAo/cT9zVdLZl0xuLxYm3asK1IONJ/evxyVZima3il6MPvhe
58Hwz+m+HdqHxi24b/1J/VKYbISG4huOQCdLzeNXgvwFlGPUmHSnnKo1/KbQDAR5
llG/Sw5+FquFuChaA6l5KWy7F3bQyA==
-----END CERTIFICATE-----`

const clientKey string = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA4yWJpbI0RQkozfu9YKXlsa5veUyJzJECoZDJj+rEP3IoozYV
X3IqEtfCSA+5gvPUUgTrZ7cyPhG4tnU0GnBMLNmtGgCMIIqPezVooolelmnmcnvi
i3ghvHZzQSDQi1+zP5zAlsOPdFHH8AE/d1CYWyImO4mRZFZ+ev+1xB6OGnOapQZS
1MCc+B5KR9qzfSB6q5KFqJOvYGoPEccmSAPnQOH1gM/cd9DWjeuHWhmzQo+PM/Ou
lLRD6VMDrz+Mz/ArQraKXrO5lKZIAGI1cMWnv7t/uyrqveDYYiYnb3W/J8xai/oP
apqdr8DtPP9rci26RbG/2EwZXtKt7BC+FqWfUxHmJzUGfiySyqtZM1e715XFW9nf
wI8nwmsaE93/xwNJRYXz6mTTnNxGxJ5dEnn4c3plP4dL+Wr6NbdCr8BwenFpjnLx
zMZQRtFFgh8ZgQDoYnkYTBrYylyQb2+cjqUtHVAMcAO8/3xu8Z/Tyj2M+0bb4v+D
11xtpJBzQmeF3/+SeYhwcV7UBCNvriH7jCm4qzK8FWyCeYBp/m3iCqQZe/7c2/zs
2X0Wa6h6xJCYwV1/f9++7v6CiAKXoN/r/QdzzXQkmPi/go1pSS9TjAb9RskG2deo
ZQ1DpH1IJJxpRpS/5riUMIxO9XD2UG4f+e+6o6flJLqE2vcURVfLV6q/g5BSSoBV
mFPx4nqyqNfwaoYEP4DV12f+2Wg/olVPHIxKd48UCy4D2uEEwy+FK/cjBrJ1O4Wp
S0RZShU7fsdyIO9Bt0LJr08QsZmxFS/Vq5M0eWuJ+gSnyjXZLgq2y0muGwmZlHtE
oiudFIv9PhXP6nnhFJQGN1dE0vTLHTRi2dRULelOLRv+iNbLzjkw7Z7gQZTSfY61
zenjNAn3s6fi12vPZjy7IZSLHnrTon2qjsI6kC3raRnDH21snhFH4VJ44TMoXVyO
faZq26zQp/UYZ2j2TOXOXGadwOpyHxWz7KgJ1xnbI4COIelP1+OSbVpx7y7DsvwA
sBoqw2mBZZqLUslnPywA/Fo6VqKcLeig0jnz56jcnkDayA9kyw0UMKy+Nj8yUH44
DQ9dR+blhGEe8ymfrmPwUSWPCQ90R9EReIEkIaDVgcEDZ/0zZ0F5Vk0KjJDDJ1s1
k3gEu1WAaxq/H7MORmMxEa6iJqQrzhG+K5F+yt+WjMKISzbiUpGyL3kdWGvgfrJ7
h7Yb3N5esOY7yGIAxhI9rzNIZpEUbITARykut3MOxbizl5Zkjr+DInkJTAAJjEjA
DHeO1mPKE+OQBQuzxV/vLwxww61SoqK5HLxWFHykv7YKw0f92+fO3RSOW71cM4bw
oy6NsQ6wFgTnchXrnhkDpg2/clicGibxc305F/eSRdihfiMQnc1RtkhQ6mvNxBQ5
lUTWF+CHbIGBJcpsD2Sm1aHY0RXEQFq9bHWId2dv6slQd+p7j6FY8X03KFiPP8hR
sW2VAlJMKU4IZz0I7NvaFPgTQi9OMIknN90kwOPfQoh9QROoqMkpTp1duw0rsXvL
+EqBLbEOWtamQmoocWyqpAyHTpr1gRrhONcxkNCv/2/GZYlaRP47AJY=
-----END RSA PRIVATE KEY-----`
