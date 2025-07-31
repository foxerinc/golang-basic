package main

import (
	"fmt"
	"maps"
)

type Product struct {
	Barcode string
	Name    string
	Price   float32
}

type PurchasedItem struct {
	Barcode  string
	Name     string
	Subtotal float32
	Price    float32
	Quantity int
}

type PurchaseSummary struct {
	PurchasedItems []*PurchasedItem
	TotalPrice     float32
}

func GetProducts() []*Product {
	return []*Product{
		{
			Barcode: "00001",
			Name:    "Coca-cola",
			Price:   3000,
		},
		{
			Barcode: "00002",
			Name:    "Sprite",
			Price:   2500,
		},
		{
			Barcode: "00003",
			Name:    "Fanta",
			Price:   2500,
		},
		{
			Barcode: "00004",
			Name:    "Instant Noodles",
			Price:   3500,
		},
		{
			Barcode: "00005",
			Name:    "Coffee",
			Price:   5000,
		},
	}
}

var purchaseSum = PurchaseSummary{}

func CalculatePurchase(barcodes []string) *PurchaseSummary {

	pi := []*PurchasedItem{}
	if len(barcodes) < 1 {
		purchaseSum.PurchasedItems = pi
		purchaseSum.TotalPrice = 0
		return &purchaseSum
	}

	getProduct := GetProducts()

	subTotal := 0

	countItem := make(map[string]*PurchasedItem)

	for i := range barcodes {
		switch {
		case barcodes[i] == "00001":
			val, ok := countItem["00001"]
			if ok {
				val.Quantity += 1
				val.Subtotal += getProduct[0].Price
			} else {
				countItem["00001"] = &PurchasedItem{Barcode: getProduct[0].Barcode, Name: getProduct[0].Name, Subtotal: getProduct[0].Price, Price: getProduct[0].Price, Quantity: 1}
			}

			subTotal += int(getProduct[0].Price)

		case barcodes[i] == "00002":

			val, ok := countItem["00002"]
			if ok {
				val.Quantity += 1
				val.Subtotal += getProduct[1].Price
			} else {
				countItem["00002"] = &PurchasedItem{Barcode: getProduct[1].Barcode, Name: getProduct[1].Name, Subtotal: getProduct[1].Price, Price: getProduct[1].Price, Quantity: 1}
			}

			subTotal += int(getProduct[1].Price)

		case barcodes[i] == "00003":

			val, ok := countItem["00003"]
			if ok {
				val.Quantity += 1
				val.Subtotal += getProduct[2].Price
			} else {
				countItem["00003"] = &PurchasedItem{Barcode: getProduct[2].Barcode, Name: getProduct[2].Name, Subtotal: getProduct[2].Price, Price: getProduct[2].Price, Quantity: 1}
			}

			subTotal += int(getProduct[2].Price)

		case barcodes[i] == "00004":

			val, ok := countItem["00004"]
			if ok {
				val.Quantity += 1
				val.Subtotal += getProduct[3].Price
			} else {
				countItem["00004"] = &PurchasedItem{Barcode: getProduct[3].Barcode, Name: getProduct[3].Name, Subtotal: getProduct[3].Price, Price: getProduct[3].Price, Quantity: 1}
			}

			subTotal += int(getProduct[3].Price)
		case barcodes[i] == "00005":

			val, ok := countItem["00005"]
			if ok {
				val.Quantity += 1
				val.Subtotal += getProduct[4].Price
			} else {
				countItem["00005"] = &PurchasedItem{Barcode: getProduct[4].Barcode, Name: getProduct[4].Name, Subtotal: getProduct[4].Price, Price: getProduct[4].Price, Quantity: 1}
			}

			subTotal += int(getProduct[4].Price)
		}
	}

	for v := range maps.Values(countItem) {
		pi = append(pi, v)
	}

	purchaseSum.PurchasedItems = pi
	purchaseSum.TotalPrice = float32(subTotal)

	println(purchaseSum.PurchasedItems, purchaseSum.TotalPrice)
	return &purchaseSum

}

func GenerateReceiptText(purchaseSummary *PurchaseSummary) string {
	nota := ""
	if len(purchaseSummary.PurchasedItems) < 1 {
		nota += "\n--------------------\n" + "Total = 0"
		return nota
	}

	nota += "\n"
	for i := range purchaseSummary.PurchasedItems {
		nota += fmt.Sprintf("%s %.0f x %d = %.0f\n", purchaseSummary.PurchasedItems[i].Name, purchaseSummary.PurchasedItems[i].Price, purchaseSummary.PurchasedItems[i].Quantity, purchaseSummary.PurchasedItems[i].Subtotal)
	}

	nota += fmt.Sprintf("--------------------\nTotal = %.0f", purchaseSummary.TotalPrice)

	return nota

}

func main() {
	barcodes := []string{
		"00001",
		"00001",
		"00002",
		"00003",
	}
	CalculatePurchase(barcodes)
	GenerateReceiptText(&purchaseSum)
}
