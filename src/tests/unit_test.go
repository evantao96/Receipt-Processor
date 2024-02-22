package tests

import (
	"testing"
	"src/models"
)

type modelTest struct {
	arg string
	expected int
}

type itemTest struct {
	arg []models.Item
	expected int
}

var alphaTests = []modelTest{
	modelTest{"Target", 6},
	modelTest{"Walgreens", 9},
	modelTest{"M&M Corner Market", 14},
	modelTest{"Grove 34", 7},
}

var roundTotalTests = []modelTest{
	modelTest{"1.25", 0},
	modelTest{"2.50", 0},
	modelTest{"35.35", 0},
	modelTest{"9.00", 50},
}

var multiple25Tests = []modelTest{
	modelTest{"1.25", 25},
	modelTest{"2.65", 0},
	modelTest{"0.25", 25},
	modelTest{"9.00", 25},
	modelTest{"8.75", 25},
}

var dayTests = []modelTest{
	modelTest{"2022-01-02", 0},
	modelTest{"2022-01-01", 6},
	modelTest{"2022-03-20", 0},
	modelTest{"2022-03-17", 6},
}

var timeTests = []modelTest{
	modelTest{"13:13", 0},
	modelTest{"08:13", 0},
	modelTest{"14:01", 10},
	modelTest{"15:59", 10},
	modelTest{"16:00", 0},
}

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