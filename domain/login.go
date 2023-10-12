package domain

import (
	jwt "github.com/appleboy/gin-jwt/v2"
)

// user for authentication
type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// user logged in application
type UserLogged struct {
	UserName      string `json:"user_name"`
	DepartamentId string `json:"departament_id"`
	Token         string `json:"token"`
	Expire        string `json:"expire"`
}

// port for authentication
type ForLogin interface {
	GetJwtMiddleware() *jwt.GinJWTMiddleware
}

// driven for login authentication
type RepositoryLogin interface {
	GetJwtMidd() *jwt.GinJWTMiddleware
}
