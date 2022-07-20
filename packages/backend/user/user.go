package user

// TODO: prepare migration or init User table in a proper way

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	UserName  string `json:"user_name"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
