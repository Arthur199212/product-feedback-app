package auth

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	users "product-feedback/user"
	"product-feedback/utils"
)

type AuthService interface {
	getGithubAccessToken(code string) (string, error)
	getUserDataFromGitHub(code string) (userDataFromGitHub, error)
	getUserEmailFromGitHub(accessToken string) (string, error)
}

type authService struct {
	repo users.UserRepository
}

func NewAuthService(userRepo users.UserRepository) AuthService {
	return &authService{userRepo}
}

const (
	ghLoginOauthAccessTokenURI = "https://github.com/login/oauth/access_token"
	ghLoginOauthAuthorizeURI   = "https://github.com/login/oauth/authorize"
	ghRedirectURI              = "http://localhost:8000/api/auth/github/callback"
	ghUserEmailsURI            = "https://api.github.com/user/emails"
	ghUserEmailScope           = "user:email"
)

type ghAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func (s *authService) getGithubAccessToken(code string) (string, error) {
	bodyMap := map[string]string{
		"client_id":     os.Getenv("GITHUB_CLIENT_ID"),
		"client_secret": os.Getenv("GITHUB_CLIENT_SECRET"),
		"code":          code,
	}
	body := new(bytes.Buffer)
	err := utils.ToJSON(body, bodyMap)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, ghLoginOauthAccessTokenURI, body)
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

	var ghResp ghAccessTokenResponse
	err = utils.FromJSON(res.Body, &ghResp)
	if err != nil {
		return "", err
	}

	return ghResp.AccessToken, nil
}

type userDataFromGitHub struct {
	AvatarUrl string `json:"avatar_url"`
	Email     string `json:"email"`
	Id        string `json:"id"`
	Name      string `json:"name"`
}

func (s *authService) getUserDataFromGitHub(
	code string,
) (userDataFromGitHub, error) {
	accessToken, err := s.getGithubAccessToken(code)
	if err != nil {
		log.Println(err) // todo: use loggrus instance
		return userDataFromGitHub{}, err
	}

	// avatarUrl: userGitHub.avatar_url,
	// githubId: userGitHub.id,
	// name: userGitHub.name

	// email, err := h.getUserEmailFromGitHub(accessToken)
	// if err != nil {
	// 	log.Println(err)
	// 	c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
	// 		"message": "login failed",
	// 	})
	// 	return
	// }

	email, err := s.getUserEmailFromGitHub(accessToken)
	if err != nil {
		log.Println(err) // todo: use loggrus instance
		return userDataFromGitHub{}, err
	}

	mockUserData := userDataFromGitHub{
		Email: email,
	}

	return mockUserData, nil
}

type ghUserEmailResponse struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

func (s *authService) getUserEmailFromGitHub(
	accessToken string,
) (string, error) {
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
	err = utils.FromJSON(res.Body, &emailsData)
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
