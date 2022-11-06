package all

import (
	"/v2fly_core/main/commands/all/api"
	"/v2fly_core/main/commands/all/tls"
	"/v2fly_core/main/commands/base"
)

//go:generate go run v2fly_core/common/errors/errorgen

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
