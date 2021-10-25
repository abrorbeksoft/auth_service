package services

import (
	"context"
	"database/sql"
	auth "github.com/abrorbeksoft/auth_service/genproto/auth_service"
	"github.com/abrorbeksoft/auth_service/storage"
)

type authService struct {
	storage storage.StorageI
}

func NewAuthService(session *sql.DB) *authService {
	return &authService{
		storage: storage.NewStorage(session),
	}
}


func (service authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	id:=service.storage.UserStorage().Create(request)
	return &auth.RegisterResponse{
		Id: id,
	}, nil
}

func (service authService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	panic("implement me")
}

func (service authService) FindOne(ctx context.Context, request *auth.FindOneRequest) (*auth.FindOneResponse, error) {
	user:=service.storage.UserStorage().GetById(request.Id)
	return user,nil
}
