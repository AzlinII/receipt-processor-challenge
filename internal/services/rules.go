package services

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/AzlinII/receipt-processor-challenge/internal/model"
)

type Rule func(model.Receipt) int

func RetailerNameRule() Rule {
	return func(receipt model.Receipt) int {
		regex := regexp.MustCompile(`[A-z0-9]`)
		// One point for every alphanumeric character in the retailer name.
		return len(regex.FindAllString(receipt.Retailer, -1))
	}
}

func ReceiptTotalRule() Rule {
	return func(receipt model.Receipt) int {
		points := 0
		value, err := strconv.ParseFloat(receipt.Total, 64)
		if err != nil {
			// If there's an error in conversion, return 0 points
			return 0
		}
		valueInCents := int(value * 100)

		//  50 points if the total is a round dollar amount with no cents
		if valueInCents%100 == 0 {
			points += 50
		}

		// 25 points if the total is a multiple of 25 (since multiplied by 100)
		if valueInCents%25 == 0 {
			points += 25
		}
		return points
	}
}

func ItemsRule() Rule {
	return func(receipt model.Receipt) int {
		items := receipt.Items
		// 5 points for every two items on the receipt
		points := 5 * (len(items) / 2)

		// If the trimmed length of the item description is a multiple of 3,
		// multiply the price by `0.2` and round up to the nearest integer.
		// The result is the number of points earned
		for _, item := range items {
			trimmedDesc := strings.TrimSpace(item.ShortDescription)
			if len(trimmedDesc)%3 == 0 {
				value, err := strconv.ParseFloat(item.Price, 64)
				if err != nil {
					// If there's an error in conversion, skip item
					continue
				}
				point := int(math.Ceil(value * 0.2))
				points += point
			}
		}
		return points
	}
}

func PurchaseDateRule() Rule {
	return func(receipt model.Receipt) int {
		// 6 points if the day in the purchase date is odd
		date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
		if err != nil {
			return 0
		}
		fmt.Println(date.Day())
		if date.Day()%2 != 0 {
			return 6
		}
		return 0
	}
}

func PurchaseTimeRule() Rule {
	return func(receipt model.Receipt) int {
		format := "15:04"
		// 10 points if the time of purchase is after 2:00pm and before 4:00pm
		curTime, err := time.Parse(format, receipt.PurchaseTime)
		if err != nil {
			return 0
		}
		twoPm, _ := time.Parse(format, "14:00")
		fourPm, _ := time.Parse(format, "16:00")

		if curTime.After(twoPm) && curTime.Before(fourPm) {
			return 10
		}
		return 0
	}
}
