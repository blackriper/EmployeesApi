package usecases

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/blackriper/manager/domain"
)

type AuthAdmin struct {
	AuthRepo domain.RepositoryLogin
}

func NewAuthAdmin(authRepo domain.RepositoryLogin) *AuthAdmin {
	return &AuthAdmin{
		AuthRepo: authRepo,
	}
}

func (a *AuthAdmin) GetJwtMiddleware() *jwt.GinJWTMiddleware {
	return a.AuthRepo.GetJwtMidd()
}
