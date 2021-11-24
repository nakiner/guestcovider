package user

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"
	tool "github.com/nakiner/guestcovider/tools/metrics"
)

// NewMetricService returns an instance of an instrumenting Service.
func NewMetricsService(ctx context.Context, s Service) Service {
	counter, latency := tool.FromContext(ctx)
	return &metricService{counter, latency, s}
}

type metricService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func (s *metricService) UpdateUser(ctx context.Context, req *UpdateUserRequest) (resp *UpdateUserResponse, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "user", "handler", "UpdateUser", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "user", "handler", "UpdateUser", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.UpdateUser(ctx, req)
}

func (s *metricService) SearchUser(ctx context.Context, req *SearchUserRequest) (resp *SearchUserResponse, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "user", "handler", "SearchUser", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "user", "handler", "SearchUser", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.SearchUser(ctx, req)
}
