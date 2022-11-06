package tls_test

import (
	"context"
	"testing"

	"../v2fly_core/common"
	"../v2fly_core/common/buf"
	. "../v2fly_core/transport/internet/headers/tls"
)

func TestDTLSWrite(t *testing.T) {
	content := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	dtlsRaw, err := New(context.Background(), &PacketConfig{})
	common.Must(err)

	dtls := dtlsRaw.(*DTLS)

	payload := buf.New()
	dtls.Serialize(payload.Extend(dtls.Size()))
	payload.Write(content)

	if payload.Len() != int32(len(content))+dtls.Size() {
		t.Error("payload len: ", payload.Len(), " want ", int32(len(content))+dtls.Size())
	}
}
