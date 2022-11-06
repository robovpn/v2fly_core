package jsonv4

import "github.com/robovpn/v2fly_core/main/commands/all/api"

func init() {
	api.CmdAPI.Commands = append(api.CmdAPI.Commands,
		cmdAddInbounds,
		cmdAddOutbounds,
		cmdRemoveInbounds,
		cmdRemoveOutbounds)
}
