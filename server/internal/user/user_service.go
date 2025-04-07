package user

import (
	"context"
	"server/util"
	"time"

	"github.com/google/uuid"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUserService(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:        uuid.New(),
		UserName:  req.UserName,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	newUser, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:        newUser.ID,
		UserName:  newUser.UserName,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	return res, nil

}
