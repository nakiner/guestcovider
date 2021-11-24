//go:generate mockgen -destination service_mock.go -package user  github.com/nakiner/guestcovider/pkg/user Service
package user

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
)

type Service interface {
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	SearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error)
}
