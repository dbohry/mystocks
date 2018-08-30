package services

import(
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dbohry/mystocks/configs"
	"github.com/dbohry/mystocks/models"
	"github.com/dbohry/mystocks/tools"
)

const cTransactions = "transactions"

//GetTransactions get all transactions
//
func GetTransactions(idUser string) []models.Transaction {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(cTransactions)

	result := make([]models.Transaction, 0, 10)
	fmt.Printf("\nSearching Transactions by User [%s]", idUser)
	err = c.Find(bson.M{}).All(&result)
	tools.ValidateFatal(err)

	return result
}

//GetTransactionByID get transaction filtered by id
//
func GetTransactionByID(idUser string, id string) models.Transaction {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(cTransactions)

	result := models.Transaction{}
	fmt.Printf("\nSearching Transaction [%s] of User [%s]", id, idUser)
	err = c.Find(bson.M{"id": id}).One(&result)
	tools.ValidateFatal(err)

	return result
}

//SaveTransaction save a new transactions
//
func SaveTransaction(user string, t models.Transaction) models.Transaction {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(cTransactions)

	t.User = getUser(user)
	t.Stock = getStock(t.Stock.ID)

	println("Saving new Transaction")
	err = c.Insert(t)
	tools.ValidateFatal(err)

	return t
}

//DeleteTransaction remove a transactions by ID
//
func DeleteTransaction(idUser string, id string) bool {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(cTransactions)

	println("Removing Transaction by ID: " + id)
	err = c.Remove(bson.M{"id": id})
	return tools.ValidateExecution(err)
}

func getUser(id string) models.User {
	return GetUserById(id)
}

func getStock(id string) models.Stock {
	return GetStockByID(id)
}