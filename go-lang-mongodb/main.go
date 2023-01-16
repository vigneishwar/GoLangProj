package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	// import controller
	"github.com/vigneishwar/go-lang-mongodb/controllers"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser) // create user
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r) // start the go server on port 8080

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017") // connect to mongodb on port 27017
	if err != nil {
		panic(err)
	}
	return s

}
