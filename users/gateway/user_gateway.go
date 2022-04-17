package gateway

import (
	"database/sql"
	"fmt"

	"github.com/ramses2099/cleanarchitecturewithgo/users/model"
)

type UserGateway interface {
	GetUsers() []*model.User
}

type CreateUserInDB struct {
	UserStorage
}

func (c *CreateUserInDB) GetUsers() []*model.User {
	fmt.Println("ok")
	return nil
}

func NewUserGateway(db *sql.DB) UserGateway {
	return &CreateUserInDB{NewUserStorageGateway(db)}
}
