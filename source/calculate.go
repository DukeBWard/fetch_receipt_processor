package main

// * One point for every alphanumeric character in the retailer name.
// * 50 points if the total is a round dollar amount with no cents.
// * 25 points if the total is a multiple of `0.25`.
// * 5 points for every two items on the receipt.
// * If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
// * 6 points if the day in the purchase date is odd.
// * 10 points if the time of purchase is after 2:00pm and before 4:00pm.

import (
	"math"
	"strconv"
	"strings"
	"time"
)

func calcPoints(receipt Receipt, user User) (int, error) {
	points := 0

	if user.ReceiptCount == 1 {
		points += 1000
	} else if user.ReceiptCount == 2 {
		points += 500
	} else if user.ReceiptCount == 3 {
		points += 250
	}

	// 1 point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	// total to float
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		println("Error parsing total:", err)
		return 0, err
	}

	// 50 points if total is a whole number
	if total == float64(int(total)) {
		points += 50
	}

	// 25 points of total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every 2 items
	points += (len(receipt.Items) / 2) * 5

	// if item desc length is multiple of 3
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				price = 0.0
			}
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if day is odd
	// go handles date and time formats weirdly
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		println("Error parsing date:", err)
		return 0, err
	}

	if date.Day()%2 == 1 {
		points += 6
	}

	// 10 points if time is between 2pm and 4pm
	time, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		println("Error parsing time:", err)
		return 0, err
	}

	if time.Hour() >= 14 && time.Hour() < 16 {
		points += 10
	}

	//println("Points:", points)
	return points, nil

}

func isAlphanumeric(char rune) bool {
	return (char >= '0' && char <= '9') || (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}
