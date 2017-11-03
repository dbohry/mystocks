package services

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dbohry/mystocks/configs"
	"github.com/dbohry/mystocks/models"
	"github.com/dbohry/mystocks/tools"
)

const collection = "stocks"

//GetStocks get all stocks
//
func GetStocks() []models.Stock {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(collection)

	result := make([]models.Stock, 0, 10)
	err = c.Find(bson.M{}).All(&result)
	tools.ValidateFatal(err)

	return result
}

//GetStockByID get stock filtered by id
//
func GetStockByID(id string) models.Stock {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(collection)

	result := models.Stock{}
	err = c.Find(bson.M{"id": id}).One(&result)
	tools.ValidateFatal(err)

	return result
}

//CreateStock save a new stock
//
func CreateStock(s models.Stock) models.Stock {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(collection)

	err = c.Insert(s)
	tools.ValidateFatal(err)

	return s
}

//DeleteStock remove a stock by ID
//
func DeleteStock(id string) {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(collection)
	err = c.Remove(bson.M{"id": id})
	tools.ValidateFatal(err)
}