package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/blackriper/manager/domain"
)

type AuthController struct {
	UseAuth domain.ForLogin
}

func NewAuthController(useAuth domain.ForLogin) *AuthController {
	return &AuthController{
		UseAuth: useAuth,
	}
}

func (a *AuthController) JwtMiddleware() *jwt.GinJWTMiddleware {
	return a.UseAuth.GetJwtMiddleware()
}
