package api

import (
	"github.com/gorilla/mux"
	"github.com/ramses2099/cleanarchitecturewithgo/users/web"
)

func routes(services *web.UserHTTPService) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", services.GetUsersHandler)

	return r
}
