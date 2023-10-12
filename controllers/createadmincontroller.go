package controllers

import (
	"net/http"

	"github.com/blackriper/manager/domain"
	"github.com/blackriper/manager/utilities"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	UseAdmin domain.ForAdmin
}

func NewAdminController(useAdmin domain.ForAdmin) *AdminController {
	return &AdminController{
		UseAdmin: useAdmin,
	}
}

func (ad *AdminController) AdminCreate(c *gin.Context) {
	var user domain.User
	// validate userdata
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user.Password = utilities.HashPassword(user.Password)
	err := ad.UseAdmin.CreateAdmin(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Admin created successfully"})
}
