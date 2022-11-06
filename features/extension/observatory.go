package extension

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/robovpn/v2fly_core/features"
)

type Observatory interface {
	features.Feature
	GetObservation(ctx context.Context) (proto.Message, error)
}

func ObservatoryType() interface{} {
	return (*Observatory)(nil)
}
