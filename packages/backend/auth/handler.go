package auth

import (
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type authHandler struct {
	l       *logrus.Logger
	service AuthService
}

func NewAuthHandler(l *logrus.Logger, service AuthService) AuthHandler {
	return &authHandler{
		l:       l,
		service: service,
	}
}

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
	// 	refreshTokenCookieName,
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
		Path:     loginCallbackUrl,
		RawQuery: q.Encode(),
	}

	c.Redirect(http.StatusFound, loginCallbackUrlWithTokens.RequestURI())
}

func (h *authHandler) redirectToGitHubLoginURL(c *gin.Context) {
	q := url.Values{}
	q.Set("client_id", os.Getenv("GITHUB_CLIENT_ID"))
	q.Set("redirect_uri", ghRedirectURI)
	q.Set("scope", ghUserEmailScope)
	location := url.URL{Path: ghLoginOauthAuthorizeURI, RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())
}

func (h *authHandler) refreshAccessToken(c *gin.Context) {
	refreshToken, err := c.Cookie(refreshTokenCookieName)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	userId, err := h.service.verifyRefreshToken(refreshToken)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	token, err := h.service.generateAccessToken(strconv.Itoa(userId))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": token,
	})
}
