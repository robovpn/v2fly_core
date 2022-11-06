package engineering

import "../v2fly_core/main/commands/base"

//go:generate go run ../v2fly_core/common/errors/errorgen

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
