package domain

type Departament struct {
	IDDepartament string `gorm:"primaryKey" json:"id_departament"`
	Name          string `json:"name"`
}
