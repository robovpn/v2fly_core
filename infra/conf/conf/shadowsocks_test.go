package conf_test

import (
	"testing"

	"github.com/robovpn/v2fly_core/common/net"
	"github.com/robovpn/v2fly_core/common/protocol"
	"github.com/robovpn/v2fly_core/common/serial"
	. "github.com/robovpn/v2fly_core/infra/conf"
	"github.com/robovpn/v2fly_core/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() Buildable {
		return new(ShadowsocksServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"method": "aes-128-cfb",
				"password": "v2ray-password"
			}`,
			Parser: loadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				User: &protocol.User{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_128_CFB,
						Password:   "v2ray-password",
					}),
				},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
