package models

import (
	testify "github.com/stretchr/testify/assert"
	"testing"
)

var basicItem = Item{
	ShortDescription: "Company",
	Price:            "15.00",
}

func TestShortDescriptionValidation(t *testing.T) {
	assert := testify.New(t)

	testItem := basicItem
	testItem.ShortDescription = "  Company-haha"
	err := testItem.Validate()
	assert.NoError(err)

	testItem = basicItem
	testItem.ShortDescription = "Company & haha"
	err = testItem.Validate()
	assert.Error(err)

	testItem = basicItem
	testItem.ShortDescription = "Company ! haha"
	err = testItem.Validate()
	assert.Error(err)
}

func TestPriceValidation(t *testing.T) {
	assert := testify.New(t)

	testItem := basicItem
	testItem.Price = "12.34"
	err := testItem.Validate()
	assert.NoError(err)

	testItem = basicItem
	testItem.Price = "12.341"
	err = testItem.Validate()
	assert.Error(err)

	testItem = basicItem
	testItem.Price = "12:34"
	err = testItem.Validate()
	assert.Error(err)

	testItem = basicItem
	testItem.Price = "12.010"
	err = testItem.Validate()
	assert.Error(err)
}
