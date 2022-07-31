package user

// swagger:model User
type User struct {
	// Id of the user
	//
	// min: 1
	Id int `json:"id" db:"id" validate:"gt=0"`
	// Email of the user
	//
	// example: user@company.com
	Email string `json:"email" db:"email" validate:"email"`
	// Name of the user
	//
	// min length: 2
	// max length: 50
	Name string `json:"name" db:"name" validate:"min=2,max=50"`
	// User name
	//
	// min length: 2
	// max length: 50
	UserName string `json:"user_name" db:"user_name" validate:"min=2,max=50"`
	// Avatar URL
	AvatarUrl string `json:"avatar_url" db:"avatar_url" validate:"url,omitempty"`
	// Time the user was created at
	CreatedAt string `json:"created_at" db:"created_at"`
	// Time the user was updated at
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
