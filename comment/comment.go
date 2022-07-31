package comment

// swagger:model Comment
type Comment struct {
	// Id of the comment
	//
	// min: 1
	Id string `json:"id" db:"id"`
	// Body of the comment
	//
	// min length: 5
	// max length: 255
	Body string `json:"body" db:"body"`
	// Id of the feedback this comment is related to
	//
	// min: 1
	FeedbackId int `json:"feedbackId" db:"feedback_id"`
	// Id of the user who created the comment
	//
	// min: 1
	UserId int `json:"userId" db:"user_id"`
	// Time comment was created at
	CreatedAt string `json:"createdAt" db:"created_at"`
	// Time comment was updated at
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}
