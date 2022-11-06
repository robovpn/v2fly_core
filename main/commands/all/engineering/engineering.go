package engineering

import "github.com/robovpn/github.com/robovpn/v2fly_core/main/commands/base"

//go:generate go run github.com/robovpn/github.com/robovpn/v2fly_core/common/errors/errorgen

var cmdEngineering = &base.Command{
	UsageLine: "{{.Exec}} engineering",
	Commands: []*base.Command{
		cmdConvertPb,
		cmdReversePb,
	},
}

func init() {
	base.RegisterCommand(cmdEngineering)
}
