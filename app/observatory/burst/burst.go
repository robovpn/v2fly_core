package burst

import (
	"math"
	"time"
)

//go:generate go run github.com/robovpn/github.com/robovpn/v2fly_core/common/errors/errorgen

const (
	rttFailed      = time.Duration(math.MaxInt64 - iota)
	rttUntested    // nolint: varcheck
	rttUnqualified // nolint: varcheck
)
