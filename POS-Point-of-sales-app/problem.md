# POS – Point of Sales App

## Problem Overview

A cashier needs a simple Point of Sales (POS) application that can:

1. Take a list of product barcodes as input.
2. Group items by barcode.
3. Calculate quantity, subtotal per product, and grand total.
4. Generate a text receipt for printing.

This exercise is focused on basic Go concepts: `struct`, `for` loops, slices, and simple formatting.

## Product Catalog

The store has a fixed product list:

| Barcode | Product Name     | Price |
|--------:|------------------|------:|
| 00001   | Coca-cola        | 3000  |
| 00002   | Sprite           | 2500  |
| 00003   | Fanta            | 2500  |
| 00004   | Instant Noodles  | 3500  |
| 00005   | Coffee           | 5000  |

## Data Structures

The core types used:

```go
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
```

- `Product` represents an item in the store catalog.
- `PurchasedItem` represents a grouped purchase line (one product, many units).
- `PurchaseSummary` represents the complete checkout result.

## Required Functions

### 1. `CalculatePurchase(barcodes []string) *PurchaseSummary`

Responsibilities:

- Input: slice of barcodes (e.g. `[]string{"00001", "00001", "00002", "00003"}`).
- Group barcodes by product.
- For each product:
  - Count `Quantity`.
  - Compute `Subtotal = Price * Quantity`.
- Compute `TotalPrice` as the sum of all subtotals.
- Return a `*PurchaseSummary` containing:
  - `PurchasedItems` (one entry per distinct barcode).
  - `TotalPrice`.

Example (from unit test):

Input barcodes:

```go
[]string{
    "00001",
    "00001",
    "00002",
    "00003",
}
```

Expected `PurchaseSummary`:

```go
&PurchaseSummary{
    PurchasedItems: []*PurchasedItem{
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
    },
    TotalPrice: 11000,
}
```

If the cart is empty (`len(barcodes) == 0`), the function should return:

```go
&PurchaseSummary{
    PurchasedItems: []*PurchasedItem{},
    TotalPrice:     0,
}
```

### 2. `GenerateReceiptText(purchaseSummary *PurchaseSummary) string`

Responsibilities:

- Input: a `*PurchaseSummary` (usually output of `CalculatePurchase`).
- Output: a formatted receipt text.

Format used in tests:

```text
Coca-cola 3000 x 2 = 6000
Sprite 2500 x 1 = 2500
Fanta 2500 x 1 = 2500
--------------------
Total = 11000
```

The actual string (including leading newline) expected by the test:

```go
expectedResult := "
" +
    "Coca-cola 3000 x 2 = 6000
" +
    "Sprite 2500 x 1 = 2500
" +
    "Fanta 2500 x 1 = 2500
" +
    "--------------------
" +
    "Total = 11000"
```

For an empty purchase summary:

```go
&PurchaseSummary{
    PurchasedItems: []*PurchasedItem{},
    TotalPrice:     0,
}
```

Expected receipt:

```text
--------------------
Total = 0
```

Exact string in the test:

```go
expectedResult := "
" +
    "--------------------
" +
    "Total = 0"
```

## How to Run

Run the program:

```bash
go run .
```

Run the tests:

```bash
go test ./...
```

The tests validate:

- Grouping logic in `CalculatePurchase`.
- Correct quantities, subtotals, and total.
- Exact formatting of the receipt string in `GenerateReceiptText`.