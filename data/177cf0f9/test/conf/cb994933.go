package conf_test

import (
	"testing"

	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	. "github.com/xtls/xray-core/infra/conf"
	"github.com/xtls/xray-core/proxy/mtproto"
)

func TestMTProtoServerConfig(t *testing.T) {
	creator := func() Buildable {
		return new(MTProtoServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"users": [{
					"email": "love@example.com",
					"level": 1,
					"secret": "a3cfcbe2e805c3394273af87a3d39e9e"
				}]
			}`,
			Parser: loadJSON(creator),
			Output: &mtproto.ServerConfig{
				User: []*protocol.User{
					{
						Email: "love@example.com",
						Level: 1,
						Account: serial.ToTypedMessage(&mtproto.Account{
							Secret: []byte{207, 189, 101, 150, 191, 144, 185, 23, 136, 144, 225, 88, 153, 172, 14, 213},
						}),
					},
				},
			},
		},
	})
}
