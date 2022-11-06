package environment

import "github.com/robovpn/github.com/robovpn/v2fly_core/common/log"

type ConnectionCapabilitySet interface {
	ConnectionLogCapabilitySet
}

type ConnectionEnvironment interface {
	ConnectionCapabilitySet
	doNotImpl()
}

type ConnectionLogCapabilitySet interface {
	RecordConnectionLog(msg log.Message)
}
