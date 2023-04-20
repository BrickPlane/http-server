package purchases_type

type Purchases struct {
	IdBuyer      uint64  `json:"id_buyer" db:"id_buyer"`
	IdGoods      uint64  `json:"id_goods" db:"id_goods"`
	MoneySpended *float64 `db:"money_spended"`
}
