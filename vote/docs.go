package vote

// OK response
// swagger:response getAllVotesResponse
type getAllVotesResponseWrapper struct {
	// A list of votes
	// in: body
	Body []Vote
}

// OK response
// swagger:response createVoteResponse
type createVoteResponseWrapper struct {
	// createVoteResponse
	// in: body
	Body struct {
		// id of a newly created vote
		VoteId string `json:"voteId"`
	}
}

// OK response
// swagger:response okResponse
type okResponse struct {
	// okResponse
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

// swagger:parameters deleteVote
type voteIdParam struct {
	// The id of the vote for which the operation relates
	// in: path
	// required: true
	Id int `json:"id"`
}

// swagger:parameters createVote
type createVoteInputParamsWrapper struct {
	// Vote data structure to create vote
	// in: body
	// required: true
	Body createVoteInput
}

// swagger:parameters getAllVotes
type feedbackIdQueryParam struct {
	// Feedback id can be used to filter out
	// votes by feedback they relate to.
	// It's possible to pass multimple feedback ids
	// (ids should be comma delimited).
	// in: query
	// required: false
	// example: /api/votes/?feedbackId=1,2
	FeedbackId string `json:"feedbackId"`
}

// swagger:parameters toggleVote
type toggleVoteInputParamsWrapper struct {
	// Data structure to toggle vote
	// in: body
	// required: true
	Body toggleVoteInput
}
