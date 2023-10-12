package repository

import (
	"github.com/blackriper/manager/domain"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Adminrepository struct{}

func (r *Adminrepository) CreateAdmin(user domain.User) error {
	if err := mgm.Coll(&user).Create(&user); err != nil {
		return err
	}
	return nil
}

func (r *Adminrepository) ExistsAdmin(email string) bool {
	count, _ := mgm.Coll(&domain.User{}).CountDocuments(mgm.Ctx(), bson.M{"email": email})
	return count > 0
}

func (r *Adminrepository) GetAdmin(email string) (domain.User, error) {
	var user domain.User
	if err := mgm.Coll(&user).First(bson.M{"email": email}, &user); err != nil {
		return user, err
	}
	return user, nil
}
