package user

// todo: add validation

type User struct {
	Id        int    `json:"id" db:"id"`
	Email     string `json:"email" db:"email" validate:"email"`
	Name      string `json:"name" db:"name" validate:"min=2,max=50"`
	UserName  string `json:"user_name" db:"user_name" validate:"min=2,max=50"`
	AvatarUrl string `json:"avatar_url" db:"avatar_url" validate:"base64url,omitempty"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
