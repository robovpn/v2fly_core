package command_test

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/anypb"

	core "/v2fly_core"
	"/v2fly_core/app/dispatcher"
	"/v2fly_core/app/log"
	. "/v2fly_core/app/log/command"
	"/v2fly_core/app/proxyman"
	_ "/v2fly_core/app/proxyman/inbound"
	_ "/v2fly_core/app/proxyman/outbound"
	"/v2fly_core/common"
	"/v2fly_core/common/serial"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
