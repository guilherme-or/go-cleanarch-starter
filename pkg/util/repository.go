package util

type CRUDRepository[T any] interface {
	FindAll() ([]T, error)
	FindById(id int) (T, error)
	Create(entity T) (T, error)
	Update(id int, entity T) error
	Delete(id int) error
}
