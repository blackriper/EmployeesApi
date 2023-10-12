package controllers

import (
	"net/http"

	"github.com/blackriper/manager/domain"
	"github.com/gin-gonic/gin"
)

type EmpController struct {
	UseEmp domain.ForEmployee
}

func NewEmpController(useEmp domain.ForEmployee) *EmpController {
	return &EmpController{
		UseEmp: useEmp,
	}
}

func (e *EmpController) CreateEmployee(c *gin.Context) {
	var employee domain.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := e.UseEmp.NewEmp(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "employee created successfully"})

}

func (e *EmpController) GetEmployees(c *gin.Context) {
	iddep := c.Query("iddep")
	if iddep == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "parameter iddep incorrect"})
		return
	}
	empleoyess, err := e.UseEmp.GetAllEmp(iddep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"body": empleoyess})

}

func (e *EmpController) UpdateEmp(c *gin.Context) {
	var employee domain.Employee
	idemp := c.Query("idemp")

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := e.UseEmp.UpdateEmp(idemp, employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "employee updated successfully"})
}

func (e *EmpController) DeleteEmp(c *gin.Context) {
	idemp := c.Query("idemp")
	if err := e.UseEmp.DeleteEmp(idemp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "employee delete successfully"})
}
