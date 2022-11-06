package v4_test

import (
	"testing"

	"../v2fly_core/common/net"
	"../v2fly_core/common/protocol"
	"../v2fly_core/common/serial"
	"../v2fly_core/infra/conf/cfgcommon"
	"../v2fly_core/infra/conf/cfgcommon/testassist"
	v4 "../v2fly_core/infra/conf/v4"
	"../v2fly_core/proxy/socks"
)

func TestSocksInboundConfig(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.SocksServerConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"auth": "password",
				"accounts": [
					{
						"user": "my-username",
						"pass": "my-password"
					}
				],
				"udp": false,
				"ip": "127.0.0.1",
				"timeout": 5,
				"userLevel": 1
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &socks.ServerConfig{
				AuthType: socks.AuthType_PASSWORD,
				Accounts: map[string]string{
					"my-username": "my-password",
				},
				UdpEnabled: false,
				Address: &net.IPOrDomain{
					Address: &net.IPOrDomain_Ip{
						Ip: []byte{127, 0, 0, 1},
					},
				},
				Timeout:   5,
				UserLevel: 1,
			},
		},
	})
}

func TestSocksOutboundConfig(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.SocksClientConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"servers": [{
					"address": "127.0.0.1",
					"port": 1234,
					"users": [
						{"user": "test user", "pass": "test pass", "email": "test@email.com"}
					]
				}]
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &socks.ClientConfig{
				Server: []*protocol.ServerEndpoint{
					{
						Address: &net.IPOrDomain{
							Address: &net.IPOrDomain_Ip{
								Ip: []byte{127, 0, 0, 1},
							},
						},
						Port: 1234,
						User: []*protocol.User{
							{
								Email: "test@email.com",
								Account: serial.ToTypedMessage(&socks.Account{
									Username: "test user",
									Password: "test pass",
								}),
							},
						},
					},
				},
			},
		},
	})
}
