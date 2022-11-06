package udp

import (
	"context"

	"../v2fly_core/common"
	"../v2fly_core/common/buf"
	"../v2fly_core/common/net"
)

type DispatcherI interface {
	common.Closable
	Dispatch(ctx context.Context, destination net.Destination, payload *buf.Buffer)
}
