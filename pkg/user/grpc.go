package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	pb "github.com/nakiner/guestcovider/internal/guestcoviderpb"
	"github.com/nakiner/guestcovider/tools/logging"
	"github.com/nakiner/guestcovider/tools/tracing"
	googlegrpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type grpcServer struct {
	updateUser grpctransport.Handler
	searchUser grpctransport.Handler
}

type ContextGRPCKey struct{}

type GRPCInfo struct{}

// NewGRPCServer makes a set of endpoints available as a gRPC userServer.
func NewGRPCServer(ctx context.Context, s Service) pb.UserServiceServer {
	logger := logging.FromContext(ctx)
	logger = log.With(logger, "grpc handler", "user")
	tracer := tracing.FromContext(ctx)

	options := []grpctransport.ServerOption{
		// grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(grpcToContext()),
		grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "grpc server", logger)),
		grpctransport.ServerFinalizer(closeGRPCTracer()),
	}

	return &grpcServer{
		updateUser: grpctransport.NewServer(
			makeUpdateUserEndpoint(s),
			decodeGRPCUpdateUserRequest,
			encodeGRPCUpdateUserResponse,
			options...,
		),
		searchUser: grpctransport.NewServer(
			makeSearchUserEndpoint(s),
			decodeGRPCSearchUserRequest,
			encodeGRPCSearchUserResponse,
			options...,
		),
	}
}

func JoinGRPC(ctx context.Context, s Service) func(*googlegrpc.Server) {
	return func(g *googlegrpc.Server) {
		pb.RegisterUserServiceServer(g, NewGRPCServer(ctx, s))
	}
}

func grpcToContext() grpc.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		return context.WithValue(ctx, ContextGRPCKey{}, GRPCInfo{})
	}
}

func closeGRPCTracer() grpc.ServerFinalizerFunc {
	return func(ctx context.Context, err error) {
		span := stdopentracing.SpanFromContext(ctx)
		span.Finish()
	}
}

func (s *grpcServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	_, rep, err := s.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserResponse), nil
}

func (s *grpcServer) SearchUser(ctx context.Context, req *pb.SearchUserRequest) (*pb.SearchUserResponse, error) {
	_, rep, err := s.searchUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SearchUserResponse), nil
}

func decodeGRPCUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.UpdateUserRequest)
	if !ok {
		return nil, errors.New("decodeGRPCUpdateUserRequest wrong request")
	}

	req := PBToUpdateUserRequest(inReq)
	if err := validate(req); err != nil {
		return nil, err
	}
	return *req, nil
}

func decodeGRPCSearchUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.SearchUserRequest)
	if !ok {
		return nil, errors.New("decodeGRPCSearchUserRequest wrong request")
	}

	req := PBToSearchUserRequest(inReq)
	if err := validate(req); err != nil {
		return nil, err
	}
	return *req, nil
}

func encodeGRPCUpdateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*UpdateUserResponse)
	if !ok {
		return nil, errors.New("encodeGRPCUpdateUserResponse wrong response")
	}

	return UpdateUserResponseToPB(inResp), nil
}

func encodeGRPCSearchUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*SearchUserResponse)
	if !ok {
		return nil, errors.New("encodeGRPCSearchUserResponse wrong response")
	}

	return SearchUserResponseToPB(inResp), nil
}

func SearchUserRequestToPB(d *SearchUserRequest) *pb.SearchUserRequest {
	if d == nil {
		return nil
	}

	resp := pb.SearchUserRequest{
		Surname: d.Surname,
	}

	return &resp
}

func PBToSearchUserRequest(d *pb.SearchUserRequest) *SearchUserRequest {
	if d == nil {
		return nil
	}

	resp := SearchUserRequest{
		Surname: d.Surname,
	}

	return &resp
}

func SearchUserResponseToPB(d *SearchUserResponse) *pb.SearchUserResponse {
	if d == nil {
		return nil
	}

	resp := pb.SearchUserResponse{
		Status: StatusToPB(d.Status),
	}

	for _, v := range d.Data {
		resp.Data = append(resp.Data, UserToPB(&v))
	}

	return &resp
}

func PBToSearchUserResponse(d *pb.SearchUserResponse) *SearchUserResponse {
	if d == nil {
		return nil
	}

	resp := SearchUserResponse{
		Status: PBToStatus(d.Status),
	}

	for _, v := range d.Data {
		resp.Data = append(resp.Data, *PBToUser(v))
	}

	return &resp
}

func StatusToPB(d *Status) *pb.Status {
	if d == nil {
		return nil
	}

	resp := pb.Status{
		Status:  d.Status,
		Message: d.Message,
	}

	return &resp
}

func PBToStatus(d *pb.Status) *Status {
	if d == nil {
		return nil
	}

	resp := Status{
		Status:  d.Status,
		Message: d.Message,
	}

	return &resp
}

func UpdateDataToPB(d *UpdateData) *pb.UpdateData {
	if d == nil {
		return nil
	}

	resp := pb.UpdateData{
		CovidPass: d.CovidPass,
		Checkin:   d.Checkin,
	}

	return &resp
}

func PBToUpdateData(d *pb.UpdateData) *UpdateData {
	if d == nil {
		return nil
	}

	resp := UpdateData{
		CovidPass: d.CovidPass,
		Checkin:   d.Checkin,
	}

	return &resp
}

func UpdateUserRequestToPB(d *UpdateUserRequest) *pb.UpdateUserRequest {
	if d == nil {
		return nil
	}

	resp := pb.UpdateUserRequest{
		Id:   d.Id,
		Data: UpdateDataToPB(d.Data),
	}

	return &resp
}

func PBToUpdateUserRequest(d *pb.UpdateUserRequest) *UpdateUserRequest {
	if d == nil {
		return nil
	}

	resp := UpdateUserRequest{
		Id:   d.Id,
		Data: PBToUpdateData(d.Data),
	}

	return &resp
}

func UpdateUserResponseToPB(d *UpdateUserResponse) *pb.UpdateUserResponse {
	if d == nil {
		return nil
	}

	resp := pb.UpdateUserResponse{
		Status: StatusToPB(d.Status),
	}

	return &resp
}

func PBToUpdateUserResponse(d *pb.UpdateUserResponse) *UpdateUserResponse {
	if d == nil {
		return nil
	}

	resp := UpdateUserResponse{
		Status: PBToStatus(d.Status),
	}

	return &resp
}

func UserToPB(d *User) *pb.User {
	if d == nil {
		return nil
	}

	resp := pb.User{
		Id:           d.Id,
		Status:       d.Status,
		Company:      d.Company,
		Surname:      d.Surname,
		Name:         d.Name,
		Guest:        d.Guest,
		CovidPass:    d.CovidPass,
		Rank:         d.Rank,
		ContactPhone: d.ContactPhone,
		ContactMail:  d.ContactMail,
		Checkin:      d.Checkin,
	}

	return &resp
}

func PBToUser(d *pb.User) *User {
	if d == nil {
		return nil
	}

	resp := User{
		Id:           d.Id,
		Status:       d.Status,
		Company:      d.Company,
		Surname:      d.Surname,
		Name:         d.Name,
		Guest:        d.Guest,
		CovidPass:    d.CovidPass,
		Rank:         d.Rank,
		ContactPhone: d.ContactPhone,
		ContactMail:  d.ContactMail,
		Checkin:      d.Checkin,
	}

	return &resp
}
