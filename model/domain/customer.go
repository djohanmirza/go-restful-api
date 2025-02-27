package domain

type Customer struct {
	CustomerID uint64 `gorm:"primary_key;column:id;autoIncrement"`
	Name       string `gorm:"column:customer_name; type:varchar(100);"`
	Email      string `gorm:"column:customer_email; type:varchar(255);"`
	Phone      string `gorm:"column:customer_phone; type:varchar(20);"`
	Address    string `gorm:"column:customer_address; type:varchar(255);"`
	LoyaltyPts int    `gorm:"column:loyalty_pts; type:int(11);"`
}
