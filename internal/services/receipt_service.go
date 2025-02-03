package services

import (
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
}

func NewPointsService(pointsDBAccessor PointsDBAccessor) PointsService {
	return PointsService{
		pointsDBAccessor: pointsDBAccessor,
	}
}

func (p PointsService) Process(receipt model.Receipt) (string, error) {
	// TODO: Calculate points from receipt
	if !isValidReceipt(receipt) {
		return "", customerror.NewInvalidReceiptError()
	}
	totalPoints := 42
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
	_, err := time.Parse("01/02/2006", text)
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
