package controllers
    
import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "src/models"
    "src/dbs"
)

// POST /receipts/process
// Process takes in a JSON receipt and responds with a JSON object containing an id for the receipt
func ProcessReceipt(c *gin.Context) {

    // The new receipt
    var newReceipt models.Receipt

    // Call BindJSON to bind the received JSON to newReceipt
    if err := c.BindJSON(&newReceipt); err != nil {
        c.String(http.StatusBadRequest, "The receipt is invalid")
        return
    }

    // A new ID is generated
    id := uuid.New()

    // Add the new receipt ID and JSON object to the receipts map
    dbs.AllReceipts[id.String()] = newReceipt

    // Returns a JSON object with the ID
    c.JSON(http.StatusOK, gin.H{"id": id.String()})
}

// GET /receipts/{id}/points
// GetPoints responds with a JSON object containing the number of points awarded
func GetPoints(c *gin.Context) {

    // Retrieves the ID from the URL
    id := c.Param("id")

    // Looks up the receipt in the receipts map
    myReceipt, ok := dbs.AllReceipts[id]
    if !ok {
        c.String(http.StatusNotFound, `No receipt found for that id`)
        return
    }

    // The number of points awarded to newReceipt
    points := models.GetAlphaPoints(myReceipt.Retailer) + models.GetRoundTotalPoints(myReceipt.Total) + models.GetMultiple25Points(myReceipt.Total) + models.GetNumItemsPoints(myReceipt.Items) + models.GetDayPoints(myReceipt.PurchaseDate) + models.GetTimePoints(myReceipt.PurchaseTime) + models.GetDescriptionPoints(myReceipt.Items)

    // Returns an object specifying the points awarded
    c.JSON(http.StatusOK, gin.H{"points": points})
}