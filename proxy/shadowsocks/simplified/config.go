package simplified

import (
	"context"

	"../v2fly_core/common"
	"../v2fly_core/common/protocol"
	"../v2fly_core/common/serial"
	"../v2fly_core/proxy/shadowsocks"
)

func init() {
	common.Must(common.RegisterConfig((*ServerConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		simplifiedServer := config.(*ServerConfig)
		fullServer := &shadowsocks.ServerConfig{
			User: &protocol.User{
				Account: serial.ToTypedMessage(&shadowsocks.Account{
					Password:   simplifiedServer.Password,
					CipherType: simplifiedServer.Method,
				}),
			},
			Network:        simplifiedServer.Networks.GetNetwork(),
			PacketEncoding: simplifiedServer.PacketEncoding,
		}

		return common.CreateObject(ctx, fullServer)
	}))

	common.Must(common.RegisterConfig((*ClientConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		simplifiedClient := config.(*ClientConfig)
		fullClient := &shadowsocks.ClientConfig{
			Server: []*protocol.ServerEndpoint{
				{
					Address: simplifiedClient.Address,
					Port:    simplifiedClient.Port,
					User: []*protocol.User{
						{
							Account: serial.ToTypedMessage(&shadowsocks.Account{
								Password:                       simplifiedClient.Password,
								CipherType:                     simplifiedClient.Method,
								ExperimentReducedIvHeadEntropy: simplifiedClient.ExperimentReducedIvHeadEntropy,
							}),
						},
					},
				},
			},
		}

		return common.CreateObject(ctx, fullClient)
	}))
}
