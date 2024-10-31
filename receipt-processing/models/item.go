package models

import(
	"strconv"
	"encoding/json"
)

type Item struct{
	ShortDescription	string	`json:"shortDescription"`
	Price				float64	`json:"price"`
}

func (i *Item) UnmarshalJSON(data []byte) error {
    var temp struct {
        ShortDescription string `json:"shortDescription"`
        Price            string `json:"price"` // JSON input as string
    }
    if err := json.Unmarshal(data, &temp); err != nil {
        return err
    }

    // Parse the string price into float64
    price, err := strconv.ParseFloat(temp.Price, 64)
    if err != nil {
        return err
    }

    i.ShortDescription = temp.ShortDescription
    i.Price = price
    return nil
}