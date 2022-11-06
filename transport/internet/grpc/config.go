package grpc

import (
	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/transport/internet"
)

const protocolName = "gun"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
