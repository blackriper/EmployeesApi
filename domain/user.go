package domain

import "github.com/kamva/mgm/v3"

// user model for admin users
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email" binding:"required"`
	UserName         string `json:"user_name" bson:"user_name" binding:"required"`
	Password         string `json:"password" bson:"password" binding:"required"`
	DepartamentId    string `json:"departament_id" bson:"departament_id" binding:"required"`
}

// port foradmin
type ForAdmin interface {
	CreateAdmin(user User) error
}

// driven repository admin
type RepositoryAdmin interface {
	CreateAdmin(user User) error
	ExistsAdmin(email string) bool
	GetAdmin(email string) (User, error)
}
