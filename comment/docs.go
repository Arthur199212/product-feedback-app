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

// Returns a comment with the specified Id
// swagger:response getCommentByIdResponse
type getCommentByIdResponseWrapper struct {
	// getCommentByIdResponse
	// in: body
	Body Comment
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

// swagger:parameters getAllComments
type feedbackIdQueryParam struct {
	// Feedback id can be used to filter out
	// comments by feedback they relate to.
	// in: query
	// required: false
	FeedbackId string `json:"feedbackId"`
}

// swagger:parameters getCommentById
type commentIdParamsWrapper struct {
	// The id of the commnet for which the operation relates
	// in: path
	// required: true
	Id int
}

// swagger:parameters createComment
type createCommentInputParamsWrapper struct {
	// Comment data structure to create comment
	// in: body
	// required: true
	Body createCommentInput
}
