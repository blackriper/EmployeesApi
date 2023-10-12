package main

import (
	"github.com/blackriper/manager/controllers"
	"github.com/blackriper/manager/domain"
	"github.com/blackriper/manager/repository"
	"github.com/blackriper/manager/routes"
	"github.com/blackriper/manager/usecases"

	//"github.com/blackriper/manager/mock"
	"github.com/blackriper/manager/utilities"
)

func main() {
	// basic configuration
	utilities.LoadEnv()
	utilities.ConnectMongoDB()
	db := utilities.ConeectMySql()
	// create basic mock data for database
	//mock := mock.MockData{DB: db}
	//mock.CreateMock()

	// create user admins
	var userRepo domain.RepositoryAdmin = &repository.Adminrepository{}
	useAdmin := usecases.NewCreateAdmin(userRepo)
	adminCon := controllers.NewAdminController(useAdmin)
	// login  user admin
	authRepo := repository.NewAuthRepository(userRepo)
	useAuth := usecases.NewAuthAdmin(authRepo)
	authCon := controllers.NewAuthController(useAuth)
	// crud operations for admin empleoyess
	var empRepo domain.EmployeeRepository = &repository.EmpRepository{DB: db}
	useEmp := usecases.NewCrudEmp(empRepo)
	empCon := controllers.NewEmpController(useEmp)
	// crud operations for project
	var projectRepo domain.ProjectRepository = &repository.ProjectRepository{DB: db}
	usePro := usecases.NewCrudProj(projectRepo)
	proCon := controllers.NewProyCon(usePro)

	routes := routes.NewRoutes(adminCon, authCon, empCon, proCon)
	routes.NewRouter()

}
