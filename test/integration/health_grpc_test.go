//go:build integration && !unit
// +build integration,!unit

package integration

import (
	"context"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/nakiner/guestcovider/pkg/health"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const grpcAddrhealth = "localhost:9194"

func TestGRPCHealthServiceLiveness(t *testing.T) {

	conn, err := grpc.Dial(grpcAddrhealth, grpc.WithInsecure())
	if err != nil {
		t.Errorf("connection to grpc server: %s", err)
	}
	defer conn.Close()

	client := health.NewGRPCClient(conn, opentracing.GlobalTracer(), log.NewNopLogger())
	_, err = client.Liveness(context.Background(), &health.LivenessRequest{})

	assert.NoError(t, err)
}

func TestGRPCHealthServiceReadiness(t *testing.T) {

	conn, err := grpc.Dial(grpcAddrhealth, grpc.WithInsecure())
	if err != nil {
		t.Errorf("connection to grpc server: %s", err)
	}
	defer conn.Close()

	client := health.NewGRPCClient(conn, opentracing.GlobalTracer(), log.NewNopLogger())
	_, err = client.Readiness(context.Background(), &health.ReadinessRequest{})

	assert.NoError(t, err)
}

func TestGRPCHealthServiceVersion(t *testing.T) {

	conn, err := grpc.Dial(grpcAddrhealth, grpc.WithInsecure())
	if err != nil {
		t.Errorf("connection to grpc server: %s", err)
	}
	defer conn.Close()

	client := health.NewGRPCClient(conn, opentracing.GlobalTracer(), log.NewNopLogger())
	_, err = client.Version(context.Background(), &health.VersionRequest{})

	assert.NoError(t, err)
}
