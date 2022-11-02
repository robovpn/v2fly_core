package testassist

import (
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/infra/conf/cfgcommon"
)

func LoadJSON(creator func() cfgcommon.Buildable) func(string) (proto.Message, error) {
	return func(s string) (proto.Message, error) {
		instance := creator()
		if err := json.Unmarshal([]byte(s), instance); err != nil {
			return nil, err
		}
		return instance.Build()
	}
}

type TestCase struct {
	Input  string
	Parser func(string) (proto.Message, error)
	Output proto.Message
}

func RunMultiTestCase(t *testing.T, testCases []TestCase) {
	for _, testCase := range testCases {
		actual, err := testCase.Parser(testCase.Input)
		common.Must(err)
		if !proto.Equal(actual, testCase.Output) {
			t.Fatalf("Failed in test case:\n%s\nActual:\n%v\nExpected:\n%v", testCase.Input, actual, testCase.Output)
		}
	}
}
