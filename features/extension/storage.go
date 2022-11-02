package extension

import (
	"context"

	"github.com/robovpn/v2fly_core/features"
)

type PersistentStorageEngine interface {
	features.Feature
	PersistentStorageEngine()
	Put(ctx context.Context, key []byte, value []byte) error
	Get(ctx context.Context, key []byte) ([]byte, error)
	List(ctx context.Context, keyPrefix []byte) ([][]byte, error)
}
