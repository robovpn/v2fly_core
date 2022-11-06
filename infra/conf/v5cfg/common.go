package v5cfg

import (
	"context"
	"encoding/json"

	"github.com/golang/protobuf/proto"

	"github.com/robovpn/github.com/robovpn/v2fly_core/common/environment/envctx"
	"github.com/robovpn/github.com/robovpn/v2fly_core/common/environment/envimpl"
	"github.com/robovpn/github.com/robovpn/v2fly_core/common/registry"
)

func loadHeterogeneousConfigFromRawJSON(interfaceType, name string, rawJSON json.RawMessage) (proto.Message, error) {
	fsdef := envimpl.NewDefaultFileSystemDefaultImpl()
	ctx := envctx.ContextWithEnvironment(context.TODO(), fsdef)
	if len(rawJSON) == 0 {
		rawJSON = []byte("{}")
	}
	return registry.LoadImplementationByAlias(ctx, interfaceType, name, []byte(rawJSON))
}

// LoadHeterogeneousConfigFromRawJSON private API
func LoadHeterogeneousConfigFromRawJSON(ctx context.Context, interfaceType, name string, rawJSON json.RawMessage) (proto.Message, error) {
	return loadHeterogeneousConfigFromRawJSON(interfaceType, name, rawJSON)
}
