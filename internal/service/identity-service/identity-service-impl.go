package identity_service

import (
	"firebase.google.com/go/auth"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity/enum"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/net/context"
)

type identityService struct {
	identityRepo repository.IdentityRepository
}

func (i identityService) CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error) {
	return i.identityRepo.CreateUser(ctx, createUserRequest)
}

func (i identityService) VerifyTokenAndGenerateCustomToken(ctx *gin.Context, idToken string) {
	i.identityRepo.VerifyAndGenerateToken(ctx, idToken)
}

func (i identityService) GetUserInfo(id string) (entity.User, error) {
	return i.identityRepo.GetUserInfo(id)
}

func (i identityService) VerifyCustomToken(ctx context.Context, firebaseAuth *auth.Client, customToken string, allowedRoles []enum.Role) (jwt.MapClaims, error) {
	return i.identityRepo.VerifyCustomToken(ctx, firebaseAuth, customToken, allowedRoles)
}

func NewIdentityService(
	identityRepo *repository.IdentityRepository,
) IdentityService {
	return &identityService{
		identityRepo: *identityRepo,
	}
}
