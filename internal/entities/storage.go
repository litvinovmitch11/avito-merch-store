package entities

type SendCoin struct {
	Amount   int    `json:"amount"`
	ToUser   string `json:"toUser"`
	FromUser string
}
