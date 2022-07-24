package vote

type Vote struct {
	Id         int    `json:"id" db:"id"`
	FeedbackId int    `json:"feedbackId" db:"feedback_id"`
	UserId     int    `json:"userId" db:"user_id"`
	CreatedAt  string `json:"createdAt" db:"created_at"`
	UpdatedAt  string `json:"updatedAt" db:"updated_at"`
}
