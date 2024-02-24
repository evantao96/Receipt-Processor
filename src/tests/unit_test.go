package tests

import (
	"github.com/stretchr/testify/assert"
	"src/models"
	"testing"
)

// Struct for expected points based on retailer, purchase date, purchase time and total
type modelTest struct {
	arg string
	expected int
}

// Struct for expected points based on array of items
type itemTest struct {
	arg []models.Item
	expected int
}

// Tests for the number of points earned based on the alphanumeric characters in the receipt
var alphaTests = []modelTest{
				{"Target", 6},
				{"Walgreens", 9},
				{"M&M Corner Market", 14},
				{"Grove 34", 7},
}

// Tests for the number of points earned if the total is a round dollar amount with no cents
var roundTotalTests = []modelTest{
				{"1.25", 0},
				{"2.50", 0},
				{"35.35", 0},
				{"9.00", 50},
}

// Tests for the number of points earned if the total is a multiple of 0.25
var multiple25Tests = []modelTest{
				{"1.25", 25},
				{"2.65", 0},
				{"0.25", 25},
				{"9.00", 25},
				{"8.75", 25},
}

// Tests for the number of points earned based on the purchase date of the receipt
var dayTests = []modelTest{
				{"2022-01-02", 0},
				{"2022-01-01", 6},
				{"2022-03-20", 0},
				{"2022-03-17", 6},
}

// Tests for the number of points earned based on the purchase time of the receipt
var timeTests = []modelTest{
				{"13:13", 0},
				{"08:13", 0},
				{"14:01", 10},
				{"15:59", 10},
				{"16:00", 0},
}

// Tests for the number of points earned based on the number of items in the receipt
var numItemsTests = []itemTest{
	itemTest{[]models.Item{{"Pepsi - 12-oz", "1.25"}}, 0},
	itemTest{[]models.Item{{"Pepsi - 12-oz", "1.25"}, 
						   {"Dasani", "1.40"}}, 5},
	itemTest{[]models.Item{{"Mountain Dew 12PK", "6.49"}, 
						   {"Emils Cheese Pizza", "12.25"},
						   {"Knorr Creamy Chicken", "1.26"},
						   {"Doritos Nacho Cheese", "3.35"},
						   {"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},}, 10},
	itemTest{[]models.Item{{"Gatorade", "2.25"}, 
						   {"Gatorade", "2.25"}, 
						   {"Gatorade", "2.25"}, 
						   {"Gatorade", "2.25"}, 
						   {"Gatorade", "2.25"}, }, 10},
}

// Tests for the number of points earned based on the description of the items in the receipt
var descriptionTests = []itemTest{
	itemTest{[]models.Item{{"Pepsi - 12-oz", "1.25"}}, 0},
	itemTest{[]models.Item{{"Pepsi - 12-oz", "1.25"}, 
						   {"Dasani", "1.40"}}, 1},
	itemTest{[]models.Item{{"Mountain Dew 12PK", "6.49"}, 
						   {"Emils Cheese Pizza", "12.25"},
						   {"Knorr Creamy Chicken", "1.26"},
						   {"Doritos Nacho Cheese", "3.35"},
						   {"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},}, 6},
	itemTest{[]models.Item{{"Gatorade", "2.25"}, 
						   {"Gatorade", "2.25"}, 
						   {"Gatorade", "2.25"}, 
						   {"Gatorade", "2.25"}, 
						   {"Gatorade", "2.25"}, }, 0},
}

// Iterates through tests and prints the results
func TestAlpha(t *testing.T){
	for _, test := range alphaTests{
		output := models.GetAlphaPoints(test.arg)
		assert.Equal(t, test.expected, output)
	}
}

func TestRoundTotal(t *testing.T){
	for _, test := range roundTotalTests{
		output := models.GetRoundTotalPoints(test.arg)
		assert.Equal(t, test.expected, output)
	}
}

func TestMultiple25(t *testing.T){
	for _, test := range multiple25Tests{
		output := models.GetMultiple25Points(test.arg)
		assert.Equal(t, test.expected, output)
	}
}

func TestDay(t *testing.T){
	for _, test := range dayTests{
		output := models.GetDayPoints(test.arg)
		assert.Equal(t, test.expected, output)
	}
}

func TestTime(t *testing.T){
	for _, test := range timeTests{
		output := models.GetTimePoints(test.arg)
		assert.Equal(t, test.expected, output)
	}
}

func TestNumItems(t *testing.T){
	for _, test := range numItemsTests{
		output := models.GetNumItemsPoints(test.arg)
		assert.Equal(t, test.expected, output)
	}
}

func TestDescription(t *testing.T){
	for _, test := range descriptionTests{
		output := models.GetDescriptionPoints(test.arg)
		assert.Equal(t, test.expected, output)
	}
}