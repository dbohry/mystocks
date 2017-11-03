package tests

import (
	"testing"
	"reflect"

	"github.com/dbohry/mystocks/models"
	"github.com/dbohry/mystocks/services"
)

func TestStockService(t *testing.T) {
	actualResult := services.GetStocks()
	expectedResult := []models.Stock{}
	if reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}