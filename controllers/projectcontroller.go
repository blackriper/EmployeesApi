package controllers

import (
	"net/http"

	"github.com/blackriper/manager/domain"
	"github.com/gin-gonic/gin"
)

type ProController struct {
	UsePro domain.ForProject
}

func NewProyCon(usePro domain.ForProject) *ProController {
	return &ProController{UsePro: usePro}
}

func (p *ProController) NewProject(c *gin.Context) {
	var project domain.ProjectResquest
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := p.UsePro.CreatedProject(project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "proyect created successfully"})
}

func (p *ProController) DeleteProject(c *gin.Context) {
	idpro := c.Query("idpro")
	if err := p.UsePro.DeleteProject(idpro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "project delete successfully"})
}
