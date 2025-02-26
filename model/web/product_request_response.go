package web

type ProductCreateRequest struct {
	Name        string  `validate:"required,min=1,max=100" json:"name"`
	Description string  `json:"description"`
	Price       float64 `validate:"required,gt=0" json:"price"`
	StockQty    int     `validate:"required,gte=0" json:"stock_qty"`
	Category    string  `json:"category"`
	SKU         string  `json:"sku"`
	TaxRate     float64 `json:"tax_rate"`
}

type ProductResponse struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
	Category    string  `json:"category"`
	SKU         string  `json:"sku"`
	TaxRate     float64 `json:"tax_rate"`
}

type ProductUpdateRequest struct {
	ProductID   string  `validate:"required" json:"product_id"`
	Name        string  `validate:"required,max=100,min=1" json:"name"`
	Description string  `json:"description"`
	Price       float64 `validate:"required,gt=0" json:"price"`
	StockQty    int     `validate:"required,gte=0" json:"stock_qty"`
	Category    string  `json:"category"`
	SKU         string  `json:"sku"`
	TaxRate     float64 `json:"tax_rate"`
}
