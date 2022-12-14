package v4_test

import (
	"testing"

	"github.com/robovpn/v2fly_core/common/serial"
	"github.com/robovpn/v2fly_core/infra/conf/cfgcommon"
	"github.com/robovpn/v2fly_core/infra/conf/cfgcommon/testassist"
	v4 "github.com/robovpn/v2fly_core/infra/conf/v4"
	"github.com/robovpn/v2fly_core/proxy/blackhole"
)

func TestHTTPResponseJSON(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.BlackholeConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"response": {
					"type": "http"
				}
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &blackhole.Config{
				Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
			},
		},
		{
			Input:  `{}`,
			Parser: testassist.LoadJSON(creator),
			Output: &blackhole.Config{},
		},
	})
}
