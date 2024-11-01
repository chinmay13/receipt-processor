package models

// Receipt represents the receipt details, including retailer information, 
// purchase date, time, items, and total amount.
type Receipt struct {
	Retailer     string  `json:"retailer"`
	Total        float64 `json:"total"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
}