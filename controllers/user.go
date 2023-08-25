package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shashwat5522/golan_mongodb/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/Shashwat5522/golan_mongodb/models"
// 	"github.com/julienschmidt/httprouter"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"gopkg.in/mgo.v2/bson"
// )

// type UserController struct {
// 	Client *mongo.Client
// }

// func NewUserController(s *mongo.Client) *UserController {
// 	return &UserController{s}
// }

// func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// 	id := p.ByName("id")

// 	oid := bson.ObjectIdHex(id)
// 	u := models.User{}

// 	if err := uc.Client.Database("mongo-golang").Collection("users").FindId(oid).One(&u); err != nil {
// 		w.WriteHeader(404)
// 		return
// 	}
// 	uj, err := json.Marshal(u)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "%s\n", uj)

// }
var user models.User

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	user.Id = primitive.NewObjectID()
	err := user.CreateUser()
	// uc.Client.Database("mongo-golang").Collection("users").Insert(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "object inserted sucessfully!!")
}

// func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	id := p.ByName("id")
// 	if !bson.IsObjectIdHex(id) {
// 		w.WriteHeader(404)
// 		return
// 	}
// 	oid := bson.ObjectIdHex(id)
// 	if err := uc.Client.Database("mongo-golang").Collection("users").RemoveId(oid); err != nil {
// 		w.WriteHeader(404)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprint(w, "Deleted user", oid, "\n")
// }

func GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	Users, err := user.GetAllUsers()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Users)
	response, err := json.Marshal(Users)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(response)

}

func GetUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Println(id)
	User, err := user.GetUserById(id)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := json.Marshal(User)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(resp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	user, err := user.GetUserById(id)
	if err != nil {
		log.Fatal(err)
	}

	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	
	user.Name = newUser.Name
	user.Gender = newUser.Gender
	user.Age = newUser.Age
	fmt.Println(user.Id)
	resperr := user.UpdateUser()
	if resperr != nil {
		log.Fatal(resperr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("data update successfully"))

}

func DeleteUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	userId:=p.ByName("id")
	
	user,NotFoundError:=user.GetUserById(userId)
	if NotFoundError!=nil{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found!!!"))
	}
	err:=user.DeleteUser()
	if err!=nil{
		log.Fatal(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("data deleted Successfully"))
}

func GlobalSearch(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	searchWord:=r.URL.Query()
	fmt.Println(searchWord.Get("search"))
	resp,err:=user.Search(searchWord.Get("search"))
	if err!=nil{
		log.Fatal(err)
	}
	response,jerr:=json.Marshal(resp)
	if jerr!=nil{
		log.Fatal(jerr)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	
}