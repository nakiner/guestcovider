//go:generate easyjson -all endpoint.go
package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	_ "github.com/mailru/easyjson/gen"
)

//easyjson:json
type SearchUserRequest struct {
	Surname string `json:"surname,omitempty"`
}

//easyjson:json
type SearchUserResponse struct {
	Status *Status `json:"status,omitempty"`
	Data   []User  `json:"data,omitempty"`
}

//easyjson:json
type Status struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

//easyjson:json
type UpdateData struct {
	CovidPass string `json:"covidPass,omitempty"`
	Checkin   bool   `json:"checkin,omitempty"`
}

//easyjson:json
type UpdateUserRequest struct {
	Id   uint64      `json:"id,omitempty"`
	Data *UpdateData `json:"data,omitempty"`
}

//easyjson:json
type UpdateUserResponse struct {
	Status *Status `json:"status,omitempty"`
}

//easyjson:json
type User struct {
	Id           uint64 `json:"id"`
	Status       string `json:"status"`
	Company      string `json:"company"`
	Surname      string `json:"surname"`
	Name         string `json:"name"`
	Guest        string `json:"guest"`
	CovidPass    string `json:"covidPass"`
	Rank         string `json:"rank,omitempty"`
	ContactPhone string `json:"contactPhone"`
	ContactMail  string `json:"contactMail"`
	Checkin      bool   `json:"checkin"`
}

//easyjson:skip
type endpoints struct {
	UpdateUserEndpoint endpoint.Endpoint
	SearchUserEndpoint endpoint.Endpoint
}

func (e endpoints) UpdateUser(ctx context.Context, req *UpdateUserRequest) (resp *UpdateUserResponse, err error) {
	response, err := e.UpdateUserEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(UpdateUserResponse)
	return &r, err
}

func (e endpoints) SearchUser(ctx context.Context, req *SearchUserRequest) (resp *SearchUserResponse, err error) {
	response, err := e.SearchUserEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(SearchUserResponse)
	return &r, err
}

func makeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		return s.UpdateUser(ctx, &req)
	}
}

func makeSearchUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchUserRequest)
		return s.SearchUser(ctx, &req)
	}
}
