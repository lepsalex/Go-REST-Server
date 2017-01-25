package common

import (
	"gopkg.in/mgo.v2"
)

// DBConfig - connect config
type DBConfig struct {
	address string
	name    string
}

// DB - Pointer to mongo type
var DB *mongo

// Defines all our collections
type mongo struct {
	Expenses          *mgo.Collection
	ExpenseLineItems  *mgo.Collection
	ExpenseCategories *mgo.Collection
	Suppliers         *mgo.Collection
	Budgets           *mgo.Collection
	Horses            *mgo.Collection
	Media             *mgo.Collection
	Currencies        *mgo.Collection
}

// Start a MongoDB session
func connectDB(mongoConfig DBConfig) {
	session, err := mgo.Dial(mongoConfig.address)
	if err != nil {
		panic(err)
	}
	DB = &mongo{
		session.DB(mongoConfig.name).C("expenses"),
		session.DB(mongoConfig.name).C("expenseLineItems"),
		session.DB(mongoConfig.name).C("expenseCategories"),
		session.DB(mongoConfig.name).C("suppliers"),
		session.DB(mongoConfig.name).C("budgets"),
		session.DB(mongoConfig.name).C("horses"),
		session.DB(mongoConfig.name).C("media"),
		session.DB(mongoConfig.name).C("currencies"),
	}
}
