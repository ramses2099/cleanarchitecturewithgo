package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

const user string = "sa"
const password string = "P@ssword01"
const port string = "1433"
const database string = "dbtest"
const driverName string = "sqlserver"

func NewDB() *sql.DB {
	connectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, database)

	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		fmt.Println(fmt.Errorf("Error opening database %s", err.Error()))
		return nil
	}

	return db
}
