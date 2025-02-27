package domain

type Employee struct {
	EmployeeID string `gorm:"column:id;primary_key"`
	Name       string `gorm:"column:name"`
	Role       string `gorm:"column:role"` // e.g., Cashier, Manager
	Email      string `gorm:"column:email"`
	Phone      string `gorm:"column:phone"`
	DateHired  string `gorm:"column:date_hired"`
}
