package repository

import (
	"errors"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity/enum"
	authorizationHelper "gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/helpers"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/mapper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type IdentityRepository struct {
	db       *gorm.DB
	firebase *firebase.App
}

func NewIdentityRepository(db *gorm.DB, firebase *firebase.App) *IdentityRepository {
	return &IdentityRepository{db: db, firebase: firebase}
}

func (r *IdentityRepository) CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error) {
	client, err := r.firebase.Auth(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	toEntity, err := mapper.MapUserRequestToEntity(createUserRequest)
	if err != nil {
		return nil, err
	}

	params :=
		(&auth.UserToCreate{}).
			Email(createUserRequest.Email).
			Password(createUserRequest.Password).
			DisplayName(createUserRequest.Username)
	token, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user := &entity.User{
		Email:       toEntity.Email,
		Username:    toEntity.Username,
		FirebaseUID: token.UID,
	}

	err = r.db.Create(user).Error
	if err != nil {
		log.Println("Database error:", err)

		delErr := client.DeleteUser(ctx, token.UID)
		if delErr != nil {
			//TODO: Add metric
			log.Println("Failed to delete Firebase user:", delErr)
		}

		return nil, err
	}

	return user, nil
}

func (r *IdentityRepository) VerifyAndGenerateToken(ctx *gin.Context, firebaseToken string) {
	client, err := r.firebase.Auth(ctx)
	if err != nil {
		log.Println("Firebase auth client error:", err)
		return
	}
	_, err = r.VerifyFirebaseToken(ctx, client, firebaseToken)
	if err != nil {
		log.Println("Firebase token verification error:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
		return

	}

	//parser
	token, _, err := new(jwt.Parser).ParseUnverified(firebaseToken, jwt.MapClaims{})
	if err != nil {
		log.Println("Token ayrıştırılamadı:", err)
	}

	// Token claim'lerini al
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Claims ayrıştırılamadı")
	}

	// `sub` claim'ini (Firebase User ID) al
	userID, ok := claims["sub"].(string)
	if !ok {
		log.Println("Firebase User ID ayrıştırılamadı")
	}
	/// Token'ı imzala ve string'e dönüştür

	userInfo, err := r.GetUserInfoFromFirebaseToken(userID)
	if err != nil {
		log.Println("User info error:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: " + err.Error()})
		return
	}

	GenerateCustomToken(ctx, client, userInfo)

}

func (r *IdentityRepository) GetUserInfo(id string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not found for ID:", id)
			return entity.User{}, fmt.Errorf("user not found")
		}
		log.Println("Database error:", err)
		return entity.User{}, err
	}

	return user, nil
}

func (r *IdentityRepository) VerifyFirebaseToken(ctx context.Context, firebaseAuth *auth.Client, customToken string) (*auth.Token, error) {
	token, err := firebaseAuth.VerifyIDToken(ctx, customToken)
	if err != nil {
		log.Println("Geçersiz token:", err)
		return nil, err
	}
	return token, nil
}

func (r *IdentityRepository) VerifyCustomToken(ctx context.Context, firebaseAuth *auth.Client, customToken string, allowedRoles []enum.Role) (jwt.MapClaims, error) {
	claims, err := authorizationHelper.VerifyToken(customToken, allowedRoles)
	if err != nil {
		log.Println("Geçersiz token:", err)
		return nil, err
	}

	// Token geçerliyse token bilgilerini döndür
	return claims, nil
}

func GenerateCustomToken(ctx *gin.Context, firebaseAuth *auth.Client, userInfo entity.User) {
	//claims := map[string]interface{}{
	//	"role": "user",
	//}
	authorizationHelper.GenerateTokenHandler(ctx, userInfo.ID, userInfo.Role)
	//if err != nil {
	//	log.Println("Failed to create custom token:", err)
	//	return nil, err
	//}

	//return &dto.TokenResponse{
	//	Token:     customToken,
	//	Type:      "Bearer",
	//	UID:       uid,
	//	Claims:    dto.Claims{CustomClaims: claims},
	//	ExpiresIn: 3600, // 1 saat (3600 saniye)
	//	IssuedAt:  int(time.Now().Unix()),
	//}, nil
}

func (r *IdentityRepository) GetUserInfoFromFirebaseToken(firebaseUID string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("firebase_uid = ?", firebaseUID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not found for Firebase UID:", firebaseUID)
			return entity.User{}, fmt.Errorf("user not found")
		}
		log.Println("Database error:", err)
		return entity.User{}, err
	}

	return user, nil
}
