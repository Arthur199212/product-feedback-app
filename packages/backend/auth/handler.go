package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

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
	userData, err := h.service.getUserDataFromGitHub(code)
	if err != nil {
		// fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"message": "login failed",
		})
		return
	}

	// todo: login user
	fmt.Println(userData)

	c.Header("refresh-token", userData.Email)
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

func (h *authHandler) signIn(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "signIn not implemented",
	})
}

func (h *authHandler) signUp(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "signUp not implemented",
	})
}
