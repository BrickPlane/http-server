package types

type UpdateUserRequestDTO struct {
	ID       int     `json:"id"`
	Email    *string `json:"email,omitempty"`
	Login    *string `json:"login,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UpdateUserResponseDTO struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Login string `json:"login"`
}
