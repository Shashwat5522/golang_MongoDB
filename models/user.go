package models

import (
	"context"
	"fmt"
	"log"

	"github.com/Shashwat5522/golan_mongodb/initializers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
}

func (user *User) CreateUser() error {
	fmt.Println(user)
	orgCollection := initializers.Mgr.Connection.Database("goguru123").Collection("collectiongoguru")

	result, err := orgCollection.InsertOne(context.TODO(), user)
	fmt.Println(result.InsertedID)
	fmt.Println(result)
	return err
}

func (user *User) GetAllUsers() ([]User, error) {
	var users []User
	fmt.Println("get all users called")
	orgCollection := initializers.Mgr.Connection.Database("goguru123").Collection("collectiongoguru")

	curr, err := orgCollection.Find(context.TODO(), bson.M{})
	fmt.Println(curr.Err(), err)

	for curr.Next(context.TODO()) {
		fmt.Println("inside loop")
		var user User
		decodeerr := curr.Decode(&user)
		fmt.Println(user, "hello")
		if decodeerr != nil {
			log.Fatal(decodeerr)
		}
		users = append(users, user)

	}
	if cerr := curr.Err(); cerr != nil {
		return nil, cerr
	}
	defer curr.Close(context.TODO())

	return users, err

}

func (user User) GetUserById(id string) (User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	orgCollection := initializers.Mgr.Connection.Database("goguru123").Collection("collectiongoguru")

	conerr := orgCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	return user, conerr

}

func (user *User) UpdateUser() error {
	fmt.Println("inside model func")
	fmt.Println(user)
	orgCollection := initializers.Mgr.Connection.Database("goguru123").Collection("collectiongoguru")
	
	filter := bson.D{{"_id", user.Id}}
	update := bson.D{{"$set", user}}

	updated, err := orgCollection.UpdateOne(context.TODO(), filter, update)
	fmt.Println(updated)
	return err

}

func(user *User)DeleteUser()error{
	orgCollection:=initializers.Mgr.Connection.Database("goguru123").Collection("collectiongoguru")

		_,err:=orgCollection.DeleteOne(context.TODO(),bson.D{{"_id",user.Id}})
		return err

}

func(user *User)Search(search string)([]User,error){
	fmt.Println(search)
	orgCollection:=initializers.Mgr.Connection.Database("goguru123").Collection("collectiongoguru")

	query:=bson.M{"$or":[]interface{}{
		bson.M{"name":bson.M{"$regex":search}},
		bson.M{"gender":bson.M{"$regex":search}},
	}}

	curr,err:=orgCollection.Find(context.TODO(),query)
	if err!=nil{
		return nil,err
	}

	var users []User
	if scanErr:=curr.All(context.TODO(),&users);scanErr!=nil{
		log.Fatal(scanErr)
	}
	return users,nil

}