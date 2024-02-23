package tests

import (
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
	modelTest{"Target", 6},
	modelTest{"Walgreens", 9},
	modelTest{"M&M Corner Market", 14},
	modelTest{"Grove 34", 7},
}

// Tests for the number of points earned if the total is a round dollar amount with no cents
var roundTotalTests = []modelTest{
	modelTest{"1.25", 0},
	modelTest{"2.50", 0},
	modelTest{"35.35", 0},
	modelTest{"9.00", 50},
}

// Tests for the number of points earned if the total is a multiple of 0.25
var multiple25Tests = []modelTest{
	modelTest{"1.25", 25},
	modelTest{"2.65", 0},
	modelTest{"0.25", 25},
	modelTest{"9.00", 25},
	modelTest{"8.75", 25},
}

// Tests for the number of points earned based on the purchase date of the receipt
var dayTests = []modelTest{
	modelTest{"2022-01-02", 0},
	modelTest{"2022-01-01", 6},
	modelTest{"2022-03-20", 0},
	modelTest{"2022-03-17", 6},
}

// Tests for the number of points earned based on the purchase time of the receipt
var timeTests = []modelTest{
	modelTest{"13:13", 0},
	modelTest{"08:13", 0},
	modelTest{"14:01", 10},
	modelTest{"15:59", 10},
	modelTest{"16:00", 0},
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
		if output := models.GetAlphaPoints(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestRoundTotal(t *testing.T){
	for _, test := range roundTotalTests{
		if output := models.GetRoundTotalPoints(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestMultiple25(t *testing.T){
	for _, test := range multiple25Tests{
		if output := models.GetMultiple25Points(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestDay(t *testing.T){
	for _, test := range dayTests{
		if output := models.GetDayPoints(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestTime(t *testing.T){
	for _, test := range timeTests{
		if output := models.GetTimePoints(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestNumItems(t *testing.T){
	for _, test := range numItemsTests{
		if output := models.GetNumItemsPoints(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestDescription(t *testing.T){
	for _, test := range descriptionTests{
		if output := models.GetDescriptionPoints(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}