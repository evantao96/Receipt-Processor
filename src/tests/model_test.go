package models

import (
	"testing"
	"src/models"
)

func TestAlpha(t *testing.T){

	retailer := "Target"

	got := models.GetAlphaPoints(retailer)
	want := 6

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRoundTotal(t *testing.T){

	total := `"1.00"`

	got := models.GetRoundTotalPoints(total)
	want := 50

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMultiple25(t *testing.T){

	total := `"1.00"`

	got := models.GetMultiple25Points(total)
	want := 25

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
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