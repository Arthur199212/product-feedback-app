package feedback

type Feedback struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Body      string `json:"body" db:"body"`
	Category  string `json:"category" db:"category"` // ui ux enchancement bug feature
	Status    string `json:"status" db:"status"`     // idea defined in-progress done
	UserId    int    `json:"userId" db:"user_id"`
	CreatedAt string `json:"createdAt" db:"created_at"`
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}
