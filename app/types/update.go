package types

type UpdateUserRequestDTO struct {
	ID       int     `json:"id" db:"id"`
	Email    *string `json:"email,omitempty" db:"email"`
	Login    *string `json:"login,omitempty" db:"login"`
	Password *string `json:"password,omitempty" db:"password"`
	Wallet   *string `json:"wallet" db:"wallet"`
}

type UpdateUserResponseDTO struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Login    string `json:"login" db:"login"`
	Wallet   string `json:"wallet" db:"wallet"`
	Password string `json:"-" db: "passsword"`
}
