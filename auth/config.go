package auth

import "time"

const (
	accessTokenTTL  = 30 * time.Minute
	refreshTokenTTL = 72 * time.Hour

	ghLoginOauthAccessTokenURI = "https://github.com/login/oauth/access_token"
	ghLoginOauthAuthorizeURI   = "https://github.com/login/oauth/authorize"
	ghRedirectURI              = "/api/auth/github/callback"
	ghUserEmailsURI            = "https://api.github.com/user/emails"
	ghUserEmailScope           = "user:email"
)
