//go:build !confonly
// +build !confonly

package command

//go:generate go run github.com/robovpn/v2fly_core/common/errors/errorgen

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/grpc"

	core "github.com/robovpn/v2fly_core"
	"github.com/robovpn/v2fly_core/app/observatory"
	"github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/v2fly_core/features"
	"github.com/robovpn/v2fly_core/features/extension"
)

type service struct {
	UnimplementedObservatoryServiceServer
	v *core.Instance

	observatory extension.Observatory
}

func (s *service) GetOutboundStatus(ctx context.Context, request *GetOutboundStatusRequest) (*GetOutboundStatusResponse, error) {
	var result proto.Message
	if request.Tag == "" {
		observeResult, err := s.observatory.GetObservation(ctx)
		if err != nil {
			return nil, newError("cannot get observation").Base(err)
		}
		result = observeResult
	} else {
		fet, err := s.observatory.(features.TaggedFeatures).GetFeaturesByTag(request.Tag)
		if err != nil {
			return nil, newError("cannot get tagged observatory").Base(err)
		}
		observeResult, err := fet.(extension.Observatory).GetObservation(ctx)
		if err != nil {
			return nil, newError("cannot get observation").Base(err)
		}
		result = observeResult
	}
	retdata := result.(*observatory.ObservationResult)
	return &GetOutboundStatusResponse{
		Status: retdata,
	}, nil
}

func (s *service) Register(server *grpc.Server) {
	RegisterObservatoryServiceServer(server, s)
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, cfg interface{}) (interface{}, error) {
		s := core.MustFromContext(ctx)
		sv := &service{v: s}
		err := s.RequireFeatures(func(Observatory extension.Observatory) {
			sv.observatory = Observatory
		})
		if err != nil {
			return nil, err
		}
		return sv, nil
	}))
}
