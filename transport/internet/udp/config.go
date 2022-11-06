package udp

import (
	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
