package wechat_test

import (
	"context"
	"testing"

	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/common/buf"
	. "github.com/robovpn/v2fly_core/transport/internet/headers/wechat"
)

func TestUTPWrite(t *testing.T) {
	videoRaw, err := NewVideoChat(context.Background(), &VideoConfig{})
	common.Must(err)

	video := videoRaw.(*VideoChat)

	payload := buf.New()
	video.Serialize(payload.Extend(video.Size()))

	if payload.Len() != video.Size() {
		t.Error("expected payload size ", video.Size(), " but got ", payload.Len())
	}
}
