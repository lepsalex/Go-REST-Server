// Package models defines all resources and ther DB operations.
package models

import (
	"bitbucket.org/evokesolutions/sm-rest-server/common"
	"gopkg.in/mgo.v2/bson"
)

type Expense struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	ExpenseType string        `bson:"expenseType" json:"expenseType"`
	Amount      float64       `bson:"amount" json:"amount"`
}

var Expenses = new(expenses)

type expenses struct{}

func (expenses) FindAll() ([]*Expense, error) {
	var ex []*Expense
	return ex, common.DB.Expenses.Find(nil).All(&ex)
}

func (expenses) FindOne(id string) (*Expense, error) {
	var e *Expense
	return e, common.DB.Expenses.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&e)
}

func (expenses) Create(expenseType string, amount float64) (*Expense, error) {
	e := &Expense{
		ID:          bson.NewObjectId(),
		ExpenseType: expenseType,
		Amount:      amount,
	}
	if err := common.DB.Expenses.Insert(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (expenses) Update(id, expenseType string, amount float64) error {
	if err := common.DB.Expenses.UpdateId(bson.ObjectIdHex(id),
		bson.M{"$set": bson.M{
			"expenseType": expenseType,
			"amount":      amount,
		}}); err != nil {
		return err
	}
	return nil
}

func (expenses) DeleteByID(id string) error {
	if err := common.DB.Expenses.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		return err
	}
	return nil
}
