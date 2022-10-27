package arango

import (
	"context"
	"github.com/arangodb/go-driver"
	"jobsearch-demos/kanban/internal/board/entity"
)

type columnRepository struct {
	db             driver.Database
	collectionName string
}

func (c columnRepository) Create(ctx context.Context, column *entity.Column) (*entity.Column, error) {
	//TODO implement me
	panic("implement me")
}

func (c columnRepository) Get(ctx context.Context, id string) (*entity.Column, error) {
	//TODO implement me
	panic("implement me")
}

func (c columnRepository) List(ctx context.Context, filter interface{}) ([]*entity.Column, error) {
	//TODO implement me
	panic("implement me")
}

func (c columnRepository) Update(ctx context.Context, id string, column *entity.Column) (*entity.Column, error) {
	//TODO implement me
	panic("implement me")
}

func (c columnRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewColumnRepository(db driver.Database, collectionName string) entity.IColumn {
	return &columnRepository{db: db, collectionName: collectionName}
}
