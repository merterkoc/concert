package identity_service

import (
	"firebase.google.com/go/auth"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity/enum"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type IdentityService interface {
	CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error)
	VerifyTokenAndGenerateCustomToken(ctx *gin.Context, idToken string)
	GetUserInfoById(c *gin.Context, id uuid.UUID)
	VerifyCustomToken(ctx context.Context, firebaseAuth *auth.Client, customToken string, allowedRoles []enum.Role) (jwt.MapClaims, error)
	PatchUserInterests(ctx *gin.Context, id uuid.UUID, patchUserInterestsRequest dto.PatchUserInterestsRequest)
	GetAllInterests(ctx *gin.Context)
	GetUserImageByID(id uuid.UUID) *dto.ParticipantsAvatar
	GetUserPublicProfileByID(ctx *gin.Context, id uuid.UUID) (*dto.PublicUserProfileDTO, error)
	SearchUsersByKeyword(ctx *gin.Context, keyword string, limit int, offset int) ([]dto.PublicUserProfileDTO, error)
}
