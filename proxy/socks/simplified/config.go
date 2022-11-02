package simplified

import (
	"context"

	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/common/protocol"
	"github.com/robovpn/v2fly_core/proxy/socks"
)

func init() {
	common.Must(common.RegisterConfig((*ServerConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		simplifiedServer := config.(*ServerConfig)
		fullServer := &socks.ServerConfig{
			AuthType:       socks.AuthType_NO_AUTH,
			Address:        simplifiedServer.Address,
			UdpEnabled:     simplifiedServer.UdpEnabled,
			PacketEncoding: simplifiedServer.PacketEncoding,
		}
		return common.CreateObject(ctx, fullServer)
	}))

	common.Must(common.RegisterConfig((*ClientConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		simplifiedClient := config.(*ClientConfig)
		fullClient := &socks.ClientConfig{
			Server: []*protocol.ServerEndpoint{
				{
					Address: simplifiedClient.Address,
					Port:    simplifiedClient.Port,
				},
			},
		}
		return common.CreateObject(ctx, fullClient)
	}))
}
