package domain

type Product struct {
<<<<<<< HEAD
	ProductID   string  `gorm:"primaryKey;column:product_id" json:"product_id"`
	Name        string  `gorm:"column:name" json:"name"`
	Description string  `gorm:"column:description" json:"description"`
	Price       float64 `gorm:"column:price" json:"price"`
	StockQty    int     `gorm:"column:stock_qty" json:"stock_qty"`
	Category    string  `gorm:"column:category" json:"category"`
	SKU         string  `gorm:"column:sku" json:"sku"`
	TaxRate     float64 `gorm:"column:tax_rate" json:"tax_rate"`
=======
	ProductID   uint64   `gorm:"primaryKey;column:id"`
	Name        string   `gorm:"column:product_name; length:255"`
	Description string   `gorm:"column:product_description; length:255"`
	Price       float64  `gorm:"column:product_price"`
	StockQty    int      `gorm:"column:stock_qty"`
	CategoryId  uint64   `gorm:"column:category_id"`
	SKU         string   `gorm:"column:product_sku"`
	TaxRate     float64  `gorm:"column:tax_rate"`
	Category    Category `gorm:"foreignKey:CategoryId;references:Id"`
}

type ProductError struct {
	Product Product `json:"product"`
	Error   error   `json:"error"`
>>>>>>> e1936bc3f26fda82e78e1c6def74a6990cf9675c
}
