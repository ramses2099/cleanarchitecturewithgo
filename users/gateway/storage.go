package gateway

import (
	"database/sql"

	"github.com/ramses2099/cleanarchitecturewithgo/users/model"
)

//const sql_select = "SELECT [Id],[Name],[Age],[CreateAt],[UpdateAt],[DelteAt] FROM [dbo].[Users]"

const sql_select = "SELECT [Id] FROM [dbo].[Users]"

type UserStorage interface {
	GetUsers() []*model.User
}

type UserService struct {
	db *sql.DB
}

func NewUserStorageGateway(db *sql.DB) UserStorage {
	return &UserService{db: db}
}

func (us *UserService) GetUsers() []*model.User {

	var users []*model.User

	// if us.db == nil {
	// 	fmt.Println("db is nil")
	// 	return nil
	// }

	// rows, err := us.db.Query(sql_select)
	// if err != nil {
	// 	log.Printf("cannot execute select query: %s", err.Error())
	// 	return nil
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var user model.User
	// 	// err := rows.Scan(&user.Id, &user.Name, &user.Age,
	// 	// 	&user.CreateAt, &user.UpdateAt, &user.DeleteAt)
	// 	err := rows.Scan(&user.Id)

	// 	if err != nil {
	// 		log.Println("cannot read current row")
	// 		return nil
	// 	}
	// 	users = append(users, &user)
	// }

	return users
}
