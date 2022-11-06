package internet_test

import (
	"testing"

	"../v2fly_core/common"
	. "../v2fly_core/transport/internet"
	"../v2fly_core/transport/internet/headers/noop"
	"../v2fly_core/transport/internet/headers/srtp"
	"../v2fly_core/transport/internet/headers/utp"
	"../v2fly_core/transport/internet/headers/wechat"
	"../v2fly_core/transport/internet/headers/wireguard"
)

func TestAllHeadersLoadable(t *testing.T) {
	testCases := []struct {
		Input interface{}
		Size  int32
	}{
		{
			Input: new(noop.Config),
			Size:  0,
		},
		{
			Input: new(srtp.Config),
			Size:  4,
		},
		{
			Input: new(utp.Config),
			Size:  4,
		},
		{
			Input: new(wechat.VideoConfig),
			Size:  13,
		},
		{
			Input: new(wireguard.WireguardConfig),
			Size:  4,
		},
	}

	for _, testCase := range testCases {
		header, err := CreatePacketHeader(testCase.Input)
		common.Must(err)
		if header.Size() != testCase.Size {
			t.Error("expected size ", testCase.Size, " but got ", header.Size())
		}
	}
}
