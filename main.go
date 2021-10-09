package main

import (
	"net/http"

	"github.com/Bhavyansh-1/Golang.git/controllers" //My own package
	"github.com/julienschmidt/httprouter"           //Router Package
	"gopkg.in/mgo.v2"                               //For mongo connection the driver and this package is from the same site from where mongo driver was given
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.POST("/users", uc.CreateUser)             //Route for creating a user
	r.GET("/users/:id", uc.GetUser)             //Route for geting the user from id
	r.POST("/posts", uc.CreatePost)             //Route for creating posts
	r.GET("/posts/:id", uc.GetPost)             //Route for getting posts from post id
	r.GET("/post/users/:id", uc.GetPostOfUsers) //Route for getting posts from user id

	http.ListenAndServe("localhost:9000", r) //My development Port
}
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017") //port at which my mongo db server runs locally
	if err != nil {
		panic(err)
	}
	return s
}
