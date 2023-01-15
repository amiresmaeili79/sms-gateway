package repository

import "context"

type Repository interface {
	List(ctx context.Context) (interface{}, error)
	Get(id interface{}, ctx context.Context) (interface{}, error)
	Create(entity interface{}, ctx context.Context) (interface{}, error)
}
