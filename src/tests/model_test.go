package models

import (
	"testing"
	"src/models"
)

type modelTest struct {
	arg string
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

func TestNumItems(t *testing.T){

	items := []models.Item{
		{"Pepsi - 12-oz", "1.25"},
	}

	got := models.GetNumItemsPoints(items)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestDay(t *testing.T){

	purchaseDate := "2022-01-02"

	got := models.GetDayPoints(purchaseDate)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestTime(t *testing.T){

	purchaseTime := "13:13"

	got := models.GetTimePoints(purchaseTime)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestDescription(t *testing.T){

	items := []models.Item{
		{"Pepsi - 12-oz", "1.25"},
	}

	got := models.GetDescriptionPoints(items)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}