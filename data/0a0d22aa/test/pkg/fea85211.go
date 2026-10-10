package xml

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	gitlab = &Credential{
		Tags: map[string]string{
			"scope":       "GLOBAL",
			"id":          "gitlab",
			"description": "Gitlab admin user",
			"username":    "gitlabadmin",
			"password":    "{AQAAABAAAAAgPT7JbBVgyWiivobt0CJEduLyP0lB3uyTj+D5WBvVk6jyG6BQFPYGN4Z3VJN2JLDm}",
		},
	}
	bastion = &Credential{
		Tags: map[string]string{
			"scope":            "GLOBAL",
			"id":               "production-bastion",
			"description":      "Production bastion ssh key",
			"username":         "root",
			"passphrase":       "{OLBAXFRTYHINhXFyi4Bg9nYuLTBElp7DH9BuO8v/mLrAu/GFPTlxpUi=}",
			"privateKeySource": "",
			"privateKey":       "{XWHPQJAPQDQTXgQOIHprbcZGqgey/t3uJiI/hwZLgk/RuRDUWuvSdRhaH40gUyqruU7bJSr0Zc3aOWZBrTqUS5b2yk5/v/fzf+HFd0tOA1yWlOuQFgdlq6YW4LFersNB0wKpU05f5WKeBHGv9Dd8JHW5YR2MbY83tByNeWBaqUPHlA6cMvJkNmZX51wTVmNV31TF58bMI1iN6nGao1JfgMyov3A3e/Z2FLoqDZYTChuBGN1E4JdW6cKMxViFqbdgq2h3/TBtbLDoGOQ0dC6ddf+yq0hROM6N5ZQTk6NmbZOfMdkK/sqPfNdBmgmtGXWMz2DTXQq4kH6gK4NmFEHxrT6/4vYn/qzNQrNKrvbBjK2u9YFg2PZY8UBDbHxG40/V7xxUt3Wzsl3N/dD4ijC/mgI4MDS3Lz42A0IYvhSoWUUKQv+RUSU7vU2mKI0wAyqerwQU1MAgtSJMe5Swuh2m1RBDgG67bXtOGOyATZbhTeswrsmwuLu+GODzKAM9apTSpOCNbtkMWvme31jhrX7WhsCe2ow57xu6pM+17lP/lgrL4cXmncBYhYetkpZt2YAPFO6Dch0fvgVuGlrgx9KwdPdcCuXD/Lha4OtD8H0W/yQ73Lm/QSNzBDFe5FPfHppzFDGQ+BAnaKM6JVkDXLfsjyCfYaIuAR9uqNnJKBuNLXxDGBM8MfZkMyjmpS7kSUhLtvuIioIs2F1FEkosZv3drMB/lM9OVc6DuysAonvyTFiIlz+KGVB8JVvNkXFQ91YdEpvBmAGqqIPJlEr3mx2UM9MtCQkWXUarUry+qfZBf442MNVticUUC+mqcjlnY4rtpRLgDiTf+Pb7fnWCZ6GF0o7+ICMIIY/oA/yhJm+qIl6TUkqRCX01iB8peGEpwhLReTMgGdnozqdg9gaHUDUgiz4/TJvFqi0vi6EdN0nwAgzAkJ+jPWS536dxXiZESfDdtpNMYX2QnmlAtRcHSBmynlIVWx1JVmSeOQD2utfaNWoeYbHHkPGGNNzY0YVoLzuB2+tisoX+cXh5yAne7S3ZDNY3wOt6wWk6L5wXGtcnj/phOBxRZGzsLK9XRXsV34lsaUzjDhvQUjAU8ZoRa0IOO075SL0Vwro8CjlPTo4CerVHWuZlKAaXmzf4QHNuW0elw44EjbV9ndb0IKs8k5V2iNkCA9bIg32e8VkTp37OgZfG4BBRFLlDbAHEXWW9yUwDDhuFMynCaYsD}",
		},
	}
	exampleCredential = &Credential{
		Tags: map[string]string{
			"id":       "exampleId",
			"username": "exampleUsername",
			"password": "djabprxYdjpbbac",
		},
	}
)

func Test_reads_credentials_from_xml_file(t *testing.T) {
	expectedCredentials := []Credential{*gitlab, *bastion, *exampleCredential}
	credentialsXml, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/credentials.xml")

	actualCredentials, _ := ParseCredentialsXml(credentialsXml)

	assert.Equal(t, expectedCredentials, *actualCredentials)
}
