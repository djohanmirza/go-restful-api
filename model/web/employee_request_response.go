package web

type EmployeeCreateRequest struct {
	Name      string `validate:"required,min=1,max=100" json:"name"`
	Role      string `validate:"required" json:"role"`
	Email     string `validate:"required,email" json:"email"`
	Phone     string `validate:"required" json:"phone"`
	DateHired string `json:"date_hired"`
}

type EmployeeResponse struct {
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	DateHired  string `json:"date_hired"`
}

type EmployeeUpdateRequest struct {
	EmployeeID string `validate:"required" json:"employee_id"`
	Name       string `validate:"required,max=100,min=1" json:"name"`
	Role       string `validate:"required" json:"role"`
	Email      string `validate:"required,email" json:"email"`
	Phone      string `validate:"required" json:"phone"`
	DateHired  string `json:"date_hired"`
}
