# Variable Hands-On — Shopee Annual Gross Profit

## Objectives
- Use variables and operators correctly.
- Choose appropriate data types for currency and counts.
- Compute annual gross profit from 10 products.

## Problem
Calculate **annual gross profit** for Shopee from the 10 products below.

| Product | Price (IDR) | Total Sold | Discount (%) |
|---|---:|---:|---:|
| A | 100,000 | 200 | 0 |
| B | 67,000 | 12 | 20 |
| C | 56,000 | 80 | 0 |
| D | 1,000 | 1,350 | 0 |
| E | 20,000 | 1 | 0 |
| F | 38,455 | 7 | 15 |
| G | 76,000 | 5,644 | 0 |
| H | 530,120 | 30 | 10 |
| I | 143,000 | 54 | 0 |
| J | 16,000 | 109 | 0 |

**Formula:**  
\[ \text{Gross Profit} = \sum \big((\text{price} \times \text{qty}) - (\text{price} \times \text{discount} \times \text{qty})\big) \]
with `discount` in percent \((0..100)\).

> **Rounding note:** Use **floor** for percent division. Compute the discount as `(net * discount) / 100` using integer division.

## Data Assumptions
- Use **int64** for rupiah and quantities. Avoid `float64` to prevent rounding errors.
- `discount` is an integer percent `0..100`.
- All values are non-negative.

## Approach (brief)
1. For each product, compute `net = price * qty`.
2. Compute `disc = (net * discount) / 100` (**floor**).
3. Accumulate `gross += net - disc`.
4. Print the total in rupiah.


## How to run
```
go run .
```