package user

import (
	"context"
	"strconv"

	"github.com/getsentry/sentry-go"
)

func NewSentryService(s Service) Service {
	return &sentryService{s}
}

type sentryService struct {
	Service
}

type sentryLog interface {
	SentryLog() []interface{}
}

func (s *sentryService) getSentryLog(req interface{}, resp interface{}) (out map[string][]interface{}) {
	out = make(map[string][]interface{})
	if sentry, ok := interface{}(req).(sentryLog); ok {
		out["request"] = append(out["request"], sentry.SentryLog()...)
	}

	if sentry, ok := interface{}(resp).(sentryLog); ok {
		out["response"] = append(out["response"], sentry.SentryLog()...)
	}
	return
}

func (s *sentryService) UpdateUser(ctx context.Context, req *UpdateUserRequest) (resp *UpdateUserResponse, err error) {
	defer func() {
		if err != nil {
			log := s.getSentryLog(req, resp)
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetTag("code", strconv.Itoa(getHTTPStatusCode(err)))
				scope.SetTag("method", "UpdateUser")
				scope.SetExtra("request", log["request"])
				scope.SetExtra("response", log["response"])
			})
			sentry.CaptureException(err)
		}
	}()
	return s.Service.UpdateUser(ctx, req)
}

func (s *sentryService) SearchUser(ctx context.Context, req *SearchUserRequest) (resp *SearchUserResponse, err error) {
	defer func() {
		if err != nil {
			log := s.getSentryLog(req, resp)
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetTag("code", strconv.Itoa(getHTTPStatusCode(err)))
				scope.SetTag("method", "SearchUser")
				scope.SetExtra("request", log["request"])
				scope.SetExtra("response", log["response"])
			})
			sentry.CaptureException(err)
		}
	}()
	return s.Service.SearchUser(ctx, req)
}
