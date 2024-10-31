package models

import (
	"encoding/json"
	"strconv"
)

type Receipt struct {
	Retailer     string  `json:"retailer"`
	Total        float64 `json:"total"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
}

func (r *Receipt) UnmarshalJSON(data []byte) error {
    var temp struct {
        Retailer     string `json:"retailer"`
        Total        string `json:"total"` // JSON input as string
        PurchaseDate string `json:"purchaseDate"`
        PurchaseTime string `json:"purchaseTime"`
        Items        []Item `json:"items"`
    }
    if err := json.Unmarshal(data, &temp); err != nil {
        return err
    }

    // Parse the string total into float64
    total, err := strconv.ParseFloat(temp.Total, 64)
    if err != nil {
        return err
    }

    r.Retailer = temp.Retailer
    r.Total = total
    r.PurchaseDate = temp.PurchaseDate
    r.PurchaseTime = temp.PurchaseTime
    r.Items = temp.Items
    return nil
}