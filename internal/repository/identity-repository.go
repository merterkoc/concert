package repository

import (
	"cloud.google.com/go/storage"

	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity/enum"
	authorizationHelper "gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/helpers"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/mapper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type IdentityRepository struct {
	db            *gorm.DB
	firebase      *firebase.App
	storageClient *storage.Client
}

func NewIdentityRepository(db *gorm.DB, firebase *firebase.App, storageClient *storage.Client) *IdentityRepository {
	return &IdentityRepository{db: db, firebase: firebase, storageClient: storageClient}
}

func (r *IdentityRepository) CreateUser(ctx context.Context, createUserRequest dto.CreateUserRequest) (*entity.User, error) {
	var imageURL string
	if createUserRequest.Image != nil {
		url, err := r.uploadImageFirebaseStorage(ctx, createUserRequest.Image)
		if err != nil {
			return nil, fmt.Errorf("image upload failed: %w", err)
		}
		imageURL = url
	}

	client, err := r.firebase.Auth(ctx)
	if err != nil {
		log.Printf("failed to get Firebase Auth client: %v", err)
		return nil, err
	}

	userEntity, err := mapper.MapUserRequestToEntityWithImage(createUserRequest, &imageURL)
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
		Email:       userEntity.Email,
		Username:    userEntity.Username,
		FirebaseUID: token.UID,
		UserImage:   userEntity.UserImage,
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

func (r *IdentityRepository) GetUserInfoById(id uuid.UUID) (entity.User, error) {
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

func (r *IdentityRepository) uploadImageFirebaseStorage(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {

	// Dosyayı aç
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("File not opened: %v", err)
	}
	defer file.Close()

	// Firebase Storage bucket'ını belirt
	bucketName := "gigbuddy-dev.firebasestorage.app"
	bucket := r.storageClient.Bucket(bucketName)

	// Dosya adını belirle (örneğin, benzersiz bir ad oluşturabilirsiniz)
	objectName := fmt.Sprintf("uploads/user_images/%d-%s", time.Now().UnixNano(), fileHeader.Filename)
	object := bucket.Object(objectName)

	// Dosyayı yüklemek için bir Writer oluştur
	writer := object.NewWriter(ctx)
	writer.ContentType = fileHeader.Header.Get("Content-Type")

	// Dosya içeriğini Firebase Storage'a yaz
	if _, err := io.Copy(writer, file); err != nil {
		return "", fmt.Errorf("Dosya Firebase Storage'a yazılamadı: %v", err)
	}

	// Writer'ı kapat
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("Writer kapatılamadı: %v", err)
	}

	// Dosyanın herkese açık olarak erişilebilir olmasını sağla
	if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", fmt.Errorf("Dosya herkese açık yapılamadı: %v", err)
	}

	// Dosyanın herkese açık URL'sini oluştur
	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
	return publicURL, nil
}
