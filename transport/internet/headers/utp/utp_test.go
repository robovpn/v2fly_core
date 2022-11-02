package utp_test

import (
	"context"
	"testing"

	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/common/buf"
	. "github.com/robovpn/v2fly_core/transport/internet/headers/utp"
)

func TestUTPWrite(t *testing.T) {
	content := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	utpRaw, err := New(context.Background(), &Config{})
	common.Must(err)

	utp := utpRaw.(*UTP)

	payload := buf.New()
	utp.Serialize(payload.Extend(utp.Size()))
	payload.Write(content)

	if payload.Len() != int32(len(content))+utp.Size() {
		t.Error("unexpected payload length: ", payload.Len())
	}
}
