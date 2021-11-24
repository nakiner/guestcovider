package user

import (
	"context"
	"github.com/nakiner/guestcovider/internal/userRepository"
)

type userService struct {
	repo userRepository.Repository
}

func NewUserService(repo userRepository.Repository) Service {
	return &userService{repo: repo}
}

func (s *userService) UpdateUser(ctx context.Context, req *UpdateUserRequest) (resp *UpdateUserResponse, err error) {
	resp = &UpdateUserResponse{}

	if req.Id < 0 {
		return resp, ErrInvalidRequest
	}

	user := userRepository.User{
		ID: req.Id,
		Checkin: req.Data.Checkin,
		CovidPass: req.Data.CovidPass,
	}

	if err := s.repo.UpdateUser(ctx, &user); err != nil {
		return resp, err
	}

	resp.Status = &Status{
		Status: true,
		Message: "OK",
	}

	return resp, nil
}

func (s *userService) SearchUser(ctx context.Context, req *SearchUserRequest) (resp *SearchUserResponse, err error) {
	resp = &SearchUserResponse{}
	if len(req.Surname) < 1 {
		return resp, nil
	}

	users, err := s.repo.FindBySurname(ctx, req.Surname)
	if err != nil {
		return resp, err
	}

	resp.Status = &Status{
		Status: true,
		Message: "OK",
	}

	return resp.FromRepo(users), nil
}

func (pp *SearchUserResponse) FromRepo(in []*userRepository.User) *SearchUserResponse {
	if pp == nil {
		pp = new(SearchUserResponse)
	}

	for _, i := range in {
		pp.Data = append(pp.Data, User{
			Id:           i.ID,
			Status:       i.Status,
			Company:      i.Company,
			Surname:      i.Surname,
			Name:         i.Name,
			Guest:        i.Guest,
			CovidPass:    i.CovidPass,
			Rank:         i.Rank,
			ContactPhone: i.ContactPhone,
			ContactMail:  i.ContactMail,
			Checkin:      i.Checkin,
		})
	}

	return pp
}
