package user

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/nakiner/guestcovider/tools/logging"
)

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(ctx context.Context, s Service) Service {
	logger := logging.FromContext(ctx)
	logger = log.With(logger, "component", "user")
	return &loggingService{logger, s}
}

type logged interface {
	Log() []interface{}
}

type loggingService struct {
	logger log.Logger
	Service
}

func (s *loggingService) getLog(req interface{}, resp interface{}) (out []interface{}) {
	if logger, ok := interface{}(req).(logged); ok {
		out = append(out, logger.Log()...)
	}

	if logger, ok := interface{}(resp).(logged); ok {
		out = append(out, logger.Log()...)
	}

	return
}

func (s *loggingService) UpdateUser(ctx context.Context, req *UpdateUserRequest) (resp *UpdateUserResponse, err error) {
	defer func(begin time.Time) {
		m := getInfoFromContext(ctx)
		m = append(m,
			"code", getHTTPStatusCode(err),
			"method", "UpdateUser",
			"took", time.Since(begin),
		)

		m = append(m, s.getLog(req, resp)...)

		if getHTTPStatusCode(err) == 404 {
			m = append(m, "msg", err)
			level.Warn(s.logger).Log(m...)
		} else if err != nil {
			m = append(m, "err", err)
			level.Error(s.logger).Log(m...)
		} else {
			level.Info(s.logger).Log(m...)
		}
	}(time.Now())
	return s.Service.UpdateUser(ctx, req)
}

func (s *loggingService) SearchUser(ctx context.Context, req *SearchUserRequest) (resp *SearchUserResponse, err error) {
	defer func(begin time.Time) {
		m := getInfoFromContext(ctx)
		m = append(m,
			"code", getHTTPStatusCode(err),
			"method", "SearchUser",
			"took", time.Since(begin),
		)

		m = append(m, s.getLog(req, resp)...)

		if getHTTPStatusCode(err) == 404 {
			m = append(m, "msg", err)
			level.Warn(s.logger).Log(m...)
		} else if err != nil {
			m = append(m, "err", err)
			level.Error(s.logger).Log(m...)
		} else {
			level.Info(s.logger).Log(m...)
		}
	}(time.Now())
	return s.Service.SearchUser(ctx, req)
}

func getInfoFromContext(ctx context.Context) []interface{} {
	m := make([]interface{}, 0)
	{
		val := ctx.Value(ContextGRPCKey{})
		if _, ok := val.(GRPCInfo); ok {
			m = append(m, "protocol", "GRPC")
		}
	}

	{
		val := ctx.Value(ContextHTTPKey{})
		if i, ok := val.(HTTPInfo); ok {
			m = append(m,
				// "protocol", i.Protocol,
				// "http_method", i.Method,
				// "from", i.From,
				"url", i.URL,
			)
		}
	}

	return m
}
