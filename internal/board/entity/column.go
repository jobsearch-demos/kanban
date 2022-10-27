package entity

import "context"

type Column struct {
	ID          string `json:"id,omitempty" validate:"omitempty"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type IColumn interface {
	List(ctx context.Context, filter interface{}) ([]*Column, error)
	Get(ctx context.Context, id string) (*Column, error)
	Create(ctx context.Context, column *Column) (*Column, error)
	Update(ctx context.Context, id string, column *Column) (*Column, error)
	Delete(ctx context.Context, id string) error
}
