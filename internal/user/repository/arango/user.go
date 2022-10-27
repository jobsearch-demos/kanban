package arango

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver"
	"jobsearch-demos/kanban/internal/user/entity"
	errors "jobsearch-demos/kanban/pkg/error"
)

type userRepository struct {
	db             driver.Database
	collectionName string
}

func (u userRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	user := entity.User{}
	query := fmt.Sprintf("FOR u IN users FILTER u.id == %d RETURN u", id)
	cursor, err := u.db.Query(ctx, query, nil)

	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error fetching user: %s\n", err.Error()),
			InfoContext: nil,
		}
	}
	defer func() {
		if err := cursor.Close(); err != nil {
			// TODO: handle logging the error instead of panic
			panic(err)
		}
	}()
	_, err = cursor.ReadDocument(ctx, &user)

	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error decoding the fetched results: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	return &user, nil
}

func (u userRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	query := fmt.Sprintf("INSERT {id: %d, name: '%s', email: '%s', password: '%s'} INTO users RETURN NEW",
		user.ID, user.Username, user.Email, user.Password)

	cursor, err := u.db.Query(ctx, query, nil)

	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error fetching user: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	defer func() {
		if err := cursor.Close(); err != nil {
			// TODO: handle logging the error instead of panic
			panic(err)
		}
	}()

	_, err = cursor.ReadDocument(ctx, &user)

	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error decoding the fetched results: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	return &user, nil
}

func (u userRepository) Update(ctx context.Context, user entity.User) (*entity.User, error) {
	query := fmt.Sprintf("FOR u IN users FILTER u.id == %d UPDATE u WITH {name: '%s', email: '%s', password: '%s'} IN users RETURN NEW",
		user.ID, user.Username, user.Email, user.Password)

	cursor, err := u.db.Query(ctx, query, nil)

	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error fetching user: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	defer func() {
		if err := cursor.Close(); err != nil {
			// TODO: handle logging the error instead of panic
			panic(err)
		}
	}()

	_, err = cursor.ReadDocument(ctx, &user)

	if err != nil {
		return nil, errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error decoding the fetched results: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	return &user, nil
}

func (u userRepository) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf("FOR u IN users FILTER u.id == %d REMOVE u IN users", id)

	_, err := u.db.Query(ctx, query, nil)

	if err != nil {
		return errors.CustomError{
			Code:        500,
			Message:     fmt.Sprintf("error deleting user: %s\n", err.Error()),
			InfoContext: nil,
		}
	}

	return nil
}

func (u userRepository) List(ctx context.Context, user *entity.UserFilter) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db driver.Database, collectionName string) entity.IUser {
	return &userRepository{db, collectionName}
}
