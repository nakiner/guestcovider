package health

import (
	"context"

	"github.com/bxcodec/faker"
)

var (
	BuildTime string
	Commit    string
	Version   string
)

type healthService struct {
}

func NewHealthService() Service {
	return &healthService{}
}

func (s *healthService) Liveness(ctx context.Context, req *LivenessRequest) (resp *LivenessResponse, err error) {
	a := LivenessResponse{}
	err = faker.FakeData(&a)
	if err != nil {
		return &a, err
	}
	return &a, nil
}

func (s *healthService) Readiness(ctx context.Context, req *ReadinessRequest) (resp *ReadinessResponse, err error) {
	a := ReadinessResponse{}
	err = faker.FakeData(&a)
	if err != nil {
		return &a, err
	}
	return &a, nil
}

func (s *healthService) Version(ctx context.Context, req *VersionRequest) (resp *VersionResponse, err error) {
	a := VersionResponse{}
	err = faker.FakeData(&a)
	if err != nil {
		return &a, err
	}
	return &a, nil
}
