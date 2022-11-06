package v4_test

import (
	"testing"

	"github.com/robovpn/v2fly_core/common/net"
	"github.com/robovpn/v2fly_core/common/protocol"
	"github.com/robovpn/v2fly_core/common/serial"
	"github.com/robovpn/v2fly_core/infra/conf/cfgcommon"
	"github.com/robovpn/v2fly_core/infra/conf/cfgcommon/testassist"
	v4 "github.com/robovpn/v2fly_core/infra/conf/v4"
	"github.com/robovpn/v2fly_core/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.ShadowsocksServerConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"method": "aes-256-GCM",
				"password": "v2fly-password"
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				User: &protocol.User{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_256_GCM,
						Password:   "v2fly-password",
					}),
				},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
