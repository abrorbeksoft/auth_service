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


func (a authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	panic("implement me")
}

func (a authService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	panic("implement me")
}