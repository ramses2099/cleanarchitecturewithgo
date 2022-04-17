package web

import (
	"database/sql"
	"net/http"

	"github.com/ramses2099/cleanarchitecturewithgo/internal/web"
	"github.com/ramses2099/cleanarchitecturewithgo/users/gateway"
	"github.com/ramses2099/cleanarchitecturewithgo/users/model"
)

type UserHTTPService struct {
	gtw gateway.UserGateway
}

func NewUserHTTPService(db *sql.DB) *UserHTTPService {
	return &UserHTTPService{gtw: gateway.NewUserGateway(db)}
}

func (s *UserHTTPService) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := s.gtw.GetUsers()
	if users == nil || len(users) == 0 {
		users = []*model.User{}
	}
	web.Success(&users, http.StatusOK).Send(w)
}
