// Package main is our entry point into the application
package main

import (
	"log"
	"net/http"

	"bitbucket.org/evokesolutions/sm-rest-server/common"
	"bitbucket.org/evokesolutions/sm-rest-server/routers"
	"github.com/urfave/negroni"
)

// Runs start-up functions/methods and starts
// our app listening for incoming connections
func main() {
	common.StartUp()

	r := routers.InitRouters()

	n := negroni.Classic()
	n.UseHandler(r)

	log.Println("Listening ...")
	http.ListenAndServe(":8080", n)
}
