package entity

import "context"

type Group struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type IGroup interface {
	Create(ctx context.Context, group *Group) (*Group, error)
	Get(ctx context.Context, id string) (*Group, error)
	List(ctx context.Context, filter interface{}) ([]*Group, error)
	Update(ctx context.Context, id string, group *Group) (*Group, error)
	Delete(ctx context.Context, id string) error
}
