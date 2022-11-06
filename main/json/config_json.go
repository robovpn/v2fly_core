package json

//go:generate go run github.com/robovpn/v2fly_core/common/errors/errorgen

import (
	"io"
	"os"

	"github.com/robovpn/v2fly_core"
	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/common/cmdarg"
	"github.com/robovpn/v2fly_core/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				r, err := confloader.LoadExtConfig(v, os.Stdin)
				if err != nil {
					return nil, newError("failed to execute v2ctl to convert config file.").Base(err).AtWarning()
				}
				return core.LoadConfig("protobuf", "", r)
			case io.Reader:
				r, err := confloader.LoadExtConfig([]string{"stdin:"}, os.Stdin)
				if err != nil {
					return nil, newError("failed to execute v2ctl to convert config file.").Base(err).AtWarning()
				}
				return core.LoadConfig("protobuf", "", r)
			default:
				return nil, newError("unknown type")
			}
		},
	}))
}
