package feedback

// OK response
// swagger:response createFeedbackResponse
type createFeedbackResponseWrapper struct {
	// createFeedbackResponse
	// in: body
	Body struct {
		// id of a newly created feedback
		FeedbacId string `json:"feedbackId"`
	}
}

// OK response
// swagger:response okResponse
type okResponseResponseWrapper struct {
	// okResponse
	// in: body
	Body struct {
		// OK message
		Message string `json:"message"`
	}
}

// Not found feedback response
// swagger:response notFoundResponse
type notFoundResponse struct {
	// notFoundResponse
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// Bad request response
// swagger:response badRequestResponse
type badRequestResponse struct {
	// badRequestResponse
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// Returns all product feedback in the system
// swagger:response getAllFeedbackResponse
type getAllFeedbackResponseWrapper struct {
	// getAllFeedbackResponse
	// in: body
	Body []Feedback
}

// Returns product feedback by id
// swagger:response getFeedbackByIdResponse
type getFeedbackByIdResponseWrapper struct {
	// getFeedbackByIdResponse
	// in: body
	Body Feedback
}
