package identity_service

import (
	"firebase.google.com/go/auth"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity/enum"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/net/context"
)

type IdentityService interface {
	CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error)
	VerifyTokenAndGenerateCustomToken(ctx *gin.Context, idToken string)
	GetUserInfo(id string) (entity.User, error)
	VerifyCustomToken(ctx context.Context, firebaseAuth *auth.Client, customToken string, allowedRoles []enum.Role) (jwt.MapClaims, error)
}
