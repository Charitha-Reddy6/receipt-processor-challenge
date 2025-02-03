package utils

import (
	"math"
	"receipt-processor/models"
	"regexp"
	"strconv"
	"strings"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// 1. One point for every alphanumeric character in the retailer name
	re := regexp.MustCompile("[a-zA-Z0-9]")
	points += len(re.FindAllString(receipt.Retailer, -1))

	// 2. 50 points if total is a whole number
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// 3. 25 points if total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 4. 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// 5. Points for item descriptions with length multiple of 3
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6. 6 points if the purchase date is odd
	dateParts := strings.Split(receipt.PurchaseDate, "-")
	day, _ := strconv.Atoi(dateParts[2])
	if day%2 != 0 {
		points += 6
	}

	// 7. 10 points if the purchase time is between 2:00pm and 4:00pm
	timeParts := strings.Split(receipt.PurchaseTime, ":")
	hour, _ := strconv.Atoi(timeParts[0])
	if hour >= 14 && hour < 16 {
		points += 10
	}

	return points
}
