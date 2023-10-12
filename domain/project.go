package domain

type Project struct {
	IDProject   string `gorm:"primaryKey" json:"id_project"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectResquest struct {
	IDProject   string `json:"id_project" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// driver proyect repository
type ForProject interface {
	CreatedProject(proyect ProjectResquest) error
	DeleteProject(idpro string) error
}

// driven proyect repository
type ProjectRepository interface {
	CreatedProject(proyect Project) error
	DeleteProject(idpro string) error
}
