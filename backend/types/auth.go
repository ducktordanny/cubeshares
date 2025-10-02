package types

import "github.com/golang-jwt/jwt/v5"

type WCAAuthorizationInfo struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int64  `json:"created_at"`
}

type AuthClaims struct {
	UserId int64  `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
