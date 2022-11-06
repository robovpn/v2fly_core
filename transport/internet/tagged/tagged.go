package tagged

import (
	"context"

	"github.com/robovpn/github.com/robovpn/v2fly_core/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
