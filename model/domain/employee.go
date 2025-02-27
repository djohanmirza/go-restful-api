package domain

type Employee struct {
<<<<<<< HEAD
	EmployeeID string `gorm:"primaryKey;column:employee_id" json:"employee_id"`
	Name       string `gorm:"column:name" json:"name"`
	Role       string `gorm:"column:role" json:"role"`
	Email      string `gorm:"column:email" json:"email"`
	Phone      string `gorm:"column:phone" json:"phone"`
	DateHired  string `gorm:"column:date_hired" json:"date_hired"`
=======
	EmployeeID uint64 `gorm:"column:id;primary_key"`
	Name       string `gorm:"column:name"`
	Role       string `gorm:"column:role"` // e.g., Cashier, Manager
	Email      string `gorm:"column:email"`
	Phone      string `gorm:"column:phone"`
	DateHired  string `gorm:"column:date_hired"`
>>>>>>> e1936bc3f26fda82e78e1c6def74a6990cf9675c
}
