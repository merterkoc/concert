package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
	identityService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/identity-service"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type IdentityController interface {
	CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error)
	VerifyToken(ctx *gin.Context, verifyTokenRequest dto.VerifyTokenRequest)
	GetUserInfo(idToken string) (entity.User, error)
}

type identityController struct {
	identityService identityService.IdentityService
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

// GetUserInfo is a controller method
// that gets the user info
// @Summary Get user info
// @Description Get user info
// @Tags identity
// @Accept  json
// @Produce  json
// @Param idToken path string true "IdToken"
// @Success 200 {object} string
// @Router /identity/userinfo [post]
func (c identityController) GetUserInfo(id string) (entity.User, error) {
	return c.identityService.GetUserInfo(id)
}

func NewIdentityController(identityService identityService.IdentityService) IdentityController {
	return identityController{
		identityService: identityService,
	}
}
