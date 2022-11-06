package command_test

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/robovpn/v2fly_core"
	"github.com/robovpn/v2fly_core/app/dispatcher"
	"github.com/robovpn/v2fly_core/app/log"
	. "github.com/robovpn/v2fly_core/app/log/command"
	"github.com/robovpn/v2fly_core/app/proxyman"
	_ "github.com/robovpn/v2fly_core/app/proxyman/inbound"
	_ "github.com/robovpn/v2fly_core/app/proxyman/outbound"
	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/common/serial"
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
