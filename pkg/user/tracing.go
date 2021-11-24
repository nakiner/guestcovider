package user

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/nakiner/guestcovider/tools/tracing"
)

// NewTracingService returns an instance of an instrumenting Service.
func NewTracingService(ctx context.Context, s Service) Service {
	tracer := tracing.FromContext(ctx)
	return &tracingService{tracer, s}
}

type tracingService struct {
	tracer opentracing.Tracer
	Service
}

func (s *tracingService) UpdateUser(ctx context.Context, req *UpdateUserRequest) (resp *UpdateUserResponse, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Version")
	defer span.Finish()
	return s.Service.UpdateUser(ctx, req)
}

func (s *tracingService) SearchUser(ctx context.Context, req *SearchUserRequest) (resp *SearchUserResponse, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Version")
	defer span.Finish()
	return s.Service.SearchUser(ctx, req)
}
