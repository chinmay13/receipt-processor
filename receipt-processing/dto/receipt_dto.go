package dto

// ReceiptDTO represents the receipt details, including retailer information, 
// purchase date, time, items, and total amount. This struct is used to interact
// with the outside world
type ReceiptDTO struct {
	Retailer     string    `json:"retailer" validate:"required"`
	Total        string    `json:"total" validate:"required,floatString"`
	PurchaseDate string    `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string    `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []ItemDTO `json:"items" validate:"required,dive,required"`
}