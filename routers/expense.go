package routers

import (
	"bitbucket.org/evokesolutions/sm-rest-server/controllers"
	"github.com/gorilla/mux"
)

func SetExpenseRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/expenses", controllers.Expense.Create).Methods("POST")
	router.HandleFunc("/expenses", controllers.Expense.Get).Methods("GET")
	router.HandleFunc("/expenses/{id}", controllers.Expense.Show).Methods("GET")
	router.HandleFunc("/expenses/{id}", controllers.Expense.Update).Methods("PUT")
	router.HandleFunc("/expenses/{id}", controllers.Expense.Delete).Methods("DELETE")
	return router
}
