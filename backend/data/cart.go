package data

import "time"

type Cart struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

type CartItem struct {
	ID        int     `json:"id"`
	CartID    int     `json:"cart_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	TotalCost float64 `json:"total_cost"`
}
