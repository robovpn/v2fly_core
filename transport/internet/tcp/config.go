package tcp

import (
	"github.com/robovpn/github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
