package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"jobsearch-demos/kanban/internal/board/entity"
	errors "jobsearch-demos/kanban/pkg/error"
)

type taskRepository struct {
	db             *mongo.Database
	collectionName string
}

func (t taskRepository) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	_, err := t.db.Collection(t.collectionName).InsertOne(ctx, task)
	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error creating task: %s\n", err.Error()),
			InfoContext: nil,
		}
	}
	return task, nil
}

func (t taskRepository) Get(ctx context.Context, id string) (*entity.Task, error) {
	task := &entity.Task{}

	err := t.db.Collection(t.collectionName).FindOne(ctx, bson.D{{"id", id}}).Decode(task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.CustomError{
				Code:        400,
				Message:     fmt.Sprintf("no task with provided id: %s\n", err.Error()),
				InfoContext: nil,
			}
		}
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error fetching task: %s\n", err.Error()),
			InfoContext: nil,
		}
	}
	return task, nil
}

func (t taskRepository) List(ctx context.Context, filter interface{}) ([]*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) Update(ctx context.Context, id string, task *entity.Task) (*entity.Task, error) {
	res := t.db.Collection(t.collectionName).FindOneAndUpdate(ctx, bson.D{{"id", id}}, bson.D{{"$set", task}})

	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.CustomError{
				Code:        400,
				Message:     fmt.Sprintf("no task with provided id: %s\n", err.Error()),
				InfoContext: nil,
			}
		}
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error updating task: %s\n", res.Err().Error()),
			InfoContext: nil,
		}
	}
	return task, nil
}

func (t taskRepository) Delete(ctx context.Context, id string) error {
	res, err := t.db.Collection(t.collectionName).DeleteOne(ctx, bson.D{{"id", id}})
	if err != nil {
		return errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error deleting task: %s\n", err.Error()),
			InfoContext: nil,
		}
	}
	if res.DeletedCount == 0 {
		return errors.CustomError{
			Code:        400,
			Message:     fmt.Sprintf("no task with provided id: %s\n", id),
			InfoContext: nil,
		}
	}
	return nil
}

func NewTaskRepository(db *mongo.Database) entity.ITask {
	return &taskRepository{db: db}
}
