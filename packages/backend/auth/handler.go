package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type authHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) AuthHandler {
	return &authHandler{service}
}

func (h *authHandler) loginWithGitHub(c *gin.Context) {
	code := c.Query("code")

	userId, err := h.service.loginWithGitHub(code)
	if err != nil {
		// fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "login failed",
		})
		return
	}

	token, err := h.service.generateRefreshToken(strconv.Itoa(userId))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "login failed",
		})
		return
	}

	c.SetCookie(
		refreshTokenCookieName,
		token,
		int(refreshTokenTTL),
		"/api/auth/refresh-token",
		"localhost",
		true,
		true,
	)
	c.Redirect(http.StatusFound, "http://localhost:8000/")
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
		fmt.Println(err) // todo: use loggrus
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	userId, err := h.service.verifyRefreshToken(refreshToken)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	token, err := h.service.generateAccessToken(strconv.Itoa(userId))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": token,
	})
}
