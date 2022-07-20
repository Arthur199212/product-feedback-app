package auth

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"product-feedback/user"
	"product-feedback/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	generateAccessToken(userId string) (string, error)
	generateRefreshToken(userId string) (string, error)
	getGithubAccessToken(code string) (string, error)
	getUserDataFromGitHub(code string) (userDataFromGitHub, error)
	getUserEmailFromGitHub(accessToken string) <-chan getUserEmailFromGitHubResponse
	getUserFromGitHub(accessToken string) <-chan getUserFromGitHubResponse
	loginWithGitHub(code string) (int, error)
	verifyRefreshToken(tokenStr string) (int, error)
}

type authService struct {
	userService user.UserService
}

func NewAuthService(userService user.UserService) AuthService {
	return &authService{userService}
}

const (
	ghLoginOauthAccessTokenURI = "https://github.com/login/oauth/access_token"
	ghLoginOauthAuthorizeURI   = "https://github.com/login/oauth/authorize"
	ghRedirectURI              = "http://localhost:8000/api/auth/github/callback"
	ghUserEmailsURI            = "https://api.github.com/user/emails"
	ghUserEmailScope           = "user:email"

	accessTokenTTL         = 30 * time.Minute
	refreshTokenCookieName = "refresh-token"
	refreshTokenRoute      = "/api/auth/refresh-token"
	refreshTokenTTL        = 72 * time.Hour
)

func (s *authService) generateAccessToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   userId,
	},
	)
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}

func (s *authService) generateRefreshToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   userId,
	},
	)
	return token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
}

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
	Login     string `json:"login"`
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

	userReq := s.getUserFromGitHub(accessToken)
	userEmailReq := s.getUserEmailFromGitHub(accessToken)

	userResp, userEmailResp := <-userReq, <-userEmailReq
	if userResp.err != nil {
		log.Println(userResp.err)
		return userDataFromGitHub{}, userResp.err
	}
	if userEmailResp.err != nil {
		log.Println(userEmailResp.err)
		return userDataFromGitHub{}, userEmailResp.err
	}

	userResp.user.Email = userEmailResp.email

	return userResp.user, nil
}

type getUserFromGitHubResponse struct {
	user userDataFromGitHub
	err  error
}

func (s *authService) getUserFromGitHub(
	accessToken string,
) <-chan getUserFromGitHubResponse {
	ch := make(chan getUserFromGitHubResponse, 1)

	go func() {
		req, err := http.NewRequest(
			http.MethodGet,
			"https://api.github.com/user",
			nil,
		)
		if err != nil {
			ch <- getUserFromGitHubResponse{
				user: userDataFromGitHub{},
				err:  err,
			}
		}

		authTokenHeader := fmt.Sprintf("token %s", accessToken)
		req.Header.Set("Authorization", authTokenHeader)
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			ch <- getUserFromGitHubResponse{
				user: userDataFromGitHub{},
				err:  err,
			}
		}

		// b, err := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(b))

		var user userDataFromGitHub
		err = utils.FromJSON(resp.Body, &user)

		ch <- getUserFromGitHubResponse{
			user: user,
			err:  err,
		}
	}()

	return ch
}

type ghUserEmailResponse struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

type getUserEmailFromGitHubResponse struct {
	email string
	err   error
}

func (s *authService) getUserEmailFromGitHub(
	accessToken string,
) <-chan getUserEmailFromGitHubResponse {
	ch := make(chan getUserEmailFromGitHubResponse, 1)

	go func() {
		req, err := http.NewRequest(
			http.MethodGet,
			"https://api.github.com/user/emails",
			nil,
		)
		if err != nil {
			ch <- getUserEmailFromGitHubResponse{email: "", err: err}
		}

		authTokenHeader := fmt.Sprintf("token %s", accessToken)
		req.Header.Set("Authorization", authTokenHeader)
		req.Header.Set("Accept", "application/json")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			ch <- getUserEmailFromGitHubResponse{email: "", err: err}
		}

		var emailsData []ghUserEmailResponse
		err = utils.FromJSON(res.Body, &emailsData)
		if err != nil {
			ch <- getUserEmailFromGitHubResponse{email: "", err: err}
		}

		for _, data := range emailsData {
			if data.Primary && data.Verified {
				ch <- getUserEmailFromGitHubResponse{email: data.Email, err: nil}
			}
		}

		ch <- getUserEmailFromGitHubResponse{
			email: "",
			err:   errors.New("could not retrieve an email"),
		}
	}()

	return ch
}

func (s *authService) loginWithGitHub(code string) (int, error) {
	userData, err := s.getUserDataFromGitHub(code)
	if err != nil {
		fmt.Println(err) // todo: use loggrus instance
		return 0, err
	}

	userObj, err := s.userService.GetByEmail(userData.Email)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return 0, err
	}

	// user exists: do nothing
	if err == nil {
		return userObj.Id, nil
	}

	// user does not exist: create user
	userToCreate := user.User{
		Email:     userData.Email,
		Name:      userData.Name,
		UserName:  userData.Login,
		AvatarUrl: userData.AvatarUrl,
	}

	return s.userService.Create(userToCreate)
}

func (s *authService) verifyRefreshToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt token signing method")
		}
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return strconv.Atoi(claims.Subject)
	}
	return 0, errors.New("token is invalid")
}
