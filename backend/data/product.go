package data

type Product struct {
	Id       int     `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Category string  `json:"category"`
	Rating   float32 `json:"rating"`
}
