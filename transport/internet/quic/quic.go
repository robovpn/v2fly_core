package quic

import (
	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/transport/internet"
)

//go:generate go run github.com/robovpn/v2fly_core/common/errors/errorgen

// Here is some modification needs to be done before update quic vendor.
// * use bytespool in buffer_pool.go
// * set MaxReceivePacketSize to 1452 - 32 (16 bytes auth, 16 bytes head)
//
//

const (
	protocolName   = "quic"
	internalDomain = "quic.internal.v2fly.org"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
