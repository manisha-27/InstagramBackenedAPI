package main

import(
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2"
"net/http"

"github.com/manisha-27/mongo-golang/controllers"

//...................
// "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/mongo/options"
//..................
)

func main(){

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.POST("/users", uc.CreateUser)
	r.GET("/users/:id", uc.GetUser)
	r.POST("/posts", uc.CreatePost)
	r.GET("/posts/users/:id", uc.GetPost)
	// r.GET("/posts/users/:id", uc.List)
	r.DELETE("/users/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)

	// ......................
	
	//.......................
}

func getSession() *mgo.Session{
	
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s
}
