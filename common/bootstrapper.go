// Package common contains all universal code consumed
// by various other packages
package common

// StartUp - Bootstrapping calls here
func StartUp() {
	// Connection config
	mongoConfig := DBConfig{
		"mongodb://evoke-dev:esMockingJay2@ds157187.mlab.com:57187/sm-rest-server",
		"sm-rest-server",
	}
	// Connect to Mongo
	connectDB(mongoConfig)
}
