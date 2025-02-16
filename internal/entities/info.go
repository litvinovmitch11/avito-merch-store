package entities

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type ReceivedItem struct {
	FromUser string
	Amount   int
}

type SentItem struct {
	ToUser *string
	Amount int
}

type CoinHistory struct {
	Received []ReceivedItem
	Sent     []SentItem
}

type Inventory = []InventoryItem

type Info struct {
	Coins       int
	Inventory   Inventory
	CoinHistory CoinHistory
}
