package helper

import (
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

func ToEmployeeResponse(employee domain.Employee) web.EmployeeResponse {
	return web.EmployeeResponse{
		EmployeeID: employee.EmployeeID,
		Name:       employee.Name,
		Role:       employee.Role,
		Email:      employee.Email,
		Phone:      employee.Phone,
		DateHired:  employee.DateHired,
	}
}

func ToEmployeeResponses(employees []domain.Employee) []web.EmployeeResponse {
	var employeeResponses []web.EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, ToEmployeeResponse(employee))
	}
	return employeeResponses
}
