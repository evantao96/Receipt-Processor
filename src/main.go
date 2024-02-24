package main

import (
    "github.com/gin-gonic/gin"
    "src/controllers"
    "src/models"
)

func main() {
    // "database" of all receipt IDs to JSON receipt objects
    var allReceipts = make(map[string]models.Receipt)

    r := gin.Default()
    h := controllers.ReceiptHandler{Receipts: allReceipts}
    r.POST("/receipts/process", h.ProcessReceipt)
    r.GET("/receipts/:id/points", h.GetPoints)
    r.Run(":8080")
}
