package repository

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/mapper"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
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
			DisplayName(createUserRequest.Name)
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
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (r *IdentityRepository) VerifyToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := r.firebase.Auth(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	token, err := client.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		log.Println(err)
		return token, err
	}
	return token, nil
}

func (r *IdentityRepository) GetUserInfo(ctx context.Context, idToken string) (string, error) {
	client, err := r.firebase.Auth(ctx)
	if err != nil {
		log.Println(err)
		return "", err
	}
	token, err := client.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return token.UID, nil
}
