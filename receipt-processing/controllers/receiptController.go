package controllers

import (
	"net/http"
	"receipt-processing/models"
	"receipt-processing/services"

	"github.com/gin-gonic/gin"
)

func ProcessReceipt(c *gin.Context){
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := services.ProcessReceipt(receipt)
	c.JSON(http.StatusOK, models.ProcessReceiptResponse{ID: id})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := services.GetPoints(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error":"Receipt ID not found"})
		return
	}
	c.JSON(http.StatusOK, models.PointsResponse{Points:points})
}