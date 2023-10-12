package repository

import (
	"errors"

	"github.com/blackriper/manager/domain"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}

func (p *ProjectRepository) CreatedProject(project domain.Project) error {
	result := p.DB.Create(&project)
	return result.Error
}

func (p *ProjectRepository) DeleteProject(idpro string) error {
	result := p.DB.Where("id_project = ? ", idpro).Delete(&domain.Project{})
	if result.RowsAffected == 0 {
		return errors.New("project not found")
	}
	return result.Error
}
