package repository

type Repository interface {
	Configure(db interface{})
	List() (interface{}, error)
	Get(id interface{}) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
}
