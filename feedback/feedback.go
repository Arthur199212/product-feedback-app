package feedback

// swagger:model Feedback
type Feedback struct {
	// Id of the feedback
	//
	// min: 1
	Id int `json:"id" db:"id"`
	// Title of the feedback
	//
	// min length: 5
	// max length: 50
	Title string `json:"title" db:"title"`
	// Body of the feedback
	//
	// min length: 10
	// max length: 1000
	Body string `json:"body" db:"body"`
	// Category of the feedback
	//
	// Possible categories: 'ui', 'ux', 'enchancement', 'bug', 'feature'
	Category string `json:"category" db:"category"`
	// Status of the feedback
	//
	// Possible statuses: 'idea', 'defined', 'in-progress', 'done'
	Status string `json:"status" db:"status"`
	// Id of a user who created a feedback
	//
	// min: 1
	UserId int `json:"userId" db:"user_id"`
	// Time feedback was created at
	CreatedAt string `json:"createdAt" db:"created_at"`
	// Time feedback was updated at
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}
