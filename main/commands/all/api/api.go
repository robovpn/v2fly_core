package api

import (
	"../v2fly_core/main/commands/base"
)

// CmdAPI calls an API in an v2fly process
var CmdAPI = &base.Command{
	UsageLine: "{{.Exec}} api",
	Short:     "call v2fly API",
	Long: `{{.Exec}} {{.LongName}} provides tools to manipulate v2fly via its API.
`,
	Commands: []*base.Command{
		cmdLog,
		cmdStats,
		cmdBalancerInfo,
		cmdBalancerOverride,
	},
}
