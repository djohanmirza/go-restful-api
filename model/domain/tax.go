package domain

type Tax struct {
	TaxID       string  `json:"tax_id"`
	TaxRate     float64 `json:"tax_rate"` // Percentage value of the tax rate
	TaxType     string  `json:"tax_type"` // e.g., Sales Tax, VAT
	Description string  `json:"description"`
}
