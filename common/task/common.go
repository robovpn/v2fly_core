package task

import "github.com/robovpn/v2fly_core/common"

// Close returns a func() that closes v.
func Close(v interface{}) func() error {
	return func() error {
		return common.Close(v)
	}
}
