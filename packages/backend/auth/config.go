package auth

import "time"

const (
	ghLoginOauthAccessTokenURI = "https://github.com/login/oauth/access_token"
	ghLoginOauthAuthorizeURI   = "https://github.com/login/oauth/authorize"
	ghRedirectURI              = "http://localhost:8000/api/auth/github/callback"
	ghUserEmailsURI            = "https://api.github.com/user/emails"
	ghUserEmailScope           = "user:email"

	// accessTokenTTL         = 30 * time.Minute
	// todo: remove
	accessTokenTTL         = 168 * time.Hour // for test purposes
	refreshTokenCookieName = "refresh-token"
	refreshTokenRoute      = "/api/auth/refresh-token"
	refreshTokenTTL        = 72 * time.Hour
)
