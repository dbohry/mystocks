package services

import(
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
	err = c.Find(bson.M{"id": id}).One(&result)
	tools.ValidateFatal(err)

	return result
}

//SaveTransaction save a new transactions
//
func SaveTransaction(t models.Transaction) models.Transaction {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(cTransactions)
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
	err = c.Remove(bson.M{"id": id})
	return tools.ValidateExecution(err)
}