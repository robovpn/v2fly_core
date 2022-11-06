package blackhole_test

import (
	"context"
	"testing"

	"github.com/robovpn/github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/github.com/robovpn/v2fly_core/common/buf"
	"github.com/robovpn/github.com/robovpn/v2fly_core/common/serial"
	"github.com/robovpn/github.com/robovpn/v2fly_core/proxy/blackhole"
	"github.com/robovpn/github.com/robovpn/v2fly_core/transport"
	"github.com/robovpn/github.com/robovpn/v2fly_core/transport/pipe"
)

func TestBlackHoleHTTPResponse(t *testing.T) {
	handler, err := blackhole.New(context.Background(), &blackhole.Config{
		Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
	})
	common.Must(err)

	reader, writer := pipe.New(pipe.WithoutSizeLimit())

	readerError := make(chan error)
	var mb buf.MultiBuffer
	go func() {
		b, e := reader.ReadMultiBuffer()
		mb = b
		readerError <- e
	}()

	link := transport.Link{
		Reader: reader,
		Writer: writer,
	}
	common.Must(handler.Process(context.Background(), &link, nil))
	common.Must(<-readerError)
	if mb.IsEmpty() {
		t.Error("expect http response, but nothing")
	}
}
