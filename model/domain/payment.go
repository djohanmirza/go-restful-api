package domain

type Payment struct {
	PaymentID   string  `json:"payment_id"`
	OrderID     string  `json:"order_id"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"payment_type"` // e.g., Cash, Card, Online
	PaymentDate string  `json:"payment_date"`
	Status      string  `json:"status"` // e.g., Completed, Pending
}
