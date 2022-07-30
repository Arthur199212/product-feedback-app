package comment

// Create comment response
// swagger:response createCommentResponse
type createCommentResponseWrapper struct {
	// createCommentResponse
	// in: body
	Body struct {
		// id of a newly created comment
		CommentId string `json:"commentId"`
	}
}

// Returns a list of all comments in the system
// swagger:response getAllCommentsResponse
type getAllCommentsResponseWrapper struct {
	// getAllCommentsResponse
	// in: body
	Body []Comment
}

// OK response
// swagger:response okResponse
type okResponse struct {
	// OK response
	// in: body
	Body struct {
		Message string `json:"message"`
	}
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
