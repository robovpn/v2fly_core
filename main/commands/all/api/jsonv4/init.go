package jsonv4

import "/v2fly_core/main/commands/all/api"

func init() {
	api.CmdAPI.Commands = append(api.CmdAPI.Commands,
		cmdAddInbounds,
		cmdAddOutbounds,
		cmdRemoveInbounds,
		cmdRemoveOutbounds)
}
