package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"jobsearch-demos/kanban/internal/board/entity"
)

type boardRepository struct {
	db             mongo.Database
	collectionName string
}

func (b boardRepository) Create(ctx context.Context, board *entity.Board) (*entity.Board, error) {
	//TODO implement me
	panic("implement me")
}

func (b boardRepository) Get(ctx context.Context, id string) (*entity.Board, error) {
	//TODO implement me
	panic("implement me")
}

func (b boardRepository) List(ctx context.Context, filter interface{}) ([]*entity.Board, error) {
	//TODO implement me
	panic("implement me")
}

func (b boardRepository) Update(ctx context.Context, id string, board *entity.Board) (*entity.Board, error) {
	//TODO implement me
	panic("implement me")
}

func (b boardRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewBoardRepository(db mongo.Database, collectionName string) entity.IBoard {
	return &boardRepository{db: db, collectionName: collectionName}
}
