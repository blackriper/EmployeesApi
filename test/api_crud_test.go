package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackriper/manager/controllers"
	"github.com/blackriper/manager/domain"
	"github.com/blackriper/manager/repository"
	"github.com/blackriper/manager/usecases"
	"github.com/blackriper/manager/utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// returning routing for testing employees
func ginHandler(emp *controllers.EmpController) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// use middlewares
	r.Use(utilities.CorsConfig())
	r.GET("/employees", emp.GetEmployees)
	r.POST("/newemployee", emp.CreateEmployee)
	r.PUT("/updateemployee", emp.UpdateEmp)
	r.DELETE("/deleteemployee", emp.DeleteEmp)
	return r
}

// returning routing for testing Projects
func ginHandlerP(pro *controllers.ProController) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// use middlewares
	r.Use(utilities.CorsConfig())
	r.POST("/newproject", pro.NewProject)
	r.DELETE("/deleteproject", pro.DeleteProject)
	return r
}

func TestCreateEmploye(t *testing.T) {

	// basic configuration
	utilities.LoadEnv()
	db := utilities.ConeectMySql()

	var empRepo domain.EmployeeRepository = &repository.EmpRepository{DB: db}
	useEmp := usecases.NewCrudEmp(empRepo)
	empCon := controllers.NewEmpController(useEmp)

	r := ginHandler(empCon)

	project := domain.Project{
		IDProject:   "6d9fd9ff-dcc5-47b3-82bf-57f3468e5401",
		Name:        "App IOS 17",
		Description: "app de IOS para renta de departamentos",
	}

	emp := domain.Employee{
		IDEmployee:  "6d9fd9ff-dcc5-47b3-82bf-57f3468e5401",
		Name:        "Hector",
		Position:    "Programador",
		Salary:      "$ 12,251",
		ContratDate: "2023-09-20",
		IDDep:       "b5a5988c-ddf5-4ffc-bfbb-3cb8524c8475",
		Project:     project,
	}

	t.Run("success", func(t *testing.T) {
		jsonVal, err := json.Marshal(emp)
		assert.NoError(t, err)

		req, _ := http.NewRequest("POST", "/newemployee", bytes.NewBuffer(jsonVal))
		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("error", func(t *testing.T) {

		jsonVal, err := json.Marshal(emp)
		assert.NoError(t, err)

		req, _ := http.NewRequest("POST", "/newemployee", bytes.NewBuffer(jsonVal))
		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("errorrequest", func(t *testing.T) {

		empl := domain.Employee{
			IDEmployee:  "cd4dbf47-8107-4f7a-9400-b3a26e1f973b",
			Name:        "Hector",
			Position:    "Programador",
			Salary:      "$ 12,251",
			ContratDate: "2023-09-20",
		}

		emplo, err := json.Marshal(empl)
		assert.NoError(t, err)

		req, _ := http.NewRequest("POST", "/newemployee", bytes.NewBuffer(emplo))
		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}

func TestGetEmpleoyees(t *testing.T) {

	// basic configuration
	utilities.LoadEnv()
	db := utilities.ConeectMySql()

	var empRepo domain.EmployeeRepository = &repository.EmpRepository{DB: db}
	useEmp := usecases.NewCrudEmp(empRepo)
	empCon := controllers.NewEmpController(useEmp)

	r := ginHandler(empCon)

	type Response struct {
		Body []domain.Employee
	}

	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/employees?iddep=b5a5988c-ddf5-4ffc-bfbb-3cb8524c8475", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var response Response
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, response.Body)
	})

	t.Run("error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/employees?iddep=b5a5988c-ddf5-4ffc-bfbb-3cb8524c", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var response Response
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Empty(t, response.Body)

	})

}

func TestUpdateEmployee(t *testing.T) {

	// basic configuration
	utilities.LoadEnv()
	db := utilities.ConeectMySql()

	var empRepo domain.EmployeeRepository = &repository.EmpRepository{DB: db}
	useEmp := usecases.NewCrudEmp(empRepo)
	empCon := controllers.NewEmpController(useEmp)

	r := ginHandler(empCon)

	emp := domain.Employee{
		IDEmployee:  "6d9fd9ff-dcc5-47b3-82bf-57f3468e5401",
		Name:        "Hector Leon",
		Position:    "Lider proyecto",
		Salary:      "$ 22,251",
		ContratDate: "2023-09-20",
		IDDep:       "b5a5988c-ddf5-4ffc-bfbb-3cb8524c8475",
	}

	t.Run("sucess", func(t *testing.T) {
		mockRes := `{"message":"employee updated successfully"}`
		jsonVal, err := json.Marshal(emp)
		assert.NoError(t, err)

		req, _ := http.NewRequest("PUT", fmt.Sprintf("/updateemployee?idemp=%v", emp.IDEmployee), bytes.NewBuffer(jsonVal))

		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusAccepted, w.Code)
		assert.Equal(t, mockRes, w.Body.String())

	})

	t.Run("error", func(t *testing.T) {

		emperr := domain.Employee{
			IDEmployee:  "67c90b0f-f8f5-4e5c-84b5-18c8c7b992cc",
			Name:        "Hector D Leon",
			Position:    "Lider proyecto",
			Salary:      "$ 22,251",
			ContratDate: "2023-09-20",
			IDDep:       "b5a5988c-ddf5-4ffc-bfbb-3cb8524c8475",
		}

		jsonV, err := json.Marshal(emperr)
		assert.NoError(t, err)

		req, _ := http.NewRequest("PUT", fmt.Sprintf("/updateemployee?idemp=%v", emperr.IDEmployee), bytes.NewBuffer(jsonV))

		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})

	t.Run("errorrequest", func(t *testing.T) {

		emperr := domain.Employee{
			IDEmployee:  "67c90b0f-f8f5-4e5c-84b5-18c8c7b992cc",
			Name:        "Hector D Leon",
			Position:    "Lider proyecto",
			Salary:      "$ 22,251",
			ContratDate: "2023-09-20",
		}

		jsonV, err := json.Marshal(emperr)
		assert.NoError(t, err)

		req, _ := http.NewRequest("PUT", fmt.Sprintf("/updateemployee?idemp=%v", emperr.IDEmployee), bytes.NewBuffer(jsonV))

		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

}

func TestDeleteEmployee(t *testing.T) {
	// basic configuration
	utilities.LoadEnv()
	db := utilities.ConeectMySql()

	var empRepo domain.EmployeeRepository = &repository.EmpRepository{DB: db}
	useEmp := usecases.NewCrudEmp(empRepo)
	empCon := controllers.NewEmpController(useEmp)

	r := ginHandler(empCon)
	iDEmployee := "6d9fd9ff-dcc5-47b3-82bf-57f3468e5401"

	t.Run("sucess", func(t *testing.T) {
		mockRes := `{"message":"employee delete successfully"}`

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/deleteemployee?idemp=%v", iDEmployee), nil)

		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusAccepted, w.Code)
		assert.Equal(t, mockRes, w.Body.String())

	})

	t.Run("error", func(t *testing.T) {

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/deleteemployee?idemp=%v", "7ec97545-d0cd-4e96-b3e6-4b0d979029ba"), nil)

		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})
}

func TestCreateProject(t *testing.T) {

	// basic configuration
	utilities.LoadEnv()
	db := utilities.ConeectMySql()

	var projectRepo domain.ProjectRepository = &repository.ProjectRepository{DB: db}
	usePro := usecases.NewCrudProj(projectRepo)
	proCon := controllers.NewProyCon(usePro)

	r := ginHandlerP(proCon)

	project := domain.ProjectResquest{
		IDProject:   "49a01ec4-b7e6-4a19-bbe9-945ecd803424",
		Name:        "App IOS 17",
		Description: "app de IOS para renta de departamentos",
	}

	t.Run("success", func(t *testing.T) {

		mockRes := `{"message":"proyect created successfully"}`

		jsonVal, err := json.Marshal(project)
		assert.NoError(t, err)

		req, _ := http.NewRequest("POST", "/newproject", bytes.NewBuffer(jsonVal))
		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, mockRes, w.Body.String())
	})

	t.Run("error", func(t *testing.T) {

		jsonVal, err := json.Marshal(project)
		assert.NoError(t, err)

		req, _ := http.NewRequest("POST", "/newproject", bytes.NewBuffer(jsonVal))
		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("errorrequest", func(t *testing.T) {

		projecterr := domain.ProjectResquest{
			Name:        "App IOS 17",
			Description: "app de IOS para renta de departamentos",
		}

		emplo, err := json.Marshal(projecterr)
		assert.NoError(t, err)

		req, _ := http.NewRequest("POST", "/newproject", bytes.NewBuffer(emplo))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}

func TestDeleteProyect(t *testing.T) {
	// basic configuration
	utilities.LoadEnv()
	db := utilities.ConeectMySql()

	var projectRepo domain.ProjectRepository = &repository.ProjectRepository{DB: db}
	usePro := usecases.NewCrudProj(projectRepo)
	proCon := controllers.NewProyCon(usePro)

	r := ginHandlerP(proCon)
	iDProject := "49a01ec4-b7e6-4a19-bbe9-945ecd803424"

	t.Run("sucess", func(t *testing.T) {
		mockRes := `{"message":"project delete successfully"}`

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/deleteproject?idpro=%v", iDProject), nil)

		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusAccepted, w.Code)
		assert.Equal(t, mockRes, w.Body.String())

	})

	t.Run("error", func(t *testing.T) {

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/deleteproject?idpro=%v", "7ec97545-d0cd-4e96-b3e6-4b0d979029ba"), nil)

		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})
}
