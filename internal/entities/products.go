package entities

type Product struct {
	Id    string
	Title string `json:"title"`
	Price int    `json:"price"`
}
