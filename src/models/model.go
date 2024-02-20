package models

import (
    "strings"
    "strconv"
    "encoding/json"
)

// item represents data about a receipt item
type Item struct {
    ShortDescription    string  `json:"shortDescription" binding:"required"`
    Price               json.Number  `json:"price" binding:"required,numeric,excludesall=-."`
}

// receipt represents data about a purchase receipt
type Receipt struct {
    Retailer        string      `json:"retailer" binding:"required"`
    PurchaseDate    string      `json:"purchaseDate" binding:"required"`
    PurchaseTime    string      `json:"purchaseTime" binding:"required"`
    Items           []Item      `json:"items" binding:"required,dive"`
    Total           string      `json:"total" binding:"required"`
}

/*  The Receipt model contains functions which calculate the number of points, 
 *  based on various attributes of the receipt. 
 */

// Returns the number of points earned based on the alphanumeric characters in the receipt
func GetAlphaPoints(receipt Receipt) int {

    // alphanumeric characters 
    alphanumeric := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    // the number of alphanumeric characters in the retailer name
    count := 0

    for _, c := range receipt.Retailer {
        if strings.Index(alphanumeric, string(c)) >= 0 {
            count++
        }
    }

    // One point for every alphanumeric character in the retailer name
    return count
}

// Returns the number of points earned based on the cents in the receipt
func GetCentsPoints(receipt Receipt) int {

    // The number of cents in the total
    cents,_ := strconv.Atoi(receipt.Total[strings.Index(receipt.Total, ".")+1:])

    if cents == 0 {
        return 50 // 50 points if the total is a round dollar amount with no cents
    } else if cents % 25 == 0 {  
        return 25 // 25 points if the total is a multiple of 0.25
    } else {
        return 0
    }
}

// Returns the number of points earned based on the number of items in the receipt
func GetNumItemsPoints(receipt Receipt) int {

    // The number of items on the receipt
    numItems := len(receipt.Items)

    // 5 points for every two items on the receipt
    return 5 * (numItems / 2)
}

// Returns the number of points earned based on the purchase date of the receipt
func GetDayPoints(receipt Receipt) int {

    // The day of the purchase date
    day,_ := strconv.Atoi(receipt.PurchaseDate[len(receipt.PurchaseDate)-2:])

    // 6 points if the day in the purchase date is odd
    if day % 2 == 1 {
        return 6
    } else {
        return 0
    }
}

// Returns the number of points earned based on the purchase time of the receipt
func GetTimePoints(receipt Receipt) int {

    // The hour of the purchase time 
    hour,_ := strconv.Atoi(receipt.PurchaseTime[:2])
    // hour := receipt.PurchaseTime.Hour()

    // The minutes of the purchase time
    minutes,_ := strconv.Atoi(receipt.PurchaseTime[3:])
    // minutes := receipt.PurchaseTime.Minute()

    // 10 points if the time of purchase is after 2:00pm and before 4:00pm
    if (hour == 14 && minutes > 0) || hour == 15 {
        return 10  
    } else {
        return 0
    }
}

// Returns the number of points earned based on the items in the receipt
func GetItemsPoints(receipt Receipt) int {

    sum := 0

    // Iterate through items on the receipt
    for _, c := range receipt.Items {
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

