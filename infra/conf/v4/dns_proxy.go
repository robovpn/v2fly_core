package v4

import (
	"google.golang.org/protobuf/proto"

	"github.com/robovpn/v2fly_core/common/net"
	"github.com/robovpn/v2fly_core/infra/conf/cfgcommon"
	"github.com/robovpn/v2fly_core/proxy/dns"
)

type DNSOutboundConfig struct {
	Network   cfgcommon.Network  `json:"network"`
	Address   *cfgcommon.Address `json:"address"`
	Port      uint16             `json:"port"`
	UserLevel uint32             `json:"userLevel"`
}

func (c *DNSOutboundConfig) Build() (proto.Message, error) {
	config := &dns.Config{
		Server: &net.Endpoint{
			Network: c.Network.Build(),
			Port:    uint32(c.Port),
		},
		UserLevel: c.UserLevel,
	}
	if c.Address != nil {
		config.Server.Address = c.Address.Build()
	}
	return config, nil
}
