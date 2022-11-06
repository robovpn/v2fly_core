package restfulapi

import (
	"context"
	"net"
	"sync"

	core "/v2fly_core"
	"/v2fly_core/features"
	feature_stats "/v2fly_core/features/stats"
)

//go:generate go run /v2fly_core/common/errors/errorgen

type restfulService struct {
	listener net.Listener
	config   *Config
	access   sync.Mutex

	stats feature_stats.Manager

	ctx context.Context
}

func (rs *restfulService) Type() interface{} {
	return (*struct{})(nil)
}

func (rs *restfulService) Start() error {
	defer rs.access.Unlock()
	rs.access.Lock()
	return rs.start()
}

func (rs *restfulService) Close() error {
	defer rs.access.Unlock()
	rs.access.Lock()
	if rs.listener != nil {
		return rs.listener.Close()
	}
	return nil
}

func (rs *restfulService) init(config *Config, stats feature_stats.Manager) {
	rs.stats = stats
	rs.config = config
}

func newRestfulService(ctx context.Context, config *Config) (features.Feature, error) {
	r := new(restfulService)
	r.ctx = ctx
	if err := core.RequireFeatures(ctx, func(stats feature_stats.Manager) {
		r.init(config, stats)
	}); err != nil {
		return nil, err
	}
	return r, nil
}
