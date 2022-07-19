package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// todo: use utils.FromJSON ...

func (h *AuthHandler) AddRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.GET("/github", h.redirectToGitHubLoginURL)
		auth.GET("/github/callback", h.loginWithGitHub)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}
}

const (
	ghRedirectUri              = "http://localhost:8000/api/auth/github/callback"
	ghLoginOauthAuthorizeUri   = "https://github.com/login/oauth/authorize"
	ghLoginOauthAccessTokenUri = "https://github.com/login/oauth/access_token"
)

func (h *AuthHandler) redirectToGitHubLoginURL(c *gin.Context) {
	// baseUrl, err := url.Parse(c.Request.Host + c.Request.URL.Port())
	// if err != nil {
	// 	log.Print(err)
	// }
	// redirectUrl := baseUrl.ResolveReference(&url.URL{Path: ghRedirectUri})

	q := url.Values{}
	q.Set("client_id", os.Getenv("GITHUB_CLIENT_ID"))
	q.Set("redirect_uri", ghRedirectUri)
	q.Set("scope", "user:email")
	location := url.URL{Path: ghLoginOauthAuthorizeUri, RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())
}

func (h *AuthHandler) loginWithGitHub(c *gin.Context) {
	accessToken, err := h.getGithubAccessToken(c.Query("code"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not login",
		})
		return
	}

	data, err := h.getUserEmailFromGitHub(accessToken)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not login",
		})
		return
	}

	fmt.Println(data)
	// todo
	c.Header("refresh-token", data)
	c.Redirect(http.StatusFound, "http://localhost:8000/")
}

type githubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func (h *AuthHandler) getGithubAccessToken(code string) (string, error) {
	bodyMap := map[string]string{
		"client_id":     os.Getenv("GITHUB_CLIENT_ID"),
		"client_secret": os.Getenv("GITHUB_CLIENT_SECRET"),
		"code":          code,
	}
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(bodyMap)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, ghLoginOauthAccessTokenUri, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var ghResp githubAccessTokenResponse
	err = json.NewDecoder(res.Body).Decode(&ghResp)
	if err != nil {
		return "", err
	}

	return ghResp.AccessToken, nil
}

type ghUserEmailResponse struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

func (h *AuthHandler) getUserEmailFromGitHub(accessToken string) (string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://api.github.com/user/emails",
		nil,
	)
	if err != nil {
		return "", nil
	}

	authTokenHeader := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authTokenHeader)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil
	}

	// b, err := ioutil.ReadAll(res.Body)
	// fmt.Println(string(b))

	var emailsData []ghUserEmailResponse
	err = json.NewDecoder(res.Body).Decode(&emailsData)
	if err != nil {
		return "", err
	}

	for _, data := range emailsData {
		if data.Primary && data.Verified {
			return data.Email, nil
		}
	}

	return "", errors.New("could not retrieve an email")
}
