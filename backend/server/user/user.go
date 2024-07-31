package user

// Models for user
type User struct {
	ID int64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email`
	Password string `json:"password" db:"password"`
}

// Request and response models
type CreateUserRequest struct {
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email`
	Password string `json:"password" db:"password"`
}

type CreateUserResponse struct {
	ID string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email`
}

type LoginUserRequest struct {
	Email string `json:"email" db:"email`
	Password string `json:"password" db:"password"`
}

type LoginUserResponse struct {
	accessToken string
	ID string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}
