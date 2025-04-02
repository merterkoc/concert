package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
	identityService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/identity-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type IdentityController interface {
	CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error)
	VerifyToken(ctx *gin.Context, verifyTokenRequest dto.VerifyTokenRequest)
	GetUserInfoById(ctx *gin.Context, id uuid.UUID)
}

type identityController struct {
	identityService identityService.IdentityService
}

func NewIdentityController(identityService identityService.IdentityService) IdentityController {
	return identityController{
		identityService: identityService,
	}
}

// CreateUser is a controller method
// that creates a user
// @Summary Create user
// @Description Create user
// @Tags identity
// @Accept  json
// @Produce  json
// @Param createUserRequest body dto.CreateUserRequest true "CreateUserRequest"
// @Router /identity/create [post]
func (c identityController) CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error) {
	return c.identityService.CreateUser(ctx, createUserRequest)
}

// VerifyToken is a controller method
// that verifies the token
// @Summary Verify token
// @Description Verify token
// @Tags identity
// @Accept  json
// @Produce  json
// @Param verifyTokenRequest body dto.VerifyTokenRequest true "VerifyTokenRequest"
// @Router /identity/verify [post]
func (c identityController) VerifyToken(ctx *gin.Context, verifyTokenRequest dto.VerifyTokenRequest) {
	c.identityService.VerifyTokenAndGenerateCustomToken(ctx, verifyTokenRequest.IdToken)
}

// GetUserInfoById is a controller method
// that gets the user info
// @Summary Get user info
// @Description Get user info
// @Tags identity
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.UserDto "Return user dto successfully"
// @Router /identity/userinfo [get]
// @Security AccessToken[admin, user]
func (c identityController) GetUserInfoById(ctx *gin.Context, id uuid.UUID) {
	c.identityService.GetUserInfoById(ctx, id)
}
