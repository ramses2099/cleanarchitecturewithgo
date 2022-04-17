package api

import (
	"github.com/ramses2099/cleanarchitecturewithgo/internal/storage"
	"github.com/ramses2099/cleanarchitecturewithgo/users/web"
)

func Start(port string) {
	db := storage.NewDB()
	defer db.Close()

	r := routes(web.NewUserHTTPService(db))
	server := NewServer(port, r)

	server.Start()
}
