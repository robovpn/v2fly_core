package tls

import (
	"github.com/robovpn/github.com/robovpn/v2fly_core/main/commands/base"
)

// CmdTLS holds all tls sub commands
var CmdTLS = &base.Command{
	UsageLine: "{{.Exec}} tls",
	Short:     "TLS tools",
	Long: `{{.Exec}} {{.LongName}} provides tools for TLS.
	`,

	Commands: []*base.Command{
		cmdCert,
		cmdPing,
	},
}
