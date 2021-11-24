package health

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

func (s *tracingService) Liveness(ctx context.Context, req *LivenessRequest) (resp *LivenessResponse, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Liveness")
	defer span.Finish()
	return s.Service.Liveness(ctx, req)
}

func (s *tracingService) Readiness(ctx context.Context, req *ReadinessRequest) (resp *ReadinessResponse, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Readiness")
	defer span.Finish()
	return s.Service.Readiness(ctx, req)
}

func (s *tracingService) Version(ctx context.Context, req *VersionRequest) (resp *VersionResponse, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Version")
	defer span.Finish()
	return s.Service.Version(ctx, req)
}
