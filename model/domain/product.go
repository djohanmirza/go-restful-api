package domain

type Product struct {
	ProductID   string   `gorm:"primaryKey;column:id"`
	Name        string   `gorm:"column:product_name; length:255"`
	Description string   `gorm:"column:product_description; length:255"`
	Price       float64  `gorm:"column:product_price"`
	StockQty    int      `gorm:"column:stock_qty"`
	CategoryId  int      `gorm:"column:category_id"`
	SKU         string   `gorm:"column:product_sku"`
	TaxRate     float64  `gorm:"column:tax_rate"`
	Category    Category `gorm:"foreignKey:CategoryId;references:Id"`
}

type ProductError struct {
	Product Product `json:"product"`
	Error   error   `json:"error"`
}
