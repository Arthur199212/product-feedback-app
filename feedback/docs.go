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

// swagger:response errorResponse
type errorResponse struct {
	// Message of the error
	// in: string
	Message string `json:"message"`
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

// swagger:parameters createFeedback
type createFeedbackInputParamsWrapper struct {
	// Feedback data structure to create feedback
	// in: body
	// required: true
	Body createFeedbackInput
}

// swagger:parameters updateFeedback
type updateFeedbackInputParamsWrapper struct {
	// Feedback data structure to update feedback
	// in: body
	// required: true
	Body updateFeedbackInput
}

// swagger:parameters deleteFeedback getFeedbackById updateFeedback
type feedbackIdParamsWrapper struct {
	// The id of the feedback for which the operation relates
	// in: path
	// required: true
	Id int `json:"id"`
}
