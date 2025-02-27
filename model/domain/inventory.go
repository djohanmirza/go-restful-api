package domain

type Inventory struct {
	ProductID    string `json:"product_id"`
	StockQty     int    `json:"stock_qty"`
	RestockLevel int    `json:"restock_level"`
	LastRestock  string `json:"last_restock"`
}
