package protoext

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

//go:generate go run /v2fly_core/common/errors/errorgen

func GetMessageOptions(msgDesc protoreflect.MessageDescriptor) (*MessageOpt, error) {
	msgOpt := msgDesc.Options().(*descriptorpb.MessageOptions)
	msgOptRet, err := proto.GetExtension(msgOpt, E_MessageOpt)
	if err != nil {
		return nil, newError("unable to parse extension from message").Base(err)
	}
	return msgOptRet.(*MessageOpt), nil
}

func GetFieldOptions(fieldDesc protoreflect.FieldDescriptor) (*FieldOpt, error) {
	fieldOpt := fieldDesc.Options().(*descriptorpb.FieldOptions)
	msgOptRet, err := proto.GetExtension(fieldOpt, E_FieldOpt)
	if err != nil {
		return nil, newError("unable to parse extension from message").Base(err)
	}
	return msgOptRet.(*FieldOpt), nil
}
