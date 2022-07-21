package feedback

type Feedback struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Body      string `json:"body" db:"body"`
	// ui ux enchancement bug feature
	Category  string `json:"category" db:"category"`
	// idea, defined, in-progress, done
	Stauts    string `json:"status" db:"status"`
	UserId    int    `json:"userId" db:"user_id"`
	CreatedAt string `json:"createdAt" db:"created_at"`
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}
