package entities

type SendCoin struct {
	Amount   int    `json:"amount"`
	ToUser   string `json:"toUser"`
	FromUser string
}

type Balance struct {
	ID     string
	UserID string
	Amount int
}
