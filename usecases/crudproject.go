package usecases

import "github.com/blackriper/manager/domain"

type CrudProj struct {
	RepoPro domain.ProjectRepository
}

func NewCrudProj(repoProj domain.ProjectRepository) *CrudProj {
	return &CrudProj{
		RepoPro: repoProj,
	}
}

func (p *CrudProj) CreatedProject(project domain.ProjectResquest) error {
	err := p.RepoPro.CreatedProject(domain.Project(project))
	return err
}

func (p *CrudProj) DeleteProject(idpro string) error {
	err := p.RepoPro.DeleteProject(idpro)
	return err
}
