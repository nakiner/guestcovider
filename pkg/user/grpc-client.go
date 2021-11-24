package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	pb "github.com/nakiner/guestcovider/internal/guestcoviderpb"
	"google.golang.org/grpc"
)

// NewGRPCClient returns an Service backed by a gRPC server at the other end
// of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func NewGRPCClient(conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) Service {
	// global client middlewares
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	}

	return endpoints{
		// Each individual endpoint is an grpc/transport.Client (which implements
		// endpoint.Endpoint) that gets wrapped with various middlewares. If you
		// made your own client library, you'd do this work there, so your server
		// could rely on a consistent set of client behavior.
		UpdateUserEndpoint: grpctransport.NewClient(
			conn,
			"guestcoviderpb.UserService",
			"UpdateUser",
			encodeGRPCUpdateUserRequest,
			decodeGRPCUpdateUserResponse,
			pb.UpdateUserResponse{},
			options...,
		).Endpoint(),
		SearchUserEndpoint: grpctransport.NewClient(
			conn,
			"guestcoviderpb.UserService",
			"SearchUser",
			encodeGRPCSearchUserRequest,
			decodeGRPCSearchUserResponse,
			pb.SearchUserResponse{},
			options...,
		).Endpoint(),
	}
}

func encodeGRPCSearchUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*SearchUserRequest)
	if !ok {
		return nil, errors.New("encodeGRPCSearchUserRequest wrong request")
	}

	return SearchUserRequestToPB(inReq), nil
}

func encodeGRPCUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*UpdateUserRequest)
	if !ok {
		return nil, errors.New("encodeGRPCUpdateUserRequest wrong request")
	}

	return UpdateUserRequestToPB(inReq), nil
}

func decodeGRPCSearchUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*pb.SearchUserResponse)
	if !ok {
		return nil, errors.New("decodeGRPCSearchUserResponse wrong response")
	}

	resp := PBToSearchUserResponse(inResp)

	return *resp, nil
}

func decodeGRPCUpdateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*pb.UpdateUserResponse)
	if !ok {
		return nil, errors.New("decodeGRPCUpdateUserResponse wrong response")
	}

	resp := PBToUpdateUserResponse(inResp)

	return *resp, nil
}
