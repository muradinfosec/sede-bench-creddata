package main

import "fmt"

const (
	passwd    = "s0lvg5kq" // $ Alert // NOT OK
	_password = ""         // OK
)

// generated using http://travistidwell.com/jsencrypt/demo
const totallyNotAPrivateKey = // NOT OK
`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC/tzdtXKXcX6F3v3hR6+uYyZpIeXhhLflJkY2eILLQfAnwKlT5
minw7KvSFRz5WFTh0yxOvYkf7kyzJcrHEkKdEbExgYXXHXB8KemAktfsSfkP6hiY
MmwQMVkhVt9kbmVnTPScgRrEg+zmwP8xmdzXHYjA/53myNJ4CHy6KSGIRvkhXaZn
rttjgCUJIu6YzjFPmOXJ21JkeEPBiReEfbmm9J/ff9n1Cq0N3RfgC352iDndmoti
9CLNC+VSsyz5wsQO/ukzz3gvTw95YJXZfX/FPiWDQRO57R5534GkKynlxLfbBxvW
vLFPdhdjynvzFCY9tJsBiS0Nzku/HTYhT3azF5FMHLC/au6komxZrEDvGChzn4+5
6RAIiR1lM7cS6/RWzd26l5oogiBdpcWJvBW//9hy3tF2K1lJgTNaG0U5Tkj+I5pJ
0XaWYTZLPPIOgOMdkM/XwvXi/bPvcMTlIyBCTu5lQCLSMtpZjgI8lUz6xhUHI0TS
RhrAlkhWUwTQPnVxQCJuU5YE1CVidGSIFHmcATqdl5Vn9Y/MYUL4zAo5pCDtAe4C
xteGf2Vx1/EIEl0fn/NvNEWS9v7f1m7ydxROqQOlZfgu6eoCxQQXRb6vFwmUsz+3
fkC/RkEp+W6ZIJPEerzH7frIvlk24Y5UpnRW5nKMRaXzzvsVYwzhTu6GwVerN/9R
VXgBROlu9YMvV7SBSupPuvpWbnTYRLV3S3YFe+85vY4F8SXnSThDzT7Dn+i6eauX
FfLNr8vo5uwVAPsfmNbpgF3970hQIVvbwJxDNgDG78wx
-----END RSA PRIVATE KEY-----`

const publicKey = // OK
`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC/tzdtXKXcX6F3v3hR6+uYyZpI
eXhhLflJkY2eILLQfAnwKlT5xIHW5QZcHQV9sCyZ8qSdPGif7PwgMbButMbByiZh
CSugUFb6vjVqoktmslYF4LKHiDgvmlwuJW0TvynxBLzDCwrRP+gpRT8wuAortWAx
/03POTw7Mzi2cIPNsQIDAQAB
-----END PUBLIC KEY-----`

var secretKey = "" // OK

type info struct {
	username string
	password string
}

func main() {
	password := "f1fi" // NOT OK
	tmp := password
	i := info{
		username: "me",
		password: tmp, // NOT OK
	}
	i.password = "w7os8" // NOT OK
	fmt.Println(password, i)
	testPassword := "o8ri"          // OK
	i.password = "test"             // OK
	i.password = testPassword       // OK
	secretKey = "secret"            // OK
	i.password = "--- redacted ---" // OK
	certsDir := "/certs"            // OK
	fmt.Println(certsDir)
	accountParameter := "ACCOUNT" // OK
	fmt.Println(accountParameter)
}
