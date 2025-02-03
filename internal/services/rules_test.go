package services_test

import (
	"testing"

	"github.com/AzlinII/receipt-processor-challenge/internal/model"
	"github.com/AzlinII/receipt-processor-challenge/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestRetailerNameRule(t *testing.T) {
	// given
	rule := services.RetailerNameRule()

	// when
	actual := rule(model.Receipt{
		Retailer: "M&M Corner Market",
	})

	// then
	assert.Equal(t, 14, actual)
}

func TestReceiptTotalRule(t *testing.T) {

	// given
	rule := services.ReceiptTotalRule()

	t.Run("fails parsing returns 0", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			Total: "fail",
		})

		// then
		assert.Equal(t, 0, actual)
	})

	t.Run("gives 75 points when no cents and multiple of 0.25", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			Total: "100.00",
		})

		// then
		assert.Equal(t, 75, actual)
	})

	t.Run("gives 25 if only cent is divisible by 0.25", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			Total: "0.75",
		})

		// then
		assert.Equal(t, 25, actual)
	})

	t.Run("gives 0 if has cents and not divisible by 25", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			Total: "100.42",
		})

		// then
		assert.Equal(t, 0, actual)
	})
}

func TestItemRule(t *testing.T) {
	// given
	rule := services.ItemsRule()

	t.Run("no points for single item and invalid item", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			Items: []model.Item{
				{ShortDescription: "fail", Price: "0"},
			},
		})

		// then
		assert.Equal(t, 0, actual)
	})

	t.Run("5 points given for items length", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			Items: []model.Item{
				{ShortDescription: "fail", Price: "0"},
				{ShortDescription: "fail", Price: "0"},
			},
		})

		// then
		assert.Equal(t, 5, actual)
	})

	t.Run("Points given only based off item only", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			Items: []model.Item{
				{ShortDescription: "Banana", Price: "6.00"},
			},
		})

		// then
		assert.Equal(t, 2, actual)
	})

	t.Run("Points given for items length and each item", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			Items: []model.Item{
				{ShortDescription: "Banana", Price: "6.00"},
				{ShortDescription: "Chocolate", Price: "3.00"},
			},
		})

		// then
		assert.Equal(t, 8, actual)
	})
}

func TestPurchaseDateRule(t *testing.T) {
	// given
	rule := services.PurchaseDateRule()

	t.Run("0 points for fail parsing", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			PurchaseDate: "02-03-2025",
		})

		// then
		assert.Equal(t, 0, actual)
	})

	t.Run("6 points for odd days", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			PurchaseDate: "2025-02-03",
		})

		// then
		assert.Equal(t, 6, actual)
	})

	t.Run("no points for even days", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			PurchaseDate: "2025-02-04",
		})

		// then
		assert.Equal(t, 0, actual)
	})
}

func TestPurchaseTimeRule(t *testing.T) {
	// given
	rule := services.PurchaseTimeRule()

	t.Run("0 points for fail parsing", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			PurchaseDate: "25:00",
		})

		// then
		assert.Equal(t, 0, actual)
	})

	t.Run("no points outside for 14:00 - 16:00", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			PurchaseTime: "11:00",
		})

		// then
		assert.Equal(t, 0, actual)
	})

	t.Run("no points at 14:00", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			PurchaseTime: "14:00",
		})

		// then
		assert.Equal(t, 0, actual)
	})

	t.Run("no points at 16:00", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			PurchaseTime: "16:00",
		})

		// then
		assert.Equal(t, 0, actual)
	})

	t.Run("10 points between 14:00 - 16:00", func(t *testing.T) {
		// when
		actual := rule(model.Receipt{
			PurchaseTime: "14:23",
		})

		// then
		assert.Equal(t, 10, actual)
	})
}
