package utilities

import (
	"log"
	"os"

	"github.com/blackriper/manager/domain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// function for loading .env files
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading.env file")
	}
}

// function to generate hash password
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// function to check password
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// function to configure cors settings
func CorsConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	return cors.New(config)

}

// conected with mongodb database
func ConnectMongoDB() {
	uri := os.Getenv("MONGO_URI")
	err := mgm.SetDefaultConfig(nil, "Admins", options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}

// conected with mysql database
func ConeectMySql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Departament{}, &domain.Employee{}, &domain.Project{})
	return db
}
