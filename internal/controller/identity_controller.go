package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
	identityService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/identity-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type IdentityController interface {
	CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error)
	VerifyToken(ctx *gin.Context, verifyTokenRequest dto.VerifyTokenRequest)
	GetUserInfoById(ctx *gin.Context, id uuid.UUID)
	PatchUserInterests(ctx *gin.Context, id uuid.UUID, patchUserInterestsRequest dto.PatchUserInterestsRequest)
	GetAllInterests(ctx *gin.Context)
	SearchUsersByKeyword(ctx *gin.Context, keyword string, limit int, offset int) ([]dto.PublicUserProfileDTO, error)
	GetUserPublicProfileByID(ctx *gin.Context, id uuid.UUID) (*dto.PublicUserProfileDTO, error)
}

type identityController struct {
	identityService identityService.IdentityService
}

func (c identityController) SearchUsersByKeyword(ctx *gin.Context, keyword string, limit int, offset int) ([]dto.PublicUserProfileDTO, error) {
	//TODO implement me
	panic("implement me")
}

// GetUserPublicProfileByID is a controller method
// that gets a user public profile by id
// @Summary Get user public profile by id
// @Description Get user public profile by id
// @ID get-user-public-profile-by-id
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} dto.PublicUserProfileDTO
// @Router /identity/profile/{id} [get]
// @Tags identity
// @Security AccessToken[admin, user]
func (c identityController) GetUserPublicProfileByID(ctx *gin.Context, id uuid.UUID) (*dto.PublicUserProfileDTO, error) {
	publicUserProfile, err := c.identityService.GetUserPublicProfileByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return publicUserProfile, nil
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
// @Accept  multipart/form-data
// @Produce  json
// @Param email formData string true "User's email"
// @Param password formData string true "User's password"
// @Param username formData string true "User's username"
// @Param image formData file false "User's profile image"
// @Success 200 {object} entity.User "Return user successfully"
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

// PatchUserInterests is a controller method
// that patches user interests
// @Summary
// @Description  The operation to perform on the user's interests (either "add" or "remove") and the interest's ID.
// @Tags identity
// @Accept  json
// @Produce  json
// @Param patchUserInterestsRequest body dto.PatchUserInterestsRequest true "PatchUserInterestsRequest"
// @Success 200
// @Router /identity/userinfo/interests [patch]
// @Security AccessToken[admin, user]
func (c identityController) PatchUserInterests(ctx *gin.Context, id uuid.UUID, patchUserInterestsRequest dto.PatchUserInterestsRequest) {
	c.identityService.PatchUserInterests(ctx, id, patchUserInterestsRequest)
}

// GetAllInterests is a controller method
// that gets all interests
// @Summary Get all interests
// @Tags identity
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GetAllInterestsResponse
// @Router /identity/userinfo/interests [get]
// @Security AccessToken[admin, user]
func (c identityController) GetAllInterests(ctx *gin.Context) {
	c.identityService.GetAllInterests(ctx)
}
