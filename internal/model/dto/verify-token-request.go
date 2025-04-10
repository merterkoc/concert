package dto

type VerifyTokenRequest struct {
	IdToken string `json:"idToken" binding:"required"`
}
