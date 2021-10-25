package repos

import "github.com/abrorbeksoft/auth_service/genproto/auth_service"

type UserStorageI interface {
	Create(user *auth_service.RegisterRequest) string
	GetById(id string)
}
