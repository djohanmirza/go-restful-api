package domain

type Customer struct {
<<<<<<< HEAD
	CustomerID string `gorm:"primaryKey;column:customer_id" json:"customer_id"`
	Name       string `gorm:"column:name" json:"name"`
	Email      string `gorm:"column:email" json:"email"`
	Phone      string `gorm:"column:phone" json:"phone"`
	Address    string `gorm:"column:address" json:"address"`
	LoyaltyPts int    `gorm:"column:loyalty_points" json:"loyalty_points"`
=======
	CustomerID uint64 `gorm:"primary_key;column:id;autoIncrement"`
	Name       string `gorm:"column:customer_name; type:varchar(100);"`
	Email      string `gorm:"column:customer_email; type:varchar(255);"`
	Phone      string `gorm:"column:customer_phone; type:varchar(20);"`
	Address    string `gorm:"column:customer_address; type:varchar(255);"`
	LoyaltyPts int    `gorm:"column:loyalty_pts; type:int(11);"`
>>>>>>> e1936bc3f26fda82e78e1c6def74a6990cf9675c
}
