package auth

import (
	"net/http"
	"net/url"
	"os"
	"product-feedback/validation"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type authHandler struct {
	l       *logrus.Logger
	v       *validation.Validation
	service AuthService
}

func NewAuthHandler(
	l *logrus.Logger,
	v *validation.Validation,
	service AuthService,
) AuthHandler {
	return &authHandler{
		l:       l,
		v:       v,
		service: service,
	}
}

// swagger:route GET /api/auth/github/callback auth loginWithGitHub
// Redirects to GitHub authentication
// responses:
//	302: foundResponse
//	403: errorResponse

func (h *authHandler) loginWithGitHub(c *gin.Context) {
	code := c.Query("code")

	userId, err := h.service.loginWithGitHub(code)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "login failed",
		})
		return
	}

	accessToken, err := h.service.generateAccessToken((strconv.Itoa(userId)))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "login failed",
		})
		return
	}

	refreshToken, err := h.service.generateRefreshToken(strconv.Itoa(userId))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "login failed",
		})
		return
	}

	// Option with cookies
	// c.SetCookie(
	// 	"refresh-token",
	// 	refreshToken,
	// 	int(refreshTokenTTL),
	// 	"/api/auth/refresh-token",
	// 	"localhost",
	// 	true,
	// 	true,
	// )

	// Option with tokens in callbackUrl
	q := url.Values{}
	q.Set("access_token", accessToken)
	q.Set("refresh_token", refreshToken)
	loginCallbackUrlWithTokens := url.URL{
		Path:     os.Getenv("LOGIN_CALLBACK_URL"),
		RawQuery: q.Encode(),
	}

	c.Redirect(http.StatusFound, loginCallbackUrlWithTokens.RequestURI())
}

// swagger:route GET /api/auth/github auth loginWithGitHub
// Redirects to GitHub authentication
// responses:
//	302: foundResponse

func (h *authHandler) redirectToGitHubLoginURL(c *gin.Context) {
	q := url.Values{}
	q.Set("client_id", os.Getenv("GITHUB_CLIENT_ID"))
	q.Set("redirect_uri", os.Getenv("FRONTEND_URL")+ghRedirectURI)
	q.Set("scope", ghUserEmailScope)
	location := url.URL{Path: ghLoginOauthAuthorizeURI, RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())
}

type refreshAccessTokenInput struct {
	// Refresh token in format of JWT is used to get
	// access token in exchange
	//
	// required: true
	RefreshToken string `json:"refreshToken" validate:"required,jwt"`
}

// swagger:route POST /api/auth/refresh-token auth refreshAccessToken
// Redirects to GitHub authentication
// responses:
//	200: refreshAccessTokenResponse
//	403: errorResponse

func (h *authHandler) refreshAccessToken(c *gin.Context) {
	// Option with refresh token in cookie
	// refreshToken, err := c.Cookie("refresh-token")
	// if err != nil {
	// 	h.l.Error(err)
	// 	c.AbortWithStatus(http.StatusForbidden)
	// 	return
	// }

	// Option with refresh token in request.body
	var input refreshAccessTokenInput
	if err := c.BindJSON(&input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "Forbidden",
		})
		return
	}

	if err := h.v.ValidateStruct(input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "Forbidden",
		})
		return
	}

	userId, err := h.service.verifyRefreshToken(input.RefreshToken)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "Forbidden",
		})
		return
	}

	token, err := h.service.generateAccessToken(strconv.Itoa(userId))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "Forbidden",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": token,
	})
}
