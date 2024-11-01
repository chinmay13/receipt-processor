package dto

import (
	"errors"
	"receipt-processing/models"
	"strconv"
)

// This method maps Item DTOs to Item model. It also converts the price from string to float.
func ItemDTOToItemMapper(itemDTO ItemDTO) (models.Item, error){
	var item models.Item
	price, err := strconv.ParseFloat(itemDTO.Price, 64)
    if err != nil {
        return item, errors.New("price invalid for item: " + itemDTO.ShortDescription + ". price must be a valid float in string format (e.g., '35.99')")
    }
	item.Price = price
	item.ShortDescription = itemDTO.ShortDescription
	return item, nil
}


// This method maps ReceiptDTO to Receipt model. It converts total value from string to float.
func ReceiptDTOToReceiptMapper(receiptDTO ReceiptDTO) (models.Receipt, map[string]string){
	var receipt models.Receipt
	errorMap := make(map[string]string)
	total, err := strconv.ParseFloat(receiptDTO.Total, 64)
    if err != nil {
        errorMap["Total"] = err.Error()
		return receipt, errorMap
    }
	var items []models.Item

	for _, itemDTO := range receiptDTO.Items {
		item, err := ItemDTOToItemMapper(itemDTO)
		if err != nil{
			errorMap["Items"] = err.Error()
			return receipt, errorMap
		}
		items = append(items, item)
	}
	receipt.Retailer = receiptDTO.Retailer
	receipt.PurchaseDate = receiptDTO.PurchaseDate
	receipt.PurchaseTime = receiptDTO.PurchaseTime
	receipt.Total = total
	receipt.Items = items
	return receipt, nil
}