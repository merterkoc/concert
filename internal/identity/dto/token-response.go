package dto

import "github.com/google/uuid"

type Claims struct {
	Role         string                 `json:"role"`
	CustomClaims map[string]interface{} `json:"custom_claims,omitempty"`
}

type TokenResponse struct {
	Token     string    `json:"token"`
	Type      string    `json:"type"`
	UID       uuid.UUID `json:"uid"`
	Claims    Claims    `json:"claims"`
	ExpiresIn int       `json:"expires_in"`
	IssuedAt  int       `json:"issued_at"`
}
