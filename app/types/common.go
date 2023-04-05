package types

type CommonRequest struct {
	UserID
}

type UserID struct {
	ID uint64 `json:"id" db:"id"`
}
