package entity

import "context"

type Task struct {
	ID          string `json:"id"`
	ProfileID   string `json:"profile_id" validate:"required"`
	BoardID     string `json:"board_id" validate:"required"`
	ColumnID    string `json:"column_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Order       uint64 `json:"order" validate:"required"`
	CreatedAt   string `json:"created_at" validate:"required"`
	UpdatedAt   string `json:"updated_at" validate:"required"`
}

type ITask interface {
	List(ctx context.Context, filter interface{}) ([]*Task, error)
	Get(ctx context.Context, id string) (*Task, error)
	Create(ctx context.Context, task *Task) (*Task, error)
	Update(ctx context.Context, id string, task *Task) (*Task, error)
	Delete(ctx context.Context, id string) error
}
