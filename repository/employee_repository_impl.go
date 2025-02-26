package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
}

func (repository *EmployeeRepositoryImpl) Save(ctx context.Context, employee domain.Employee) (domain.Employee, error) {
	if err := repository.db.WithContext(ctx).Create(&employee).Error; err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

func (repository *EmployeeRepositoryImpl) Update(ctx context.Context, employee domain.Employee) (domain.Employee, error) {
	if err := repository.db.WithContext(ctx).Save(&employee).Error; err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

func (repository *EmployeeRepositoryImpl) Delete(ctx context.Context, employee domain.Employee) error {
	return repository.db.WithContext(ctx).Delete(&employee).Error
}

func (repository *EmployeeRepositoryImpl) FindById(ctx context.Context, employeeId string) (domain.Employee, error) {
	var employee domain.Employee
	err := repository.db.WithContext(ctx).First(&employee, "employee_id = ?", employeeId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return employee, errors.New("employee not found")
	}
	return employee, err
}

func (repository *EmployeeRepositoryImpl) FindAll(ctx context.Context) ([]domain.Employee, error) {
	var employees []domain.Employee
	return employees, repository.db.WithContext(ctx).Find(&employees).Error
}
