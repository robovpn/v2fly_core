package routing

import (
	"context"

	"github.com/robovpn/github.com/robovpn/v2fly_core/common/net"
	"github.com/robovpn/github.com/robovpn/v2fly_core/features"
	"github.com/robovpn/github.com/robovpn/v2fly_core/transport"
)

// Dispatcher is a feature that dispatches inbound requests to outbound handlers based on rules.
// Dispatcher is required to be registered in a v2fly instance to make v2fly function properly.
//
// v2fly:api:stable
type Dispatcher interface {
	features.Feature

	// Dispatch returns a Ray for transporting data for the given request.
	Dispatch(ctx context.Context, dest net.Destination) (*transport.Link, error)
}

// DispatcherType returns the type of Dispatcher interface. Can be used to implement common.HasType.
//
// v2fly:api:stable
func DispatcherType() interface{} {
	return (*Dispatcher)(nil)
}
