package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ducktordanny/cubeshares/backend/configs"
	"github.com/ducktordanny/cubeshares/backend/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("auth/logout", handler.handleLogout)
	router.GET("auth/login", handler.handleLogin)
	router.GET("oauth/callback", handler.handleOAuthCallback)
}

func (handler *Handler) handleLogout(context *gin.Context) {
	context.SetCookie("cubeshares.oauthState", "", -1, "/", "", configs.Envs.Production, true)
	context.SetCookie("cubeshares.session", "", -1, "/", "", configs.Envs.Production, true)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "User logged out."})
}

func (handler *Handler) handleLogin(context *gin.Context) {
	state, err := generateState(32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate state"})
		return
	}

	context.SetCookie("cubeshares.oauthState", state, 300, "/", "", configs.Envs.Production, true)

	authURL, err := url.Parse("https://www.worldcubeassociation.org/oauth/authorize")
	if err != nil {
		log.Fatalf("Invalid base URL: %v", err)
	}

	params := url.Values{}
	params.Add("client_id", configs.Envs.ClientID)
	params.Add("redirect_uri", configs.Envs.RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", "public email")
	params.Add("state", state)

	authURL.RawQuery = params.Encode()
	finalURL := authURL.String()

	context.Redirect(http.StatusFound, finalURL)
}

func (handler *Handler) handleOAuthCallback(context *gin.Context) {
	code := context.Query("code")
	state := context.Query("state")
	expectedState, err := context.Cookie("cubeshares.oauthState")
	if err != nil || state != expectedState {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing state"})
		return
	}

	authInfo, err := getWCAAuthInfo(code)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to exchange code: " + err.Error(),
		})
		return
	}

	wcaUser, err := handler.store.GetWCAUser(authInfo.AccessToken)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get WCA user: " + err.Error(),
		})
		return
	}

	user, err := handler.store.RegisterOrUpdateUser(wcaUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register/update user: " + err.Error(),
		})
		return
	}
	initUserAuthSession(context, user)
	context.Redirect(http.StatusSeeOther, configs.Envs.ClientAppURL)
}

func generateState(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func getWCAAuthInfo(code string) (types.WCAAuthorizationInfo, error) {
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {configs.Envs.ClientID},
		"client_secret": {configs.Envs.ClientSecret},
		"code":          {code},
		"redirect_uri":  {configs.Envs.RedirectURI},
	}

	tokenURL := "https://www.worldcubeassociation.org/oauth/token"
	res, err := http.PostForm(tokenURL, form)
	if err != nil {
		return types.WCAAuthorizationInfo{}, fmt.Errorf("Request failed: %w", err)
	}
	defer res.Body.Close()

	var auth types.WCAAuthorizationInfo
	if err := json.NewDecoder(res.Body).Decode(&auth); err != nil {
		return types.WCAAuthorizationInfo{}, fmt.Errorf("Decode failed: %w", err)
	}

	return auth, nil
}

func initUserAuthSession(context *gin.Context, user types.User) {
	secret := configs.Envs.JWTSecret

	sessionAge := time.Hour
	expiresAt := time.Now().Add(sessionAge).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.Id,
		"email": user.Email,
		"exp":   expiresAt,
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign JWT"})
		return
	}

	if configs.Envs.Production != true {
		fmt.Printf("\nJWT: %s\n\n", tokenString)
	}
	context.SetCookie("cubeshares.session", tokenString, int(sessionAge.Seconds()), "/", "", configs.Envs.Production, true)
}
