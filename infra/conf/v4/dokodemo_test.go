package v4_test

import (
	"testing"

	"../v2fly_core/common/net"
	"../v2fly_core/infra/conf/cfgcommon"
	"../v2fly_core/infra/conf/cfgcommon/testassist"
	v4 "../v2fly_core/infra/conf/v4"
	"../v2fly_core/proxy/dokodemo"
)

func TestDokodemoConfig(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.DokodemoConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"address": "8.8.8.8",
				"port": 53,
				"network": "tcp",
				"timeout": 10,
				"followRedirect": true,
				"userLevel": 1
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &dokodemo.Config{
				Address: &net.IPOrDomain{
					Address: &net.IPOrDomain_Ip{
						Ip: []byte{8, 8, 8, 8},
					},
				},
				Port:           53,
				Networks:       []net.Network{net.Network_TCP},
				Timeout:        10,
				FollowRedirect: true,
				UserLevel:      1,
			},
		},
	})
}
