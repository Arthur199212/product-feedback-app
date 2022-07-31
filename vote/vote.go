package vote

// swagger:model Vote
type Vote struct {
	// Id of the vote
	//
	// min: 1
	Id int `json:"id" db:"id"`
	// Feedback id that this vote is related to
	//
	// min: 1
	FeedbackId int `json:"feedbackId" db:"feedback_id"`
	// Id of a user that created this vote
	//
	// min: 1
	UserId int `json:"userId" db:"user_id"`
	// Time this vote was created at
	CreatedAt string `json:"createdAt" db:"created_at"`
	// Time this vote was updated at
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}
