package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"jobsearch-demos/kanban/internal/user/entity"
	errors "jobsearch-demos/kanban/pkg/error"
)

type userRepository struct {
	db             *mongo.Database
	collectionName string
}

func (u userRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User
	err := u.db.Collection(u.collectionName).FindOne(ctx, bson.D{{"id", id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.CustomError{
				Code:        400,
				Message:     fmt.Sprintf("no user with provided id: %s\n", err.Error()),
				InfoContext: nil,
			}
		}
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error fetching user: %s\n", err.Error()),
			InfoContext: nil,
		}
	}
	return &user, nil
}

func (u userRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	_, err := u.db.Collection(u.collectionName).InsertOne(ctx, user)
	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error creating user: %s\n", err.Error()),
			InfoContext: nil,
		}
	}
	return &user, nil
}

func (u userRepository) Update(ctx context.Context, user entity.User) (*entity.User, error) {
	res := u.db.Collection(u.collectionName).FindOneAndReplace(ctx, bson.D{{"id", user.ID}}, user)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.CustomError{
				Code:        400,
				Message:     fmt.Sprintf("no user with provided id: %s\n", err.Error()),
				InfoContext: nil,
			}
		}
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error updating user: %s\n", res.Err().Error()),
			InfoContext: nil,
		}
	}
	return &user, nil
}

func (u userRepository) Delete(ctx context.Context, id int) error {
	res := u.db.Collection(u.collectionName).FindOneAndDelete(ctx, bson.D{{"id", id}})
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.CustomError{
				Code:        400,
				Message:     fmt.Sprintf("no user with provided id: %s\n", err.Error()),
				InfoContext: nil,
			}
		}
		return errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error deleting user: %s\n", res.Err().Error()),
			InfoContext: nil,
		}
	}
	return nil
}

func (u userRepository) List(ctx context.Context, filter *entity.UserFilter) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *mongo.Database, collectionName string) entity.IUser {
	return &userRepository{db, collectionName}
}
