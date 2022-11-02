package mux_test

import (
	"testing"

	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/common/buf"
	"github.com/robovpn/v2fly_core/common/mux"
	"github.com/robovpn/v2fly_core/common/net"
)

func BenchmarkFrameWrite(b *testing.B) {
	frame := mux.FrameMetadata{
		Target:        net.TCPDestination(net.DomainAddress("www.v2fly.org"), net.Port(80)),
		SessionID:     1,
		SessionStatus: mux.SessionStatusNew,
	}
	writer := buf.New()
	defer writer.Release()

	for i := 0; i < b.N; i++ {
		common.Must(frame.WriteTo(writer))
		writer.Clear()
	}
}
