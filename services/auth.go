package services

import (
	"context"
	auth "github.com/abrorbeksoft/auth_service/genproto/auth_service"
	"github.com/abrorbeksoft/auth_service/storage"
)

type authService struct {
	storage storage.StorageI
}


func NewAuthService() *authService {
	return &authService{
		//storage:
	}
}


func (a authService) LoginUser(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	var token auth.Token
	token.Id="sdmfldsmf"
	token.CreatedAt="lsdmldmf"
	token.UpdatedAt="sdmfldfm"

	return &auth.LoginResponse{
		Token: &token,
	}, nil
}