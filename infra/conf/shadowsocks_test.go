package conf_test

import (
	"testing"

	"v2fly_core/common/net"
	"v2fly_core/common/protocol"
	"v2fly_core/common/serial"
	. "v2fly_core/infra/conf"
	"v2fly_core/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() Buildable {
		return new(ShadowsocksServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"method": "aes-128-cfb",
				"password": "v2fly-password"
			}`,
			Parser: loadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				User: &protocol.User{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_128_CFB,
						Password:   "v2fly-password",
					}),
				},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
