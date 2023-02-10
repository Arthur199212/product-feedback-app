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

// swagger:parameters getUserById
type userIdParam struct {
	// The id of the user for which the operation relates
	// in: path
	// required: true
	Id int `json:"id"`
}

// swagger:parameters getAllUsers
type userIdQueryParam struct {
	// Parameter 'userId' can be used to filter out
	// users by IDs.
	// It's possible to pass multimple IDs
	// (IDs should be comma delimited).
	// in: query
	// required: false
	// example: /api/users?userId=1,2,3,4,5
	UserId string `json:"userId"`
}

// Returns a list of all users in the system
// swagger:response getAllUsersResponse
type getAllUsersResponseWrapper struct {
	// getAllUsersResponse
	// in: body
	Body []User
}
