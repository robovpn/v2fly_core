package drain

import "io"

//go:generate go run github.com/robovpn/v2fly_core/common/errors/errorgen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
