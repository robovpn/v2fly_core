package multiobservatory

import (
	"context"

	"google.golang.org/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"

	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/common/taggedfeatures"
	"github.com/robovpn/v2fly_core/features"
	"github.com/robovpn/v2fly_core/features/extension"
)

type Observer struct {
	features.TaggedFeatures
	config *Config
	ctx    context.Context
}

func (o Observer) GetObservation(ctx context.Context) (proto.Message, error) {
	return common.Must2(o.GetFeaturesByTag("")).(extension.Observatory).GetObservation(ctx)
}

func (o Observer) Type() interface{} {
	return extension.ObservatoryType()
}

func New(ctx context.Context, config *Config) (*Observer, error) {
	holder, err := taggedfeatures.NewHolderFromConfig(ctx, config.Holders, extension.ObservatoryType())
	if err != nil {
		return nil, err
	}
	return &Observer{config: config, ctx: ctx, TaggedFeatures: holder}, nil
}

func (x *Config) UnmarshalJSONPB(unmarshaler *jsonpb.Unmarshaler, bytes []byte) error {
	var err error
	x.Holders, err = taggedfeatures.LoadJSONConfig(context.TODO(), "service", "background", bytes)
	return err
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*Config))
	}))
}
