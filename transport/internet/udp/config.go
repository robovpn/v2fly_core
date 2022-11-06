package udp

import (
	"/v2fly_core/common"
	"/v2fly_core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
