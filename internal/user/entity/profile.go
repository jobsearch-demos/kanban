package entity

import "context"

type Profile struct {
	ID          string `json:"id,omitempty" validate:"omitempty"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"numeric"`
	Photo       string `json:"photo" validate:"omitempty"`
}

type IProfile interface {
	Create(ctx context.Context, profile *Profile) (*Profile, error)
	Get(ctx context.Context, id string) (*Profile, error)
	List(ctx context.Context, filter interface{}) ([]*Profile, error)
	Update(ctx context.Context, id string, profile *Profile) (*Profile, error)
	Delete(ctx context.Context, id string) error
}
