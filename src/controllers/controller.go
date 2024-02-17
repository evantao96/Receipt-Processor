package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "src/models"
)

// map of all receipt IDs to JSON receipt objects
var receipts = make(map[string]models.Receipt)

// POST /receipts/process
// postReceipts takes in a JSON receipt and responds with a JSON object containing an id for the receipt
func ReceiptID(c *gin.Context) {

    // The new receipt
    var newReceipt models.Receipt

    // Call BindJSON to bind the received JSON to newReceipt
    if err := c.BindJSON(&newReceipt); err != nil {
        return
    }

    // A new ID is generated
    id := uuid.New()

    // Add the new receipt ID and JSON object to the receipts map
    receipts[id.String()] = newReceipt

    // Returns a JSON object with the ID
    c.JSON(http.StatusOK, gin.H{"id": id.String()})
}

// GET /receipts/{id}/points
// getPoints responds with a JSON object containing the number of points awarded
func GetPoints(c *gin.Context) {

    // Retrieves the ID from the URL
    id := c.Param("id")

    // Looks up the receipt in the receipts map
    newReceipt, ok := receipts[id]

    // if the ID does not exist in the receipts map, return
    if !ok {
        return 
    }

    // The number of points awarded to newReceipt
    points := 0
    points += models.GetAlphaPoints(newReceipt)
    points += models.GetCentsPoints(newReceipt)
    points += models.GetNumItemsPoints(newReceipt)
    points += models.GetDayPoints(newReceipt)
    points += models.GetTimePoints(newReceipt)
    points += models.GetItemsPoints(newReceipt)

    // Returns an object specifying the points awarded
    c.JSON(http.StatusOK, gin.H{"points": points})
}