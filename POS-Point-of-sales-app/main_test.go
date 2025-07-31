package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var purchaseItems = []*PurchasedItem{
	{
		Barcode:  "00001",
		Name:     "Coca-cola",
		Subtotal: 6000,
		Price:    3000,
		Quantity: 2,
	},
	{
		Barcode:  "00002",
		Name:     "Sprite",
		Subtotal: 2500,
		Price:    2500,
		Quantity: 1,
	},
	{
		Barcode:  "00003",
		Name:     "Fanta",
		Subtotal: 2500,
		Price:    2500,
		Quantity: 1,
	},
}

var purchaseSummary = &PurchaseSummary{
	PurchasedItems: purchaseItems,
	TotalPrice:     11000,
}

var emptyPurchaseSummary = &PurchaseSummary{
	PurchasedItems: make([]*PurchasedItem, 0),
	TotalPrice:     0,
}

func TestCalculatePurchase(t *testing.T) {
	t.Run("should return correct object with purchased items group by barcode and totalPrice when cart contain barcodes", func(t *testing.T) {
		barcodes := []string{
			"00001",
			"00001",
			"00002",
			"00003",
		}

		result := CalculatePurchase(barcodes)

		assert.EqualValues(t, purchaseSummary, result)
	})

	t.Run("should return empty purchased items and totalPrice 0 when cart is empty", func(t *testing.T) {
		var barcodes []string

		purchaseSummary := CalculatePurchase(barcodes)

		assert.EqualValues(t, emptyPurchaseSummary, purchaseSummary)
	})
}

func TestGenerateReceiptText(t *testing.T) {
	t.Run("should return receipt text with correct format", func(t *testing.T) {
		expectedResult := "\n" +
			"Coca-cola 3000 x 2 = 6000\n" +
			"Sprite 2500 x 1 = 2500\n" +
			"Fanta 2500 x 1 = 2500\n" +
			"--------------------\n" +
			"Total = 11000"

		result := GenerateReceiptText(purchaseSummary)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("should return empty receipt when purchase item is empty and total price is 0", func(t *testing.T) {
		expectedResult := "\n" +
			"--------------------\n" +
			"Total = 0"

		result := GenerateReceiptText(emptyPurchaseSummary)

		assert.Equal(t, expectedResult, result)
	})
}
