package simplified

import (
	"context"

	"../v2fly_core/common"
	"../v2fly_core/common/protocol"
	"../v2fly_core/proxy/http"
)

func init() {
	common.Must(common.RegisterConfig((*ServerConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		simplifiedServer := config.(*ServerConfig)
		_ = simplifiedServer
		fullServer := &http.ServerConfig{}
		return common.CreateObject(ctx, fullServer)
	}))

	common.Must(common.RegisterConfig((*ClientConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		simplifiedClient := config.(*ClientConfig)
		fullClient := &http.ClientConfig{
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
