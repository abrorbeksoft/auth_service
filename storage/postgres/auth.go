package postgres

import (
	"database/sql"
	auth "github.com/abrorbeksoft/auth_service/genproto/auth_service"
	"github.com/abrorbeksoft/auth_service/storage/repos"
)

type userStorage struct {
	session *sql.DB
}

func NewAuth(session *sql.DB) repos.UserStorageI  {
	return &userStorage{
		session: session,
	}
}

func (a *userStorage) Create(user *auth.RegisterRequest) string {
	panic("implement me")
}

func (a *userStorage) GetById(id string) {
	panic("implement me")
}