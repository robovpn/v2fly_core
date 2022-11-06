package commands

import (
	"fmt"

	core "../v2fly_core"
	"../v2fly_core/main/commands/base"
)

// CmdVersion prints v2fly Versions
var CmdVersion = &base.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "print v2fly version",
	Long: `Prints the build information for v2fly.
`,
	Run: executeVersion,
}

func executeVersion(cmd *base.Command, args []string) {
	printVersion()
}

func printVersion() {
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}
