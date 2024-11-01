package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-processing/controllers"
	"receipt-processing/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/receipts/process", controllers.ProcessReceipt)
	router.GET("/receipts/:id/points", controllers.GetPoints)
	return router
}

func TestProcessReceipt_Success(t *testing.T) {
	router := setupRouter()

	receiptJSON := `{
        "retailer": "Target",
        "purchaseDate": "2022-01-01",
        "purchaseTime": "13:01",
        "items": [
            {"shortDescription": "Mountain Dew 12PK", "price": "6.49"},
			{"shortDescription": "Emils Cheese Pizza", "price": "12.25"}
        ],
        "total": "35.35"
    }`
	
	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var response models.ProcessReceiptResponse
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.ID)
}

func TestGetPoints_Success(t *testing.T) {
	router := setupRouter()
	receiptJSON := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "Klarbrunn 12-PK 12 FL OZ",
			"price": "12.00"
			}
		],
		"total": "35.35"
		}`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	var processResponse models.ProcessReceiptResponse
	_ = json.Unmarshal(res.Body.Bytes(), &processResponse)

	getReq, _ := http.NewRequest(http.MethodGet, "/receipts/"+processResponse.ID+"/points", nil)
	getRes := httptest.NewRecorder()
	router.ServeHTTP(getRes, getReq)

	assert.Equal(t, http.StatusOK, getRes.Code)
	var pointsResponse models.PointsResponse
	err := json.Unmarshal(getRes.Body.Bytes(), &pointsResponse)
	assert.Nil(t, err)
	assert.NotNil(t, pointsResponse.Points)
	assert.Equal(t, pointsResponse.Points, 28)
}

func TestProcessReceipt_InvalidFields(t *testing.T) {
	router := setupRouter()

	receiptJSON := `{
		"retailer": "Target",
		"purchaseDate": "invalid_date",
		"purchaseTime": "invalid_time",
		"items": [
			{"shortDescription": "Mountain Dew 12PK", "price": "invalid_price"},
			{"shortDescription": "Emils Cheese Pizza", "price": "12.25"}
		],
		"total": "invalid_total"
	}`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)

	var errorResponse map[string]map[string]string
	err := json.Unmarshal(res.Body.Bytes(), &errorResponse)
	assert.Nil(t, err)
	assert.Contains(t, errorResponse["Error in receipt JSON validation "], "Total")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "]["Total"], "Total must be a valid float value represented as a string")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "], "PurchaseDate")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "]["PurchaseDate"], "PurchaseDate must be in the format YYYY-MM-DD")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "], "PurchaseTime")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "]["PurchaseTime"], "PurchaseTime must be in the format HH:MM")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "], "Price")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "]["Price"], "Price must be a valid float value represented as a string")
}

func TestProcessReceipt_MissingFields(t *testing.T) {
	router := setupRouter()

	receiptJSON := `{
		
	}`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)

	var errorResponse map[string]map[string]string
	err := json.Unmarshal(res.Body.Bytes(), &errorResponse)
	assert.Nil(t, err)
	assert.Contains(t, errorResponse["Error in receipt JSON validation "], "Retailer")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "]["Retailer"], "Retailer is required and cannot be empty")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "], "PurchaseDate")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "]["PurchaseDate"], "PurchaseDate is required and cannot be empty")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "], "PurchaseTime")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "]["PurchaseTime"], "PurchaseTime is required and cannot be empty")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "], "Items")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "]["Items"], "At least one item is required in the receipt")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "], "Total")
	assert.Contains(t, errorResponse["Error in receipt JSON validation "]["Total"], "Total is required and cannot be empty")
}

func TestGetPoints_NonExistentID(t *testing.T) {
	router := setupRouter()

	getReq, _ := http.NewRequest(http.MethodGet, "/receipts/invalidID/points", nil)
	getRes := httptest.NewRecorder()
	router.ServeHTTP(getRes, getReq)

	assert.Equal(t, http.StatusNotFound, getRes.Code)

	var errorResponse map[string]string
	err := json.Unmarshal(getRes.Body.Bytes(), &errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, "Receipt ID not found", errorResponse["error"])
}