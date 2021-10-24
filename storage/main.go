package storage

type StorageI interface {
	UserStorage() UserStorage
}

type UserStorage interface {
	Create(id string) string
}
type Just struct {

}

func NewStorage() StorageI {
	return &Just{

	}
}


func (j *Just) UserStorage() UserStorage {
	return NewJust1()
}

type Just1 struct {

}

func NewJust1() UserStorage {
	return &Just1{

	}
}

func (j *Just1) Create(id string) string {
	return "Hello"
}
