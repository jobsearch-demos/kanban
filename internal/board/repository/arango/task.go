package arango

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver"
	"jobsearch-demos/kanban/internal/board/entity"
	errors "jobsearch-demos/kanban/pkg/error"
)

type taskRepository struct {
	db             driver.Database
	collectionName string
}

func (t taskRepository) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	query := `INSERT { "title": @title, "description": @description, "created_at": @created_at, "updated_at": @updated_at } INTO tasks RETURN NEW`
	bindVars := map[string]interface{}{
		"title":       task.Name,
		"description": task.Description,
		"created_at":  task.CreatedAt, // TODO: convert to auto-generated timestamp
		"updated_at":  task.UpdatedAt, // TODO: convert to auto-generated time.Time
	}

	cursor, err := t.db.Query(ctx, query, bindVars)
	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error creating task: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	defer func() {
		if err := cursor.Close(); err != nil {
			// TODO: handle logging the error instead of panic
			panic(err)
		}
	}()

	var newTask entity.Task
	_, err = cursor.ReadDocument(ctx, &newTask)
	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error decoding the fetched results: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	return &newTask, nil
}

func (t taskRepository) Get(ctx context.Context, id string) (*entity.Task, error) {
	query := fmt.Sprintf("FOR t IN tasks FILTER t._key == '%s' RETURN t", id)
	cursor, err := t.db.Query(ctx, query, nil)

	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error fetching task: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	defer func() {
		if err := cursor.Close(); err != nil {
			// TODO: handle logging the error instead of panic
			panic(err)
		}
	}()

	var task entity.Task
	_, err = cursor.ReadDocument(ctx, &task)
	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error decoding the fetched results: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	return &task, nil
}

func (t taskRepository) List(ctx context.Context, filter interface{}) ([]*entity.Task, error) {
	query := "FOR t IN tasks RETURN t"
	cursor, err := t.db.Query(ctx, query, nil)
	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error fetching tasks: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	defer func() {
		if err := cursor.Close(); err != nil {
			// TODO: handle logging the error instead of panic
			panic(err)
		}
	}()

	var tasks []*entity.Task
	for {
		var task entity.Task
		_, err := cursor.ReadDocument(ctx, &task)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, errors.CustomError{
				Code:        500,
				Message:     fmt.Sprintf("error decoding the fetched results: %s\n", err.Error()),
				InfoContext: nil,
			}
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t taskRepository) Update(ctx context.Context, id string, task *entity.Task) (*entity.Task, error) {
	query := fmt.Sprintf("FOR t IN tasks FILTER t._key == '%s' UPDATE t WITH { title: @title, description: @description, updated_at: @updated_at } IN tasks RETURN NEW", id)
	bindVars := map[string]interface{}{
		"title":       task.Name,
		"description": task.Description,
		"updated_at":  task.UpdatedAt, // TODO: convert to auto-generated time.Time
	}

	cursor, err := t.db.Query(ctx, query, bindVars)
	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error updating task: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	defer func() {
		if err := cursor.Close(); err != nil {
			// TODO: handle logging the error instead of panic
			panic(err)
		}
	}()

	var updatedTask entity.Task
	_, err = cursor.ReadDocument(ctx, &updatedTask)
	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error decoding the fetched results: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	return &updatedTask, nil
}

func (t taskRepository) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("FOR t IN tasks FILTER t._key == '%s' REMOVE t IN tasks", id)
	_, err := t.db.Query(ctx, query, nil)
	if err != nil {
		return errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error deleting task: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	return nil
}

func NewTaskRepository(db driver.Database, collectionName string) entity.ITask {
	return &taskRepository{db: db, collectionName: collectionName}
}
