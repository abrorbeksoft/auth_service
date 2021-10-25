package storage

import (
	"database/sql"
	"github.com/abrorbeksoft/auth_service/storage/postgres"
	"github.com/abrorbeksoft/auth_service/storage/repos"
)

type StorageI interface {
	UserStorage() repos.UserStorageI
}

type postgresStorage struct {
	postgresRepo repos.UserStorageI
}


func NewStorage(session *sql.DB) StorageI {
	return &postgresStorage{
		postgresRepo: postgres.NewAuth(session),
	}
}

func (p postgresStorage) UserStorage() repos.UserStorageI {
	return p.postgresRepo
}
