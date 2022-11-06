package jsonpb

import (
	"bytes"
	"io"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	core "../v2fly_core"
	"../v2fly_core/common"
	"../v2fly_core/common/buf"
	"../v2fly_core/common/cmdarg"
	"../v2fly_core/common/serial"
)

//go:generate go run ../v2fly_core/common/errors/errorgen

func loadJSONPB(data io.Reader) (*core.Config, error) {
	coreconf := &core.Config{}
	jsonpbloader := &jsonpb.Unmarshaler{AnyResolver: serial.GetResolver()}
	err := jsonpbloader.Unmarshal(data, coreconf)
	if err != nil {
		return nil, err
	}
	return coreconf, nil
}

func dumpJSONPb(config proto.Message, w io.Writer) error {
	jsonpbdumper := &jsonpb.Marshaler{AnyResolver: serial.GetResolver()}
	err := jsonpbdumper.Marshal(w, config)
	if err != nil {
		return err
	}
	return nil
}

func DumpJSONPb(config proto.Message, w io.Writer) error {
	return dumpJSONPb(config, w)
}

const FormatProtobufJSONPB = "jsonpb"

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      []string{FormatProtobufJSONPB},
		Extension: []string{".pb.json", ".pbjson"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case string:
				r, err := cmdarg.LoadArg(v)
				if err != nil {
					return nil, err
				}
				data, err := buf.ReadAllToBytes(r)
				if err != nil {
					return nil, err
				}
				return loadJSONPB(bytes.NewReader(data))
			case []byte:
				return loadJSONPB(bytes.NewReader(v))
			case io.Reader:
				data, err := buf.ReadAllToBytes(v)
				if err != nil {
					return nil, err
				}
				return loadJSONPB(bytes.NewReader(data))
			default:
				return nil, newError("unknown type")
			}
		},
	}))
}
