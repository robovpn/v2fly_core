package extension

import (
	"context"

	"github.com/robovpn/github.com/robovpn/v2fly_core/features"
)

// InstanceManagement : unstable
type InstanceManagement interface {
	features.Feature
	ListInstance(ctx context.Context) ([]string, error)
	AddInstance(ctx context.Context, name string, config []byte, configType string) error
	StartInstance(ctx context.Context, name string) error
	StopInstance(ctx context.Context, name string) error
	UntrackInstance(ctx context.Context, name string) error
}

func InstanceManagementType() interface{} {
	return (*InstanceManagement)(nil)
}
