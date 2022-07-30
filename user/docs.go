package user

// OK response
// swagger:response getUserResponse
type getUserResponseWrapper struct {
	// User data
	// in: body
	Body User
}

// Error response
// swagger:response errorResponse
type errorResponse struct {
	// Error message
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}
