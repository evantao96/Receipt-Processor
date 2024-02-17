package main

import (
    "github.com/gin-gonic/gin"
    "src/controllers"
)

func main() {
    r := gin.Default()
    r.POST("/receipts/process", controllers.ReceiptID)
    r.GET("/receipts/:id/points", controllers.GetPoints)
    r.Run(":8080")
}
