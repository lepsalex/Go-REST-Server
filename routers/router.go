// Package routers defines our main router and all
// routes and associated handler functions defined
// in the controllers package
package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetExpenseRouters(router)
	return router
}
