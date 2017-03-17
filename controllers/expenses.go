// Package controllers defines all the callbacks for our Enpoints
// and their various HTTP methods (GET, POST, PUT, DELETE)
package controllers

import (
	"github.com/lepsalex/go-rest-server/common"
	"github.com/lepsalex/go-rest-server/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var Expense = new(expensesController)

type expensesController struct{}

func (ec *expensesController) Create(w http.ResponseWriter, r *http.Request) {
	var e models.Expense

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	expense, err := models.Expenses.Create(e.ExpenseType, e.Amount)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(expense)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusCreated)
}

func (ec *expensesController) Get(w http.ResponseWriter, r *http.Request) {
	expenses, err := models.Expenses.FindAll()
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(expenses)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

func (ec *expensesController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	expense, err := models.Expenses.FindOne(id)
	if err != nil {
		common.JsonError(w, err, http.StatusNotFound)
		return
	}

	res, err := json.Marshal(expense)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

func (ec *expensesController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var e models.Expense

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	if err := models.Expenses.Update(id, e.ExpenseType, e.Amount); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusNoContent)
}

func (ec *expensesController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := models.Expenses.DeleteByID(id); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusNoContent)
}
