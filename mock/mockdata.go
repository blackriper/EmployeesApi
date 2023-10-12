package mock

import (
	"math/rand"
	"strconv"
	"sync"

	"github.com/blackriper/manager/domain"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MockData struct {
	DB *gorm.DB
}

// data mock static
var depNames []string = []string{"Ingenieria", "Sistemas", "Diseño"}

func (m *MockData) CreateMock() {
	//specific time to finish goroutines
	wg := sync.WaitGroup{}
	deps := make(chan []domain.Departament)

	wg.Add(2)
	go CreateDepartaments(deps, m.DB, &wg)
	go CreateEmployees(deps, m.DB, &wg)
	wg.Wait()
	close(deps)
}

// chan<- only send || <-chan only recive
func CreateDepartaments(deps chan<- []domain.Departament, db *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	var deparments []domain.Departament
	for _, name := range depNames {
		deparment := domain.Departament{
			IDDepartament: uuid.New().String(),
			Name:          name,
		}
		deparments = append(deparments, deparment)
	}
	db.Create(deparments)
	deps <- deparments

}

func CreateEmployees(deps <-chan []domain.Departament, db *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	var empleoyees []domain.Employee
	var i int8
	deparments := <-deps
	for _, deparment := range deparments {

		for i < 15 {
			salary := strconv.Itoa(rand.Intn(10000-1000) + 1000)
			idemp := uuid.New().String()

			empleoye := domain.Employee{
				IDEmployee:  idemp,
				Name:        faker.FirstName(),
				Position:    RandomPosition(deparment.Name),
				Salary:      "$ " + salary,
				ContratDate: faker.Date(),
				IDDep:       deparment.IDDepartament,
			}

			if i == 7 {
				project := RandomProject(deparment.Name, idemp)
				empleoye.Project = project
			}
			empleoyees = append(empleoyees, empleoye)
			i++
		}
		i = 0
		db.Create(empleoyees)
		empleoyees = nil
	}

}

// selected employee positions

func RandomPosition(name string) string {
	var positions []string
	switch {
	case name == "Ingenieria":
		positions = []string{"Mecanico", "Instrumentista", "Electrico"}
	case name == "Sistemas":
		positions = []string{"Programador", "Testing", "Lider proyecto"}
	case name == "Diseño":
		positions = []string{"Ui/Ux", "Maquetador", "Diseñador web"}
	}
	randomindex := rand.Intn(len(positions))
	return positions[randomindex]
}

// generate proyect mock data

func RandomProject(name string, idemp string) domain.Project {
	var nameP string
	var description string

	switch {
	case name == "Ingenieria":
		nameP = "Nueva trasmision"
		description = "Trasmicion nueva para el nuevo modelo de automovil"
	case name == "Sistemas":
		nameP = "App web consultas"
		description = "App web para consultar datos de pagos de los clientes"
	case name == "Diseño":
		nameP = "Maquetar app movil"
		description = "crear mock para app movil"
	}

	project := domain.Project{
		IDProject:   idemp,
		Name:        nameP,
		Description: description,
	}
	return project

}
