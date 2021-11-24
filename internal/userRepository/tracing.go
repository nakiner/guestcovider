package userRepository

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/nakiner/guestcovider/tools/tracing"
)

func NewTracingRepository(ctx context.Context, r Repository) Repository {
	tracer := tracing.FromContext(ctx)
	return &tracingRepository{tracer, r}
}

type tracingRepository struct {
	tracer opentracing.Tracer
	Repository
}

func (r *tracingRepository) FindBySurname(ctx context.Context, surname string) ([]*User, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, r.tracer, "FindBySurname")
	defer span.Finish()
	return r.Repository.FindBySurname(ctx, surname)
}

func (r *tracingRepository) UpdateUser(ctx context.Context, data *User) error {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, r.tracer, "UpdateUser")
	defer span.Finish()
	return r.Repository.UpdateUser(ctx, data)
}
