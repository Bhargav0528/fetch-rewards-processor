package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func calculatePoints(receipt Receipt) int {
	points := 0
	// Points for Retailer name
	for _, char := range receipt.Retailer {
		if isAlphaNumeric(char) {
			points += 1
		}
	}

	fmt.Printf("Added %d points for RetailerName :: %s \n", points, receipt.Retailer)

	//Points for Total Field :: Whole Number
	if isInteger(receipt.Total) {
		points += 50
		fmt.Printf("Added %d points for Total Field (Whole Number) :: %s \n", 50, receipt.Total)
	}

	//Points for Total Field :: Multiple of 0.25
	totalFloatValue, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && math.Mod(totalFloatValue, 0.25) == 0 {
		points += 25
		fmt.Printf("Added %d points for Total Field (Multiple of 0.25) :: %s \n", 25, receipt.Total)
	}

	//Points for number of items : 5 for every 2 items
	if receipt.Items != nil {
		pointForNumItems := int(len(receipt.Items)/2) * 5
		points += pointForNumItems
		fmt.Printf("Added %d points for No. of Items (5 for every 2 items) :: %d items \n", pointForNumItems, len(receipt.Items))
	}

	// Points for length of item description : if multiple of 3, then corresponding item price * 0.2 rounded up
	if receipt.Items != nil {
		for _, item := range receipt.Items {
			trimmedDescription := strings.TrimSpace(item.ShortDescription)
			if len(trimmedDescription)%3 == 0 {
				itemPrice, err := strconv.ParseFloat(item.Price, 64)
				if err == nil {
					pointForCurItemDescription := int(math.Ceil(itemPrice * 0.2))
					points += pointForCurItemDescription
					fmt.Printf("Added %d points for Item Description :: %s with Price :: %s \n", pointForCurItemDescription, item.ShortDescription, item.Price)

				}
			}
		}
	}

	// Points for Purchase date if day is odd : 6 points
	purchaseDate, err := time.Parse(time.DateOnly, receipt.PurchaseDate)
	if err != nil {
		fmt.Println("Error is parsing the purchaseDate. Please check the format")
	} else {
		if purchaseDate.Day()%2 != 0 {
			points += 6
			fmt.Printf("Added %d points for PurchaseDate (if day is odd) :: %s \n", 6, receipt.PurchaseDate)
		}
	}

	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		fmt.Println("Error is parsing the purchaseTime. Please check the format")
	} else {
		// Does not include 2pm and 4pm in the window
		if (purchaseTime.Hour() > 14 || (purchaseTime.Hour() == 14 && purchaseTime.Minute() > 0)) && purchaseTime.Hour() < 16 {
			points += 10
			fmt.Printf("Added %d points for PurchaseTime (if time is > 2pm and < 4pm) :: %s \n", 10, receipt.PurchaseTime)
		}
	}

	return points
}
