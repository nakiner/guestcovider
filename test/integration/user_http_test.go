//go:build integration && !unit
// +build integration,!unit

package integration

import (
	"context"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/nakiner/guestcovider/pkg/user"
)

const htttAddruser = "localhost:8081"

func TestHTTPUserServiceUpdateUser(t *testing.T) {
	client, err := user.NewHTTPClient(htttAddruser, opentracing.GlobalTracer(), log.NewNopLogger())
	assert.NoError(t, err)
	_, err = client.UpdateUser(context.Background(), &user.UpdateUserRequest{})
	assert.NoError(t, err)
}

func TestHTTPUserServiceSearchUser(t *testing.T) {
	client, err := user.NewHTTPClient(htttAddruser, opentracing.GlobalTracer(), log.NewNopLogger())
	assert.NoError(t, err)
	_, err = client.SearchUser(context.Background(), &user.SearchUserRequest{})
	assert.NoError(t, err)
}
