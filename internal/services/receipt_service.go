package services

import (
	"fmt"
	"regexp"
	"time"

	"github.com/AzlinII/receipt-processor-challenge/internal/customerror"
	"github.com/AzlinII/receipt-processor-challenge/internal/model"
)

type PointsDBAccessor interface {
	GetPoints(id string) (int, error)
	SavePoints(points int) string
}

type PointsService struct {
	pointsDBAccessor PointsDBAccessor
	rules            []Rule
}

func NewPointsService(pointsDBAccessor PointsDBAccessor) PointsService {
	rules := []Rule{
		RetailerNameRule(),
		ReceiptTotalRule(),
		ItemsRule(),
		PurchaseDateRule(),
		PurchaseTimeRule(),
	}
	return PointsService{
		pointsDBAccessor: pointsDBAccessor,
		rules:            rules,
	}
}

func (p PointsService) Process(receipt model.Receipt) (string, error) {
	if !isValidReceipt(receipt) {
		return "", customerror.NewInvalidReceiptError()
	}
	totalPoints := 0
	for _, rule := range p.rules {
		totalPoints += rule(receipt)
	}
	fmt.Println(totalPoints)
	return p.pointsDBAccessor.SavePoints(totalPoints), nil
}

func (p PointsService) GetPoints(id string) (int, error) {
	points, err := p.pointsDBAccessor.GetPoints(id)
	if err != nil {
		return 0, err
	}
	return points, nil
}

func isValidReceipt(receipt model.Receipt) bool {
	return isValidPattern(receipt.Retailer, "^[\\w\\s\\-&]+$") &&
		isValidPattern(receipt.Total, "^\\d+\\.\\d{2}$") &&
		isValidDateString(receipt.PurchaseDate) &&
		isValidTimeString(receipt.PurchaseTime) &&
		allItemsValid(receipt.Items)

}

func isValidPattern(text string, pattern string) bool {
	_, err := regexp.MatchString(pattern, text)
	return err == nil
}

func isValidDateString(text string) bool {
	_, err := time.Parse("2006-01-02", text)
	return err == nil
}

func isValidTimeString(text string) bool {
	_, err := time.Parse("15:04", text)
	return err == nil
}

func allItemsValid(items []model.Item) bool {
	if len(items) < 1 {
		return false
	}
	for _, item := range items {
		if !isValidPattern(item.Price, "^[\\w\\s\\-]+$") ||
			!isValidPattern(item.ShortDescription, "^\\d+\\.\\d{2}$") {
			return false
		}
	}
	return true
}
