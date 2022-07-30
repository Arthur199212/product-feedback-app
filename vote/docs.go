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

