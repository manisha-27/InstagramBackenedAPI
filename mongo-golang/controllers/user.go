package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/manisha-27/mongo-golang/models"
	"errors"
	"time"
	// "context"
	

)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

var id string
// var password string

func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application//json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
	fmt.Println(u.Id)
	fmt.Println(u.Name)
	fmt.Println(u.Email)
	fmt.Println(u.Password)
}

func (uc UserController) CreatePost (w http.ResponseWriter, r *http.Request, _ httprouter.Params){


	u := models.Post{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = id
	// u.Password= password

	if u.Id !=id {
        errors.New("Wrong Credentials")
    }

	time.Now()

	uc.session.DB("mongo-golang").C("posts").Insert(u)

	uj, err := json.Marshal(u)

	if err!= nil{
		fmt.Println(err)
	}


	w.Header().Set("Content-Type", "application//json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

	// fmt.Println(u.Id)
	// fmt.Println(u.Caption)
	// fmt.Println(u.ImageURL)
	// // fmt.Println(u.Password)
	// fmt.Println(time.Now())
}

func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	

	u := models.User{}
	
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u);err!=nil{
		w.WriteHeader(404)
		return
	}

	uj, err :=json.Marshal(u)
	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) GetPost (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	// GetUser()
	id1 := p.ByName(id)

	if !bson.IsObjectIdHex(id1){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id1)

	u := models.Post{}
	
	if err := uc.session.DB("mongo-golang").C("posts").FindId(oid).One(&u);err!=nil{
		w.WriteHeader(404)
		return
	}

	uj, err :=json.Marshal(u)
	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}



func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil{
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}