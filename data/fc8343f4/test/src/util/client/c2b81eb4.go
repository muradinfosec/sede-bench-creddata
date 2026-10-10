/*
Copyright 2018 The Kubernetes Authors.

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

package keyutil

import (
	"os"
	"testing"
)

const (
	// rsaPrivateKey is a RSA Private Key in PKCS#1 format
	// openssl genrsa -out rsa2048.pem 2048
	rsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA92mVjhBKOFsdxFzb/Pjq+7b5TJlODAdY5hK+WxLZTIrfhDPq
QbMswVuqEovANXJTOLtk4O+iIWeplq0iyH5609HUKJTStLfAvNdstE67fVQ0jizX
ShdcLClYzMJRA0CZMlURyF7oFkFUTXe2Jhua6qrG5QIwfswlnQf1ygAPssuC7tG9
4stOTQHrZ8XpzXAL1E5BTHt0EsoVGnJ5Cv8DV91ZnWW5hT2QkEqv2l/s7G6W98p6
4lNYzQvg3/sqt0SLq3tFrqlpcPZw7j/WrUwPjTbjJg+n5aCy6x25Xs/cMrMaeGFE
mPNhRyNGrh/VUEYYX10q6qegrL/iY+Uw153aCaoZklLQnNtFkTsJ6V8tktfqT8/t
UjtWCaHJHxD9cAN2r7XacbC0emFYfB0rTpcnPiCCBh/nu1XPzJw/jtXUez/uBm7h
7N/RU0R3iyXdKRKSW/pM3NgxfAng8OhU4AsG/qkcxQ7YWPctfOX5biTemP3puhoQ
ZxuRPfH8GVXOqaUBXElturDRweYW/PmAknO/T1kB+ua8zNlm2JMqtqunE+yD565i
MbULobPZ2Bso8NnAQ5WDLDCz19fbwc6yWseDjPfb47zoxbsbpwtM9mHrKmO0mMJz
RPaGjgvZtUIGxT3THfhHtfLkOR27TwIjI94hKTOt7FaxDd0vxzRsSHEcBGZY2rQy
wxY6aHBBNRGeiz/8MVx63lTR7Q0SOU8Lzx5laHi2GKBbxxuBcl0R8RKQb6k37JMz
8C2RHxRQVxO8huZetNmuKO7T2jC+DleiFXFAx9ZNn1+hvvACD+GodSMh3nlNocnP
rcmQZt2TxnNZcDDjNTeOFInhvIfbcSWiHdaxex4Q7aHRBIrOedt/c2oAGKEWeAzY
0nNYy6kRRhUHPXIQfq9nHUYdpYh827hUSu9ULeYnaCGBmOjWRBZ0JqxC5Cw9QG49
cIvKN5ZLFyAw6DRByRzqeo+tQ+bxqCRwf2TsVpQghVh7RV1p2mTTIkUwNv0mbIq6
px6or2Cz8BBFp89YItkn+7kz6E3p4SU+eIMzYdCIkGfsKMuy330paqlWiF1+eM3T
zWTwvatdmveis6Lk2K8is45OwyJLzGYKhbuHGnwoNWpkc8XUaalwkU5RQ6q99Mgx
4Nahf4KOWi14Un5vrIroGQKwPTSdxa4cFKkK2JJBV2G50zEzzyrHOzBoIyxrQLEt
sfOKAwGLokwtOnBqauRDqnGMxgMlo6KLMioNIgci9NRo7YvEVyY449J0soeYv7ya
wspaFom0JgL5XW48qIsWXrjC+CflTIqDWnA2Gne1aoj+uBwfXoacMNDKwLPzz18l
U8OXpcKox3cN2PseijXgBgucaoLitrfbbeD88r4HP5CK/RwhRJKzt2od9v/nClwO
cpGzZsRiyQj1rr/tnYKSomUOx6tBjYvfBiqNVmZpmauzGiGL8smlRufUjaVkv1gG
K4UMZqq1QgJTLfaDPIVcj7nJHadXSk9EvJfBw6KePjyaf4g0a9WhtIULkp5Y+nhW
K7EurY7zoaWV+Xzw6AleFO/7R7lfzrgemyZqLA9q0XQQTFs8INKywG==
-----END RSA PRIVATE KEY-----`

	// rsaPublicKey is a RSA Public Key in PEM encoded format
	// openssl rsa -in rsa2048.pem -pubout -out rsa2048pub.pem
	rsaPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA92mVjhBKOFsdxFzb/Pjq
+7b5TJlODAdY5hK+WxLZTIrfhDPqFWrGKdjSNiHbXrdEtwJh9V+RqPZVSN3aWy12
24RgkyNdMJsXhJKuCC24ZKY8SXtWxuTYmMRaMnCsv6QBGRTIbZ2EFbAObVM7lDyv
1VqY3amZIWFQMlZ9CNpxDSPa5yi43gopbXkne0oGNmey9X0qtpk7NMZIgAL6Zz4r
Z30bcfC2ag6RLOFI2E/c4n8c38R89MfXfLkj8/Cxo4JfI9NvRCpPOpFO8d/ZtWVU
uIrBQN+Y7tkN2T60Qq/TkKXUrhDefwlTlktZVJ/GztLYU41b2GcWsh/XO+PH831r
mwIDAQAB
-----END PUBLIC KEY-----`

	// certificate is an x509 certificate in PEM encoded format
	// openssl req -new -key rsa2048.pem -sha256 -nodes -x509 -days 1826 -out x509certificate.pem -subj "/C=US/CN=not-valid"
	certificate = `-----BEGIN CERTIFICATE-----
MIIDFTCCAf2gAwIBAgIJAN8B8NOwtiUCMA0GCSqGSIb3DQEBCwUAMCExCzAJBgNV
BAYTAlVTMRIwEAYDVQQDDAlub3QtdmFsaWQwHhcNMTcwMzIyMDI1NjM2WhcNMjIw
MzIyMDI1NjM2WjAhMQswCQYDVQQGEwJVUzESMBAGA1UEAwwJbm90LXZhbGlkMIIB
IjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA92mVjhBKOFsdxFzb/Pjq+7b5
TJlODAdY5hK+WxLZTIrfhDPqFWrGKdjSNiHbXrdEtwJh9V+RqPZVSN3aWy1224Rg
kyNdMJsXhJKuCC24ZKY8SXtWxuTYmMRaMnCsv6QBGRTIbZ2EFbAObVM7lDyv1VqY
3amZIWFQMlZ9CNpxDSPa5yi43gopbXkne0oGNmey9X0qtpk7NMZIgAL6Zz4rZ30b
cfC2ag6RLOFI2E/c4n8c38R89MfXfLkj8/Cxo4JfI9NvRCpPOpFO8d/ZtWVUuIrB
QN+Y7tkN2T60Qq/TkKXUrhDefwlTlktZVJ/GztLYU41b2GcWsh/XO+PH831rmwID
AQABo1AwTjAdBgNVHQ4EFgQU1I5GfinLF7ta+dBJ6UWcrYaexLswHwYDVR0jBBgw
FoAU1I5GfinLF7ta+dBJ6UWcrYaexLswDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B
AQsFAAOCAQEAUl0wUD4y41juHFOVMYiziPYr1ShSpQXdwp8FfaHrzI5hsr8UMe8D
dzb9QzZ4bx3yZhiG3ahrSBh956thMTHrKTEwAfJIEXI4cuSVWQAaOJ4Em5SDFxQe
d0E6Ui2nGh1SFGF7oyuEXyzqgRMWFNDFw9HLUNgXaO18Zfouw8+K0BgbfEWEcSi1
JLQbyhCjz088gltrliQGPWDFAg9cHBKtJhuTzZkvuqK1CLEmBhtzP1zFiGBfOJc8
v+aKjAwrPUNX11cXOCPxBv2qXMetxaovBem6AI2hvypCInXaVQfP+yOLubzlTDjS
Y708SlY38hmS1uTwDpyLOn8AKkZ8jtx75g==
-----END CERTIFICATE-----`

	// ecdsaPrivateKeyWithParams is a ECDSA Private Key with included EC Parameters block
	// openssl ecparam -name prime256v1 -genkey -out ecdsa256params.pem
	ecdsaPrivateKeyWithParams = `-----BEGIN EC PARAMETERS-----
BggqhkjOPQMBBw==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIAwSOWQqlMTZNqNF7tgua812Jxib1DVOgb2pHHyIEyNNoAoGCCqGSM49
lrGIzBWqXLbKvwiVMG7Z7AbjizBDy+P+yKMr5COipGWMKOGY7J10n9Ii1B4SV4CH
HJR5VAAH9KGFGTKaetdmIbp5shPFBrqkUr==
-----END EC PRIVATE KEY-----`

	// ecdsaPrivateKey is a ECDSA Private Key in ASN.1 format
	// openssl ecparam -name prime256v1 -genkey -noout -out ecdsa256.pem
	ecdsaPrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIP6Qw6dHDiLsSnLXUhQVTPE0fTQQrj3XSbiQAZPXnk5+oAoGCCqGSM49
VhvorBeMJEnScQvq5M0N3/mKSsR/kcgFg+N9PIW0f4SCzM7C9FtH+/CtX0cZJl+5
Ody4CCu7dYfuYzyYoXvJbgjOueo20XfWVI==
-----END EC PRIVATE KEY-----`

	// ecdsaPublicKey is a ECDSA Public Key in PEM encoded format
	// openssl ec -in ecdsa256.pem -pubout -out ecdsa256pub.pem
	ecdsaPublicKey = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEZZzi1u5f2/AEGFI/HYUhU+u6cTK1
q2bbtE7r1JMK+/sQA5sNAp+7Vdc3psr1OaNzyTyuhTECyRdFKXm63cMnGg==
-----END PUBLIC KEY-----`
)

func TestReadPrivateKey(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating tmpfile: %v", err)
	}
	defer os.Remove(f.Name())

	if _, err := PrivateKeyFromFile(f.Name()); err == nil {
		t.Fatalf("Expected error reading key from empty file, got none")
	}

	if err := os.WriteFile(f.Name(), []byte(rsaPrivateKey), os.FileMode(0600)); err != nil {
		t.Fatalf("error writing private key to tmpfile: %v", err)
	}
	if _, err := PrivateKeyFromFile(f.Name()); err != nil {
		t.Fatalf("error reading private RSA key: %v", err)
	}

	if err := os.WriteFile(f.Name(), []byte(ecdsaPrivateKey), os.FileMode(0600)); err != nil {
		t.Fatalf("error writing private key to tmpfile: %v", err)
	}
	if _, err := PrivateKeyFromFile(f.Name()); err != nil {
		t.Fatalf("error reading private ECDSA key: %v", err)
	}

	if err := os.WriteFile(f.Name(), []byte(ecdsaPrivateKeyWithParams), os.FileMode(0600)); err != nil {
		t.Fatalf("error writing private key to tmpfile: %v", err)
	}
	if _, err := PrivateKeyFromFile(f.Name()); err != nil {
		t.Fatalf("error reading private ECDSA key with params: %v", err)
	}
}

func TestReadPublicKeys(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating tmpfile: %v", err)
	}
	defer os.Remove(f.Name())

	if _, err := PublicKeysFromFile(f.Name()); err == nil {
		t.Fatalf("Expected error reading keys from empty file, got none")
	}

	if err := os.WriteFile(f.Name(), []byte(rsaPublicKey), os.FileMode(0600)); err != nil {
		t.Fatalf("error writing public key to tmpfile: %v", err)
	}
	if keys, err := PublicKeysFromFile(f.Name()); err != nil {
		t.Fatalf("error reading RSA public key: %v", err)
	} else if len(keys) != 1 {
		t.Fatalf("expected 1 key, got %d", len(keys))
	}

	if err := os.WriteFile(f.Name(), []byte(ecdsaPublicKey), os.FileMode(0600)); err != nil {
		t.Fatalf("error writing public key to tmpfile: %v", err)
	}
	if keys, err := PublicKeysFromFile(f.Name()); err != nil {
		t.Fatalf("error reading ECDSA public key: %v", err)
	} else if len(keys) != 1 {
		t.Fatalf("expected 1 key, got %d", len(keys))
	}

	if err := os.WriteFile(f.Name(), []byte(rsaPublicKey+"\n"+ecdsaPublicKey), os.FileMode(0600)); err != nil {
		t.Fatalf("error writing public key to tmpfile: %v", err)
	}
	if keys, err := PublicKeysFromFile(f.Name()); err != nil {
		t.Fatalf("error reading combined RSA/ECDSA public key file: %v", err)
	} else if len(keys) != 2 {
		t.Fatalf("expected 2 keys, got %d", len(keys))
	}

	if err := os.WriteFile(f.Name(), []byte(certificate), os.FileMode(0600)); err != nil {
		t.Fatalf("error writing certificate to tmpfile: %v", err)
	}
	if keys, err := PublicKeysFromFile(f.Name()); err != nil {
		t.Fatalf("error reading public key from certificate file: %v", err)
	} else if len(keys) != 1 {
		t.Fatalf("expected 1 keys, got %d", len(keys))
	}

}
