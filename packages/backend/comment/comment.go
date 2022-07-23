package comment

type Comment struct {
	Id         string `json:"id" db:"id"`
	Body       string `json:"body" db:"body"`
	FeedbackId int    `json:"feedbackId" db:"feedback_id"`
	UserId     int    `json:"userId" db:"user_id"`
	CreatedAt  string `json:"createdAt" db:"created_at"`
	UpdatedAt  string `json:"updatedAt" db:"updated_at"`
}
