package services

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dbohry/mystocks/configs"
	"github.com/dbohry/mystocks/models"
	"github.com/dbohry/mystocks/tools"
)

const cUsers = "users"

//GetUsers get all users
//
func GetUsers() []models.User {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(cUsers)

	result := make([]models.User, 0, 10)
	fmt.Println("Searching for all Users")
	err = c.Find(bson.M{}).All(&result)
	tools.ValidateFatal(err)

	return result
}

//GetUserByName get user filtered by Id
//
func GetUserById(id string) models.User {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(cUsers)

	result := models.User{}
	fmt.Println("Searching User by Id: " + id)
	err = c.FindId(bson.ObjectIdHex(id)).One(&result)
	tools.ValidateFatal(err)

	return result
}

//SaveUser save a new user
//
func SaveUser(t models.User) models.User {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(cUsers)

	fmt.Println("Saving new User")
	err = c.Insert(t)
	tools.ValidateFatal(err)

	return t
}

//DeleteUser remove an user by ID
//
func DeleteUser(id string) bool {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(cUsers)

	println("Removing User by ID: " + id)
	err = c.Remove(bson.M{"id": id})
	return tools.ValidateExecution(err)
}
