package dto

// ItemDTO represents each item in the receipt with a short description and price.
// This struct is used to interact with outside world
type ItemDTO struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required,floatString"`
}
