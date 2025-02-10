package models

import (
	testify "github.com/stretchr/testify/assert"
	"testing"
)

var basicReceipt = Receipt{
	ID:           "",
	Retailer:     "M&M Company",
	PurchaseDate: "2022-01-01",
	PurchaseTime: "13:00",
	Items:        nil,
	Total:        "35.00",
}

func TestRetailerValidation(t *testing.T) {
	assert := testify.New(t)

	testReceipt := basicReceipt
	testReceipt.Retailer = "   M&M Company"
	err := testReceipt.Validate()
	assert.NoError(err)

	testReceipt.Retailer = "M&!M Company"
	err = testReceipt.Validate()
	assert.Error(err)

	testReceipt.Retailer = "M&#M Company"
	err = testReceipt.Validate()
	assert.Error(err)

	testReceipt.Retailer = "M&$M Company"
	err = testReceipt.Validate()
	assert.Error(err)
}

func TestTotalValidation(t *testing.T) {
	assert := testify.New(t)

	testReceipt := basicReceipt
	testReceipt.Total = "15.00"
	err := testReceipt.Validate()
	assert.NoError(err)

	testReceipt.Total = "15:00"
	err = testReceipt.Validate()
	assert.Error(err)

	testReceipt.Total = "15.0035"
	err = testReceipt.Validate()
	assert.Error(err)

	testReceipt.Total = "a5.00"
	err = testReceipt.Validate()
	assert.Error(err)

	testReceipt.Total = "-5.00"
	err = testReceipt.Validate()
	assert.Error(err)
}

func TestPurchaseDateValidation(t *testing.T) {
	assert := testify.New(t)

	testReceipt := basicReceipt
	testReceipt.PurchaseDate = "2022-01-01"
	err := testReceipt.Validate()
	assert.NoError(err)

	testReceipt.PurchaseDate = "2022-21-01"
	err = testReceipt.Validate()
	assert.Error(err)

	testReceipt.PurchaseDate = "2022-02-32"
	err = testReceipt.Validate()
	assert.Error(err)

	testReceipt.PurchaseDate = "20221-02-32"
	err = testReceipt.Validate()
	assert.Error(err)
}

func TestPurchaseTimeValidation(t *testing.T) {
	assert := testify.New(t)

	testReceipt := basicReceipt
	testReceipt.PurchaseTime = "21:01"
	err := testReceipt.Validate()
	assert.NoError(err)

	testReceipt.PurchaseTime = "25:01"
	err = testReceipt.Validate()
	assert.Error(err)

	testReceipt.PurchaseTime = "25:60"
	err = testReceipt.Validate()
	assert.Error(err)
}
