//go:build integration && !unit
// +build integration,!unit

package integration

import (
	"context"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/nakiner/guestcovider/pkg/user"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const grpcAddruser = "localhost:9194"

func TestGRPCUserServiceUpdateUser(t *testing.T) {

	conn, err := grpc.Dial(grpcAddruser, grpc.WithInsecure())
	if err != nil {
		t.Errorf("connection to grpc server: %s", err)
	}
	defer conn.Close()

	client := user.NewGRPCClient(conn, opentracing.GlobalTracer(), log.NewNopLogger())
	_, err = client.UpdateUser(context.Background(), &user.UpdateUserRequest{})

	assert.NoError(t, err)
}

func TestGRPCUserServiceSearchUser(t *testing.T) {

	conn, err := grpc.Dial(grpcAddruser, grpc.WithInsecure())
	if err != nil {
		t.Errorf("connection to grpc server: %s", err)
	}
	defer conn.Close()

	client := user.NewGRPCClient(conn, opentracing.GlobalTracer(), log.NewNopLogger())
	_, err = client.SearchUser(context.Background(), &user.SearchUserRequest{})

	assert.NoError(t, err)
}
