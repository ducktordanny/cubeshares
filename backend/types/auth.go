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
	Sub   int64  `json:"sub"`
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
	jwt.RegisteredClaims
}
