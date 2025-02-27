package domain

type Order struct {
	OrderID     string      `json:"order_id"`
	CustomerID  string      `json:"customer_id"`
	OrderDate   string      `json:"order_date"`
	TotalAmount float64     `json:"total_amount"`
	OrderItems  []OrderItem `json:"order_items"`
}

type OrderItem struct {
	ProductID  string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}
