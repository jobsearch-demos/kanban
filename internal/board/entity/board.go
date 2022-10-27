package entity

import (
	"context"
	user "jobsearch-demos/kanban/internal/user/entity"
)

type Board struct {
	ID          string      `json:"id,omitempty" validate:"omitempty"`
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Members     []user.User `json:"members,omitempty" validate:"omitempty"`
	Columns     []Column    `json:"columns,omitempty" validate:"omitempty"`
}

type IBoard interface {
	List(ctx context.Context, filter interface{}) ([]*Board, error)
	Get(ctx context.Context, id string) (*Board, error)
	Create(ctx context.Context, board *Board) (*Board, error)
	Update(ctx context.Context, id string, board *Board) (*Board, error)
	Delete(ctx context.Context, id string) error
}
