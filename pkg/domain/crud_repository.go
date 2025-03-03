package domain

type CRUDRepository[T any] interface {
	FindAll() ([]T, error)
	FindById(id int) (T, error)
	Create(entity T) (T, error)
	Delete(id int) error
}
