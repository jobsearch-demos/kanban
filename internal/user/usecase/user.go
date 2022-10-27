package usecase

import (
	"context"
	"jobsearch-demos/kanban/internal/user/entity"
	errors "jobsearch-demos/kanban/pkg/error"
)

type IUserService interface {
	// CreateUser a new user
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
	// GetUser a user by id
	GetUser(ctx context.Context, id int) (*entity.User, error)
	// ListUsers all users
	ListUsers(ctx context.Context) ([]entity.User, error)
	// UpdateUser a user
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
	// DeleteUser a user by id
	DeleteUser(ctx context.Context, id int) error
	// FilterUsers a user by id
	FilterUsers(ctx context.Context, user entity.User) ([]entity.User, error)
	// Login a user
	Login(ctx context.Context, username, password string) (*entity.User, error)
	// Logout a user
	Logout(ctx context.Context, user entity.User) (*entity.User, error)
	// ResetPassword a user
	ResetPassword(ctx context.Context, user entity.User) (*entity.User, error)
}

type userService struct {
	userRepo entity.IUser
}

func (u userService) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	res, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, errors.ToCustomError(err)
	}
	return res, nil
}

func (u userService) GetUser(ctx context.Context, id int) (*entity.User, error) {
	res, err := u.userRepo.Get(ctx, id)
	if err != nil {
		return nil, errors.ToCustomError(err)
	}
	return res, nil
}

func (u userService) ListUsers(ctx context.Context) ([]entity.User, error) {
	res, err := u.userRepo.List(ctx)
	if err != nil {
		return nil, errors.ToCustomError(err)
	}
	return res, nil
}

func (u userService) UpdateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	res, err := u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, errors.ToCustomError(err)
	}
	return res, nil
}

func (u userService) DeleteUser(ctx context.Context, id int) error {
	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		return errors.ToCustomError(err)
	}
	return nil
}

func (u userService) FilterUsers(ctx context.Context, user entity.User) ([]entity.User, error) {
	res, err := u.userRepo.Filter(ctx, user)
	if err != nil {
		return nil, errors.ToCustomError(err)
	}
	return res, nil
}

func (u userService) Login(ctx context.Context, username, password string) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u userService) Logout(ctx context.Context, user entity.User) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u userService) ResetPassword(ctx context.Context, user entity.User) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func NewUserService(userRepo entity.IUser) IUserService {
	return &userService{userRepo: userRepo}
}
