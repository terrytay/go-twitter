package db

import "context"

type DBClient interface {
	FindOne(ctx context.Context, collection string, id string) (interface{}, error)
	FindAll(ctx context.Context, collection string) ([]interface{}, error)
	Update(ctx context.Context, collection string, data Data) error
	Delete(ctx context.Context, collection string, id string) error
	Create(ctx context.Context, collection string, data Data) error
	Close() error
}
