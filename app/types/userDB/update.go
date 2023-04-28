package user_types

type UpdateUserRequestDTO struct {
	ID       int     `json:"id" db:"id"`
	Email    *string `json:"email,omitempty" db:"email"`
	Login    *string `json:"login,omitempty" db:"login"`
	Password *string `json:"password,omitempty" db:"password"`
	Wallet   *float64 `json:"wallet" db:"wallet"`
}

type UpdateUserResponseDTO struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Login    string `json:"login" db:"login"`
	Wallet   float64 `json:"wallet" db:"wallet"`
	Password string `json:"password" db:"passsword"`
}
