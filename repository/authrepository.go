package repository

import (
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/blackriper/manager/domain"
	"github.com/blackriper/manager/utilities"
	"github.com/gin-gonic/gin"
)

type AuthRepository struct {
	UserRepo domain.RepositoryAdmin
}

func NewAuthRepository(userRepo domain.RepositoryAdmin) *AuthRepository {
	return &AuthRepository{
		UserRepo: userRepo,
	}
}

func (r *AuthRepository) GetJwtMidd() *jwt.GinJWTMiddleware {
	var (
		login       domain.UserLogin
		userlogged  domain.UserLogged
		identityKey = "id"
	)

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test",
		Key:         []byte(utilities.HashPassword("MANAGER")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {

			if err := c.ShouldBindJSON(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			success := NewLoggedUser(r.UserRepo, &userlogged, login)
			if success {
				return userlogged, nil
			} else if !success {
				return nil, errors.New("admin user not exist")

			}

			return nil, jwt.ErrFailedAuthentication
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			userlogged.Token = token
			userlogged.Expire = expire.Format(time.RFC3339)
			c.JSON(http.StatusOK, gin.H{"message": userlogged})

		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	// inicialize middleware
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}

// create user logged
func NewLoggedUser(repo domain.RepositoryAdmin, admin *domain.UserLogged, data domain.UserLogin) bool {
	var logged bool
	user, err := repo.GetAdmin(data.Email)
	if err != nil {
		logged = false
	}
	if (data.Email == user.Email) && (utilities.CheckPassword(data.Password, user.Password)) {
		*admin = domain.UserLogged{
			UserName:      user.UserName,
			DepartamentId: user.DepartamentId,
		}
		logged = true
	}

	return logged
}
