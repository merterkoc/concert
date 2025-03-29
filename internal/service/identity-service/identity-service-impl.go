package identity_service

import (
	"firebase.google.com/go/auth"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	"golang.org/x/net/context"
)

type identityService struct {
	identityRepo repository.IdentityRepository
}

func (i identityService) CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error) {
	return i.identityRepo.CreateUser(ctx, createUserRequest)
}

func (i identityService) VerifyToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return i.identityRepo.VerifyToken(ctx, idToken)
}

func (i identityService) GetUserInfo(ctx context.Context, idToken string) (string, error) {
	return i.identityRepo.GetUserInfo(ctx, idToken)
}

func NewIdentityService(
	identityRepo *repository.IdentityRepository,
) IdentityService {
	return &identityService{
		identityRepo: *identityRepo,
	}
}
