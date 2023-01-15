package repository

import (
	"reflect"
)

type Registry struct {
	registry map[string]Repository
}

func NewRepositoryRegistry(repository ...Repository) *Registry {
	r := &Registry{
		registry: map[string]Repository{},
	}

	r.registerRepositories(repository)
	return r
}

func (r *Registry) registerRepositories(repositories []Repository) {
	for _, repository := range repositories {
		repositoryName := reflect.TypeOf(repository).Elem().Name()
		r.registry[repositoryName] = repository
	}
}
