package main

import (
	"math"
	"strconv"
)

func calcPoints(receipt Receipt) int {
	points := 0

	// 1 point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	// total to float
	total, err := strconv.ParseFloat(receipt.Total, 64)

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

}
