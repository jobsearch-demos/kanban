package entity

import "context"

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username" validate:"required"`
	Password string  `json:"password" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Groups   []Group `json:"groups" validate:"required"`
}

type IUser interface {
	Get(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, user User) (*User, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter interface{}) ([]User, error)
}
