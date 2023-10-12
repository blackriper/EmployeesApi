package repository

import (
	"errors"

	"github.com/blackriper/manager/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmpRepository struct {
	DB *gorm.DB
}

func NewEmpRepository(db *gorm.DB) *EmpRepository {
	return &EmpRepository{DB: db}
}

func (e *EmpRepository) NewEmp(emp domain.Employee) error {
	result := e.DB.Create(&emp)
	return result.Error
}

func (e *EmpRepository) GetAllEmp(iddep string) ([]domain.Employee, error) {
	var employees []domain.Employee
	result := e.DB.Model(&domain.Employee{}).Where("id_dep = ?", iddep).Preload("Project").Find(&employees)
	return employees, result.Error
}

func (e *EmpRepository) UpdateEmp(idem string, emp domain.Employee) error {
	result := e.DB.Session(&gorm.Session{FullSaveAssociations: true}).Where("id_employee = ? ", idem).Updates(&emp)

	if result.RowsAffected == 0 {
		return errors.New("empleoyee not found")
	}

	return result.Error
}

func (e *EmpRepository) DeleteEmp(idemp string) error {
	var employe domain.Employee

	result := e.DB.Model(&domain.Employee{}).Where("id_employee = ?", idemp).Preload("Project").Find(&employe)

	if result.Error != nil {
		return result.Error
	}
	res := e.DB.Select(clause.Associations).Delete(&employe)
	return res.Error
}
