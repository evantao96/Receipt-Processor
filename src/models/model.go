package models

import (
    "strings"
    "encoding/json"
    "time"
)

// item represents data about a receipt item
type Item struct {
    // validates that the description is present
    ShortDescription    string       `json:"shortDescription" binding:"required"`
    // validates that the price is present and a positive number
    Price               json.Number  `json:"price" binding:"required,numeric,ne=0,excludes=-"`
}

// receipt represents data about a purchase receipt
type Receipt struct {
    // validates that the retailer is present
    Retailer        string      `json:"retailer" binding:"required"`
    // validates that the purchase date is in the appropriate date format
    PurchaseDate    string      `json:"purchaseDate" binding:"required,datetime=2006-01-02"`
    // validates that the purchase time is in the appropriate time format
    PurchaseTime    string      `json:"purchaseTime" binding:"required,datetime=15:04"`
    // validates that the items array is present, validates each item in the array
    Items           []Item      `json:"items" binding:"required,dive"`
    // validates that the total is present and a positive number
    Total           json.Number `json:"total" binding:"required,numeric,ne=0,excludes=-"`
}

/*  The Receipt model contains functions which calculate the number of points, 
 *  based on various attributes of the receipt. 
 */

// Returns the number of points earned based on the alphanumeric characters in the receipt
func GetAlphaPoints(myRetailer string) int {

    // alphanumeric characters 
    alphanumeric := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    // the number of alphanumeric characters in the retailer name
    count := 0

    for _, c := range myRetailer {
        if strings.Index(alphanumeric, string(c)) >= 0 {
            count++
        }
    }

    // One point for every alphanumeric character in the retailer name
    return count
}

// Returns 50 points if the total is a round dollar amount with no cents
func GetRoundTotalPoints(myTotal json.Number) int {

    // The number of cents in the total
    fTotal,_ := myTotal.Float64()
    cents := int(fTotal * 100)

    if cents % 100 == 0 {
        return 50 // 50 points if the total is a round dollar amount with no cents
    } else {
        return 0
    }
}

// Returns 25 points if the total is a multiple of 0.25
func GetMultiple25Points(myTotal json.Number) int {

    // The number of cents in the total
    fTotal,_ := myTotal.Float64()
    cents := int(fTotal * 100)

    if cents % 25 == 0 {  
        return 25 // 25 points if the total is a multiple of 25
    } else {
        return 0
    }
}

// Returns the number of points earned based on the number of items in the receipt
func GetNumItemsPoints(myItems []Item) int {

    // The number of items on the receipt
    numItems := len(myItems)

    // 5 points for every two items on the receipt
    return 5 * (numItems / 2)
}

// Returns the number of points earned based on the purchase date of the receipt
func GetDayPoints(myPurchaseDate string) int {

    // The day of the purchase date
    parsedDate,_ := time.Parse("2006-01-02", myPurchaseDate)
    day := parsedDate.Day()

    // 6 points if the day in the purchase date is odd
    if day % 2 == 1 {
        return 6
    } else {
        return 0
    }
}

// Returns the number of points earned based on the purchase time of the receipt
func GetTimePoints(myPurchaseTime string) int {

    // The hour and minutes of the purchase time 
    parsedTime,_ := time.Parse("15:04", myPurchaseTime)
    hour := parsedTime.Hour()
    minutes := parsedTime.Minute()

    // 10 points if the time of purchase is after 2:00pm and before 4:00pm
    if (hour == 14 && minutes > 0) || hour == 15 {
        return 10  
    } else {
        return 0
    }
}

// Returns the number of points earned based on the description of the items in the receipt
func GetDescriptionPoints(myItems []Item) int {

    sum := 0

    // Iterate through items on the receipt
    for _, c := range myItems {
        desc := c.ShortDescription
        price,_ := c.Price.Float64()

        // If the trimmed length of the item description is a multiple of 3
        if len(strings.TrimSpace(desc)) % 3 == 0 {

            // Multiply the price by 0.2
            price *= 0.2

            // Round up to the nearest integer
            roundedPrice := int(price + 1)
            sum += roundedPrice
        }
    }

    // The result is the number of points earned
    return sum 
}

