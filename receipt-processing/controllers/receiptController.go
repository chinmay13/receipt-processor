package controllers

import (
	"fmt"
	"net/http"
	"receipt-processing/dto"
	"receipt-processing/models"
	"receipt-processing/services"

	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Initializing validator to check if input jsons for apis are valid
var validate = validator.New()

func init() {
	// Registering a custom validation for checking if string is a valid float value
	validate.RegisterValidation("floatString", func(fl validator.FieldLevel) bool {
		floatRegex := regexp.MustCompile(`^[0-9]*\.?[0-9]+$`)
		return floatRegex.MatchString(fl.Field().String())
	})
}

// ProcessReceipt takes a JSON request containing receipt information, validates it, 
// and processes it to generate a unique receipt ID, which is then returned.
func ProcessReceipt(c *gin.Context){
	var receiptdto dto.ReceiptDTO
	if err := c.ShouldBindJSON(&receiptdto); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(receiptdto); err != nil {
		errors := formatValidationErrors(err.(validator.ValidationErrors))
		c.JSON(http.StatusBadRequest, gin.H{"Error in receipt JSON validation ":errors})
		return
	}

	receipt, errors := dto.ReceiptDTOToReceiptMapper(receiptdto)

	if errors != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error in receipt JSON validation ":errors})
		return
	}
	
	id := services.ProcessReceipt(receipt)
	c.JSON(http.StatusOK, models.ProcessReceiptResponse{ID: id})
}

// GetPoints retrieves the reward points for a given receipt ID.
// It checks if the receipt exists in the system and returns the points if found.
// If the receipt ID does not exist, it responds with a 404 Not Found error.
func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := services.GetPoints(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error":"Receipt ID not found"})
		return
	}
	c.JSON(http.StatusOK, models.PointsResponse{Points:points})
}


// formatValidationErrors generates readable error messages from validation errors.
// It customizes messages for specific fields like "PurchaseDate" and "Total" to
// provide more helpful information based on validation tags (e.g., "required", "datetime").
func formatValidationErrors(errs validator.ValidationErrors) map[string]string {
    errors := make(map[string]string)
    for _, err := range errs {
        switch err.Field() {
        case "PurchaseDate":
            if err.Tag() == "datetime" {
                errors["PurchaseDate"] = "PurchaseDate must be in the format YYYY-MM-DD"
            } else if  err.Tag() == "required" {
				errors["PurchaseDate"] = "PurchaseDate is required and cannot be empty"
			} else {
                errors["PurchaseDate"] = fmt.Sprintf("Invalid value for PurchaseDate: %v", err.Value())
            }
        case "PurchaseTime":
            if err.Tag() == "datetime" {
                errors["PurchaseTime"] = "PurchaseTime must be in the format HH:MM"
            } else if  err.Tag() == "required" {
				errors["PurchaseTime"] = "PurchaseTime is required and cannot be empty"
			} else {
                errors["PurchaseTime"] = fmt.Sprintf("Invalid value for PurchaseTime: %v", err.Value())
            }
        case "Retailer":
            if err.Tag() == "required" {
                errors["Retailer"] = "Retailer is required and cannot be empty"
            }
        case "Items":
            if err.Tag() == "required" {
                errors["Items"] = "At least one item is required in the receipt"
            }
		case "Total":
			if err.Tag() == "floatString" {
				errors["Total"] = "Total must be a valid float value represented as a string"
			} else {
				errors["Total"] = "Total is required and cannot be empty"
			}
		case "Price":
			if err.Tag() == "floatString" {
				errors["Price"] = "Price must be a valid float value represented as a string"
			} else {
				errors["Price"] = "Price is required and cannot be empty"
			}
        default:
            errors[err.Field()] = fmt.Sprintf("Invalid value for %s", err.Field())
        }
    }
    return errors
}