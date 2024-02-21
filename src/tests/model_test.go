package models

import (
	"testing"
	"src/models"
)

var r models.Receipt

func TestAlpha(t *testing.T){

	r.Retailer = "Target"

	got := models.GetAlphaPoints(r)
	want := 6

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRoundTotal(t *testing.T){

	r.Retailer = "1.00"

	got := models.GetRoundTotalPoints(r)
	want := 50

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMultiple25(t *testing.T){

	r.Retailer = "1.00"

	got := models.GetMultiple25Points(r)
	want := 25

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestNumItems(t *testing.T){

	r.Items = []models.Item{
		{"Pepsi - 12-oz", "1.25"},
	}

	got := models.GetNumItemsPoints(r)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestDay(t *testing.T){

	r.PurchaseDate = "2022-01-02"

	got := models.GetDayPoints(r)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestTime(t *testing.T){

	r.PurchaseTime = "13:13"

	got := models.GetTimePoints(r)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestDescription(t *testing.T){

	r.Items = []models.Item{
		{"Pepsi - 12-oz", "1.25"},
	}

	got := models.GetDescriptionPoints(r)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}