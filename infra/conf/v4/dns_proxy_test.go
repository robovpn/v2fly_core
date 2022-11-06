package v4_test

import (
	"testing"

	"github.com/robovpn/github.com/robovpn/v2fly_core/common/net"
	"github.com/robovpn/github.com/robovpn/v2fly_core/infra/conf/cfgcommon"
	"github.com/robovpn/github.com/robovpn/v2fly_core/infra/conf/cfgcommon/testassist"
	v4 "github.com/robovpn/github.com/robovpn/v2fly_core/infra/conf/v4"
	"github.com/robovpn/github.com/robovpn/v2fly_core/proxy/dns"
)

func TestDnsProxyConfig(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.DNSOutboundConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"address": "8.8.8.8",
				"port": 53,
				"network": "tcp"
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &dns.Config{
				Server: &net.Endpoint{
					Network: net.Network_TCP,
					Address: net.NewIPOrDomain(net.IPAddress([]byte{8, 8, 8, 8})),
					Port:    53,
				},
			},
		},
	})
}
