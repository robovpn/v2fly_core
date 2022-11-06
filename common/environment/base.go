package environment

import (
	"github.com/robovpn/v2fly_core/common/environment/filesystemcap"
	"github.com/robovpn/v2fly_core/features/extension/storage"
	"github.com/robovpn/v2fly_core/transport/internet"
	"github.com/robovpn/v2fly_core/transport/internet/tagged"
)

type BaseEnvironmentCapabilitySet interface {
	FeaturesLookupCapabilitySet
	LogCapabilitySet
}

type BaseEnvironment interface {
	BaseEnvironmentCapabilitySet
	doNotImpl()
}

type SystemNetworkCapabilitySet interface {
	Dialer() internet.SystemDialer
	Listener() internet.SystemListener
}

type InstanceNetworkCapabilitySet interface {
	OutboundDialer() tagged.DialFunc
}

type FeaturesLookupCapabilitySet interface {
	RequireFeatures() interface{}
}

type LogCapabilitySet interface {
	RecordLog() interface{}
}

type FileSystemCapabilitySet interface {
	filesystemcap.FileSystemCapabilitySet
}

type PersistentStorageCapabilitySet interface {
	PersistentStorage() storage.ScopedPersistentStorage
}
type TransientStorageCapabilitySet interface {
	TransientStorage() storage.ScopedTransientStorage
}
