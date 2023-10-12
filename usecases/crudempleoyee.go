package usecases

import (
	"github.com/blackriper/manager/domain"
)

type CrudEmp struct {
	RepoEmp domain.EmployeeRepository
}

func NewCrudEmp(repoEmp domain.EmployeeRepository) *CrudEmp {
	return &CrudEmp{
		RepoEmp: repoEmp,
	}
}

func (e *CrudEmp) NewEmp(emp domain.Employee) error {
	err := e.RepoEmp.NewEmp(emp)
	return err
}

func (e *CrudEmp) GetAllEmp(iddep string) ([]domain.Employee, error) {
	employees, err := e.RepoEmp.GetAllEmp(iddep)
	return employees, err
}

func (e *CrudEmp) UpdateEmp(idem string, emp domain.Employee) error {
	err := e.RepoEmp.UpdateEmp(idem, emp)
	return err
}

func (e *CrudEmp) DeleteEmp(idemp string) error {
	err := e.RepoEmp.DeleteEmp(idemp)
	return err
}
