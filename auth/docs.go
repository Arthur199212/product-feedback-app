package auth

// Redirects to login route
// swagger:response foundResponse
type foundResponse struct {
}

// Error response
// swagger:response errorResponse
type errorResponse struct {
	// Error response
	// in: body
	Body struct {
		// Error message
		Message string `json:"message"`
	}
}

// OK response
// swagger:response refreshAccessTokenResponse
type refreshAccessTokenResponseWrapper struct {
	// refreshAccessTokenResponse
	// in: body
	Body struct {
		// Access token
		AccessToken string `json:"accessToken"`
	}
}

// swagger:parameters loginWithGitHub
type codeQueryParam struct {
	// Code to handle login with GitHub.
	// GitHub itself provides it.
	// in: query
	// required: true
	Code string `json:"code"`
}

// swagger:parameters refreshAccessToken
type refreshAccessTokenParamsWrapper struct {
	// Refresh token is used to get access token
	// in: body
	// required: true
	Body refreshAccessTokenInput
}
