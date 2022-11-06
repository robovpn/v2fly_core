//go:build !linux && !freebsd
// +build !linux,!freebsd

package tcp

import (
	"github.com/robovpn/github.com/robovpn/v2fly_core/common/net"
	"github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
