package utils

import (
	"fetch-rewards/models"
	testify "github.com/stretchr/testify/assert"
	"testing"
)

// Test struct
var pc = PointsCalculator{}
var itemToEqual3PointsRoundUp = models.Item{
	ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
	Price:            "10.50",
}

var itemToEqualNoPoints = models.Item{
	ShortDescription: "Knorr Creamy Chicken",
	Price:            "8.25",
}

var testReceipt = models.Receipt{
	ID:           "",
	Retailer:     "  M&M Company",
	PurchaseDate: "2022-01-01",
	PurchaseTime: "13:01",
	Items:        nil,
	Total:        "35.00",
}

func TestCalculateAlphaNumeric(t *testing.T) {
	assert := testify.New(t)

	expected := 9
	retailer := "  M&M Company"
	points, err := pc.calculateAlphaNumeric(retailer)

	assert.NoError(err)
	assert.Equal(expected, points)
}

func TestCalculateRound(t *testing.T) {
	assert := testify.New(t)
	expectedSuccess := 50
	expectedFail := 0

	total := "35.00"
	points, err := pc.calculateRound(total)
	assert.NoError(err)
	assert.Equal(expectedSuccess, points)

	total = "21.50"
	points, err = pc.calculateRound(total)
	assert.NoError(err)
	assert.Equal(expectedFail, points)

	total = "21.12"
	points, err = pc.calculateRound(total)
	assert.NoError(err)
	assert.Equal(expectedFail, points)
}

func TestCalculateMultiple(t *testing.T) {
	assert := testify.New(t)

	expectedSuccess := 25
	expectedFail := 0

	total := "39.25"
	points, err := pc.calculateMultiple(total)
	assert.NoError(err)
	assert.Equal(expectedSuccess, points)

	total = "21.10"
	points, err = pc.calculateMultiple(total)
	assert.NoError(err)
	assert.Equal(expectedFail, points)
}

func TestCalculateDualItems(t *testing.T) {
	assert := testify.New(t)
	var items []models.Item

	items = append(items, itemToEqualNoPoints)
	expectedSuccess := 0
	points, err := pc.calculateDualItems(items)
	assert.NoError(err)
	assert.Equal(expectedSuccess, points)

	items = append(items, itemToEqualNoPoints)
	expectedSuccess = 5
	points, err = pc.calculateDualItems(items)
	assert.NoError(err)
	assert.Equal(expectedSuccess, points)

	items = append(items, itemToEqualNoPoints)
	expectedSuccess = 5
	points, err = pc.calculateDualItems(items)
	assert.NoError(err)
	assert.Equal(expectedSuccess, points)

	items = append(items, itemToEqualNoPoints)
	expectedSuccess = 10
	points, err = pc.calculateDualItems(items)
	assert.NoError(err)
	assert.Equal(expectedSuccess, points)
}

func TestCalculateDescription(t *testing.T) {
	assert := testify.New(t)
	expectedSuccess := 0
	var items []models.Item

	items = append(items, itemToEqualNoPoints)
	points, err := pc.calculateDescription(items)
	assert.NoError(err)
	assert.Equal(expectedSuccess, points)

	expectedSuccess = 3
	items = append(items, itemToEqual3PointsRoundUp)
	points, err = pc.calculateDescription(items)
	assert.NoError(err)
	assert.Equal(expectedSuccess, points)

	expectedSuccess = 6
	items = append(items, itemToEqual3PointsRoundUp)
	points, err = pc.calculateDescription(items)
	assert.NoError(err)
	assert.Equal(expectedSuccess, points)
}
