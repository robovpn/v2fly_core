package v4_test

import (
	"testing"

	"../v2fly_core/infra/conf/cfgcommon"
	"../v2fly_core/infra/conf/cfgcommon/testassist"
	v4 "../v2fly_core/infra/conf/v4"
	"../v2fly_core/proxy/http"
)

func TestHTTPServerConfig(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.HTTPServerConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"timeout": 10,
				"accounts": [
					{
						"user": "my-username",
						"pass": "my-password"
					}
				],
				"allowTransparent": true,
				"userLevel": 1
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &http.ServerConfig{
				Accounts: map[string]string{
					"my-username": "my-password",
				},
				AllowTransparent: true,
				UserLevel:        1,
				Timeout:          10,
			},
		},
	})
}
