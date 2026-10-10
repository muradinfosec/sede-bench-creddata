package cryptography

import (
	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	oldFormatEncryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "user_1",
				"password": "Y+5nTpiRIS+ndtB7lzy7V2tP+u6DU4VaBh+JOuqQzB9=",
			},
		},
		{
			Tags: map[string]string{
				"username":   "user_1",
				"privateKey": "LINE_1\nLINE_2\nLINE_2",
				"passphrase": "5cPfPhhl4cgu2nsVOUv0DhzI+e3JY7UpRy+WEnoArY6=",
			},
		},
	}

	oldFormatDecryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "user_1",
				"password": "password_1\t\t\t\t\t\t\t\t\t"},
		},
		{
			Tags: map[string]string{
				"passphrase": "passphrase\t\t\t\t\t\t\t\t\t",
				"username":   "user_1",
				"privateKey": "LINE_1\nLINE_2\nLINE_2"},
		},
	}

	newFormatEncryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "xfireadmin",
				"password": "AQAAABAAAAAQtnCexFYLFtmTQCL0x3wnirMnXVA7aZy+lfrfso+SjHI=",
			},
		},
		{
			Tags: map[string]string{
				"username": "gitlabadmin",
				"password": "{AQAAABAAAAAgPT7JbBVgyWiivobt0CJEduLyP0lB3uyTj+D5WBvVk6jyG6BQFPYGN4Z3VJN2JLDm}",
			},
		},
		{
			Tags: map[string]string{
				"username":   "root",
				"passphrase": "{AQAAABAAAAAQmEZaw8Ev9tClXWVQye1TR2KgF3p/wGoYs/TEQCmsxCk=}",
				"privateKey": "{BREGMJDKNDLKDkIAWZxvbfBEgrqi/e5qGcJ/brUCgo/PmLEDKdoEbFxkG08ePialmB9bZQz9An1iREKTsWkCW7z6xb7/g/htf+KBm6uMJ0rQdLnUMvytz3RH6DTbfuSL4kNwB93k9FAtOTFh9Ro0PEM4KP8FdU42oUoKhMPhkYDUrK5eClDcXxFK52hGVgYC81YI70jPG4iN6wPwk5YluOkdp7M3j/B6DObnBOQJPllDHL3P3QvD5bVSvVdKelkjq2y6/SCbtBWjXKJ9yC2mhd+de8nFCE1G1VAGc8BlmACuMbpW/vaNpFaNsayuLHAIt7GFOPy2jI1yJ7JgWBCppK1/4sGn/sdYHsVYyiuTuN6o9HIm0FXP4XKBeUeA47/M8vtUd0Bixb3G/wT0ulQ/kpH4UKN6Gh80D1IWheMxNBMMBe+ZITR1mS9tTW1cPmfpoyVR2YUidDDFz6Wsyf1v7WRQnK22qMvHSOmAIGjzZojvdwecpKl+XUYrNJY1iqZWyOHBargSPcmn80rbpG6DubAt9cj02na0rS+45eR/bkiV0aCcsqUFlFawujOm7PORJY6Zcy6wxbVeUdrrh5RpaNhhRjFI/Hut4WzF6U8E/pD73Nl/HKSoRCPs0QNeDhtwIZDN+IZvjVR2QBxTQUruneMgCwVnXH7bkOcJLPeBXUtAOVK1EmNpCwmbyQ9mGKiYonvNswTd6F4AKvfeTo1tuUL/gR5LMo1HrexYszckJRcAct+QWWY4UFfSfPYI68RjMboYgOSjsJDCnJf3mn8ER6VmCCfVZDaaVqq+lhCQp674HUKjgqYTJ+wedobiI0ibsJRuUvFs+Zb8ueXZI6AM4b6+DRVNTB/mJ/bqLi+xFe3AVwyGLW49kC4prCWwvlBPiHYpNvotkjsk9bbTXGErvv9/XMtXzk5ib4YrH2nvXgkYmF+jQDJ993poNcBLAyIrhjFDSM0QezmVaLaFLSrzllTEQo4MWgZsOLZ9ehopMWdfLrJLvOUSCGaP8DDhCxeG6+yyrrO+tJq1rTto9N7EEGD5sAq2bPe2U4rQDzazz/fpQYtCISjzLR5GEBgF65nprLgbKszJJnFR4OfVo8ZWZ961VG0Mnnu3HthBEn6XcjZEXbKcZQlMffj2KNKqO6ong19ZwsQ5lyv2JEn6j1A4jPgYJ2kNv28x7OyDs99HnFbN0XRILCvScFXKGRW3pJmGOvoENcqClZdR}",
			},
		},
	}
	newFormatDecryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "xfireadmin",
				"password": "uqviinvdut",
			},
		},
		{
			Tags: map[string]string{
				"username": "gitlabadmin",
				"password": "Ghqylh3NFbs1jC_26Vo",
			},
		},
		{
			Tags: map[string]string{
				"username":   "root",
				"passphrase": "SLQuV9jJsDE",
				"privateKey": `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCwgxo7cl2RajAWFseL0JAIBJbZ6dFWBGcq7+TMkP8viDwfLj4u
ToxJgO+m/Jt2GUlCybsEGwcGplwfD2b3Y+CuQarWHwjJKP68mIgADUcns4wyTtg4
g4qYas4nAsJuYihmMaQRszkpUVfH97y5PL8oTnfve1uqB9di2cS4xGb5EoIIyQbN
hQuVBQe7sKlKfsa42h6ArGlzYMIh/ovDfv+1xP0+b2JPHkp4PnIDUlCPwwZxkJxM
mgl2dqyZ4oI4BQZ/jPoJslyx2Hr2rqKHvW4Zlb6NweJuRv6nn9vdbZB7mDVpCH8l
m9EagFxFVQJbWZ28KfVU+aXvun5L/hJnjci6J9VQD9VaFwSZuymZBdeZeC3lZXJb
KAKYUhIe6m+O482O1/W9Uh8GyB8xoiRkYUdaVFMsQYOwza8XcwQ3ZK/ZBNhjO6Qw
Z9fl7E4nUejyohQgDyXad+J8N+KjgXFh0OK04ZKbjyT1c+1YB0+OSvp1Wk7X8nTB
F1QrDFBphJjnVt2bRLqO9MrwT3iXldUUvdNzw/mnE+6on7qjCIvFtCwWV6tpd1Zc
XAaJPuAYHlEzUEn6L38zs9zl2tiufbOQfIL1U0RFy72VI1Kb8eK+kKWoWpgpKBtZ
9MGnBavurCH1Q+bHX+VipYmoVJBGlamjL4eqgbwr5v4trkZ7Z180fki0yURP/a/v
7yGyBr9G12yr8DiJu4fqaScO/HQNiwjhViCpsTKk4WQnMD9GJ99xh5RGBnFohbrh
WAC6mPzZgjlzh5HsWFwMGKCecV7Rcl3emZPh1evG2x3=
-----END RSA PRIVATE KEY-----`,
			},
		},
	}
)

func Test_decrypts_old_format_credentials(t *testing.T) {
	secret, _ := ioutil.ReadFile("../../test/resources/jenkins_1.625.1/decrypted/hudson.util.Secret")

	credentials, _ := DecryptCredentials(&oldFormatEncryptedCredentials, secret)

	assert.Equal(t, credentials, oldFormatDecryptedCredentials)
}

func Test_decrypts_new_format_credentials(t *testing.T) {
	secret, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/decrypted/hudson.util.Secret")

	credentials, _ := DecryptCredentials(&newFormatEncryptedCredentials, secret)

	assert.Equal(t, credentials, newFormatDecryptedCredentials)
}
