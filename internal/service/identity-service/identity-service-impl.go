package identity_service

import (
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	authorizationHelper "gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/helpers"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/mapper"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity/enum"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type identityService struct {
	identityRepo repository.IdentityRepository
	firebase     *firebase.App
}

func NewIdentityService(
	identityRepo *repository.IdentityRepository,
	firebase *firebase.App,
) IdentityService {
	return &identityService{
		identityRepo: *identityRepo,
		firebase:     firebase,
	}
}

func (i identityService) CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error) {
	return i.identityRepo.CreateUser(ctx, createUserRequest)
}

func (i identityService) VerifyTokenAndGenerateCustomToken(ctx *gin.Context, idToken string) {
	client, err := i.firebase.Auth(ctx)
	if err != nil {
		log.Println("Firebase auth client error:", err)
		return
	}
	_, err = verifyFirebaseToken(ctx, client, idToken)
	if err != nil {
		log.Println("Firebase token verification error:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
		return

	}

	token, _, err := new(jwt.Parser).ParseUnverified(idToken, jwt.MapClaims{})
	if err != nil {
		log.Println("Token failed:", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Claims failed")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		log.Println("Firebase User ID failed")
	}

	userInfo, err := i.identityRepo.GetUserInfoFromFirebaseToken(client, userID)
	if err != nil {
		log.Println("User info error:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: " + err.Error()})
		return
	}

	generateCustomToken(ctx, client, userInfo)
}

func (i identityService) GetUserInfoById(c *gin.Context, id uuid.UUID) {
	userEntity, err := i.identityRepo.GetUserInfoById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	userDTO, err := mapper.MapUserEntityToDto(userEntity)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, userDTO)

}

func (i identityService) VerifyCustomToken(ctx context.Context, firebaseAuth *auth.Client, customToken string, allowedRoles []enum.Role) (jwt.MapClaims, error) {
	claims, err := authorizationHelper.VerifyToken(customToken, allowedRoles)
	if err != nil {
		log.Println("Geçersiz token:", err)
		return nil, err
	}

	return claims, nil
}

func (i identityService) PatchUserInterests(ctx *gin.Context, id uuid.UUID, patchUserInterestsRequest dto.PatchUserInterestsRequest) {
	i.identityRepo.PatchUserInterests(ctx, id, patchUserInterestsRequest)
}

func (i identityService) GetAllInterests(ctx *gin.Context) {
	interests := i.identityRepo.GetAllInterests(ctx)
	if interests == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all interests"})
		return
	}
	interestsDTO := dto.GetAllInterestsResponse{Interests: interests}
	ctx.JSON(200, interestsDTO)
	return
}

func verifyFirebaseToken(ctx context.Context, firebaseAuth *auth.Client, customToken string) (*auth.Token, error) {
	token, err := firebaseAuth.VerifyIDToken(ctx, customToken)
	if err != nil {
		log.Println("Geçersiz token:", err)
		return nil, err
	}
	return token, nil
}

func generateCustomToken(ctx *gin.Context, firebaseAuth *auth.Client, userInfo entity.User) {
	authorizationHelper.GenerateTokenHandler(ctx, userInfo.ID, userInfo.Role)
}
