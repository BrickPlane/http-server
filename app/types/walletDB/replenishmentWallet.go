package wallet_types

type Replenishment struct {
	ID   uint64  `json:"id" db:"id"`
	Fill float64 `json:"fill" db:"fill"`
}
