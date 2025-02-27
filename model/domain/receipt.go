package domain

type Receipt struct {
	ReceiptID   string  `json:"receipt_id"`
	OrderID     string  `json:"order_id"`
	PaymentID   string  `json:"payment_id"`
	ReceiptDate string  `json:"receipt_date"`
	TotalAmount float64 `json:"total_amount"`
	Taxes       float64 `json:"taxes"`
	Discount    float64 `json:"discount"`
	FinalAmount float64 `json:"final_amount"`
}
