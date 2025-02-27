package domain

type Discount struct {
	DiscountID  string  `json:"discount_id"`
	Description string  `json:"description"`
	DiscountPct float64 `json:"discount_pct"` // e.g., 10 for 10%
	ValidFrom   string  `json:"valid_from"`
	ValidUntil  string  `json:"valid_until"`
}
