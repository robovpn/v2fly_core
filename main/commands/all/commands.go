package all

import (
	"github.com/robovpn/v2fly_core/main/commands/all/api"
	"github.com/robovpn/v2fly_core/main/commands/all/tls"
	"github.com/robovpn/v2fly_core/main/commands/base"
)

//go:generate go run github.com/robovpn/v2fly_core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		cmdLove,
		tls.CmdTLS,
		cmdUUID,
		cmdVerify,

		// documents
		docFormat,
		docMerge,
	)
}
