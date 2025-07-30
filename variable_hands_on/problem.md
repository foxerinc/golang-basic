# Variable Hands On — Shopee Annual Gross Profit

## Objectives
- Use variables and operators correctly.
- Choose appropriate data types for currency and counts.
- Compute annual gross profit from 10 products.

## Problem
Hitung **gross profit tahunan** Shopee dari 10 produk berikut.

| Product | Price (Rp) | Total Sold | Discount (%) |
|---|---:|---:|---:|
| A | 100.000 | 200 | 0 |
| B | 67.000 | 12 | 20 |
| C | 56.000 | 80 | 0 |
| D | 1.000 | 1350 | 0 |
| E | 20.000 | 1 | 0 |
| F | 38.455 | 7 | 15 |
| G | 76.000 | 5.644 | 0 |
| H | 530.120 | 30 | 10 |
| I | 143.000 | 54 | 0 |
| J | 16.000 | 109 | 0 |

**Rumus:**  
\[ \text{Gross Profit} = \sum ((\text{price} \times \text{qty}) - (\text{price} \times \text{discount} \times \text{qty})) \]
dengan `discount` dalam persen \((0..100)\).

> **Catatan pembulatan**: Sesuai permintaan, gunakan **floor** untuk pembagian persen. Artinya, diskon dihitung dengan `(net*discount)/100` tanpa penambahan offset.

## Asumsi Data
- Gunakan **int64** untuk rupiah dan kuantitas. Hindari `float64` agar tidak ada error pembulatan.
- `discount` berupa persen bulat `0..100`.
- Semua nilai non-negatif.

## Pendekatan (ringkas)
1. Untuk tiap produk, hitung `net = price * qty`.
2. Hitung `disc = (net * discount) / 100` (**floor**).
3. Akumulasi `gross += net - disc`.
4. Cetak total dalam rupiah.