package services

import (
	"math"
	"receipt-processing/models"
	"receipt-processing/utils"
	"strings"
	"time"
)

// maintaining in memory map to keep track of receipts and points
var receipts = make(map[string]int)

func ProcessReceipt(receipt models.Receipt) string {
	points := calculatePoints(receipt)
	id := utils.GenerateID()
	receipts[id] = points
	return id
}

func GetPoints(id string) (int, bool) {
	points, exists := receipts[id]
	return points, exists
}

func calculatePoints(receipt models.Receipt) int {
	points := 0
	points += len(getAlphanumericChars(receipt.Retailer))

	// +50 points if total is round dollar
	if receipt.Total == float64(int(receipt.Total)){
		points += 50
	}

	// +25 points if total is multiple of 0.25
	if int(receipt.Total*100)%25 == 0 {
		points += 25
	}

	// +5 for every 2 items on the recipt
	points += 5 * (len(receipt.Items) / 2)

	// For every item, if desc %3 == 0: +ceil(itemPrice*0.2)  
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}
	
	// For every odd day purchase, +6
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 == 1 {
		points += 6
	}
	
	// For every purchase between 14-16, points +10
	// Design decision, including 14:00
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}
	
	return points
}

func getAlphanumericChars(str string) string {
    var result string
    for _, ch := range str {
        if ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ('0' <= ch && ch <= '9') {
            result += string(ch)
        }
    }
    return result
}