package domain

// model represents employees
type Employee struct {
	IDEmployee  string  `gorm:"primaryKey" json:"id_employee" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Position    string  `json:"position" binding:"required"`
	Salary      string  `json:"salary" binding:"required"`
	ContratDate string  `json:"contrat_date" binding:"required"`
	IDDep       string  `json:"id_dep" binding:"required"`
	Project     Project `gorm:"foreignKey:IDProject" json:"project"`
}

// port for employees
type ForEmployee interface {
	NewEmp(emp Employee) error
	GetAllEmp(iddep string) ([]Employee, error)
	UpdateEmp(idem string, emp Employee) error
	DeleteEmp(idem string) error
}

// driven for employees
type EmployeeRepository interface {
	NewEmp(emp Employee) error
	GetAllEmp(iddep string) ([]Employee, error)
	UpdateEmp(idem string, emp Employee) error
	DeleteEmp(idem string) error
}
