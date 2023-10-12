package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/blackriper/manager/controllers"
	"github.com/blackriper/manager/utilities"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	AdminCon *controllers.AdminController
	AuthCon  *controllers.AuthController
	EmpCon   *controllers.EmpController
	ProCon   *controllers.ProController
}

func NewRoutes(adminCon *controllers.AdminController, authCon *controllers.AuthController, empCon *controllers.EmpController, proCon *controllers.ProController) *Routes {
	return &Routes{
		AdminCon: adminCon,
		AuthCon:  authCon,
		EmpCon:   empCon,
		ProCon:   proCon,
	}
}

func (rou *Routes) NewRouter() {
	r := gin.Default()

	// use middlewares
	r.Use(utilities.CorsConfig())
	authMidd := rou.AuthCon.JwtMiddleware()

	r.POST("/login", authMidd.LoginHandler)

	// routes
	api := r.Group("api")
	auth := r.Group("auth")

	auth.Use(authMidd.MiddlewareFunc())
	{
		auth.POST("/newadmin", rou.AdminCon.AdminCreate)
		auth.GET("/refresh_token", authMidd.RefreshHandler)
	}

	api.Use(authMidd.MiddlewareFunc())
	{
		api.POST("/newemployee", rou.EmpCon.CreateEmployee)
		api.GET("/employees", rou.EmpCon.GetEmployees)
		api.PUT("/updateemployee", rou.EmpCon.UpdateEmp)
		api.DELETE("/deleteemployee", rou.EmpCon.DeleteEmp)
		api.POST("/newproject", rou.ProCon.NewProject)
		api.DELETE("/deleteproject", rou.ProCon.DeleteProject)
	}

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
		log.Fatal(err)
	}
}
