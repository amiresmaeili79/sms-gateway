package repository

import "github.com/amir79esmaeili/sms-gateway/internal/cfg"

type Repository interface {
	Configure(config *cfg.Config)
	List() (interface{}, error)
	Get(id interface{}) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
}
