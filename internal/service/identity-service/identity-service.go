package identity_service

import (
	"firebase.google.com/go/auth"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
	"golang.org/x/net/context"
)

type IdentityService interface {
	CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error)
	VerifyToken(ctx context.Context, idToken string) (*auth.Token, error)
	GetUserInfo(ctx context.Context, idToken string) (string, error)
}
