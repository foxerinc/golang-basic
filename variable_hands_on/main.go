package main


import (
	"fmt"
)

type Product struct {
	Price    int64
	Qty      int64
	Discount int64
}

func gross(p Product) int64 {
	total := p.Price * p.Qty
	disc := (total * p.Discount) / 100
	return total - disc
}

func formatIDR(n int64) string {
	if n == 0 {
		return "Rp 0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	s := fmt.Sprintf("%d", n)
	out := make([]byte, 0, len(s)+len(s)/3)
	for i := 0; i < len(s); i++ {
		if i != 0 && (len(s)-i)%3 == 0 {
			out = append(out, '.')
		}
		out = append(out, s[i])
	}
	return "Rp " + sign + string(out)
}

func main() {
	products := []Product{
		{Price: 100_000, Qty: 200, Discount: 0},
		{Price: 67_000, Qty: 12, Discount: 20},
		{Price: 56_000, Qty: 80, Discount: 0},
		{Price: 1_000, Qty: 1350, Discount: 0},
		{Price: 20_000, Qty: 1, Discount: 0},
		{Price: 38_455, Qty: 7, Discount: 15},
		{Price: 76_000, Qty: 5644, Discount: 0},
		{Price: 530_120, Qty: 30, Discount: 10},
		{Price: 143_000, Qty: 54, Discount: 0},
		{Price: 16_000, Qty: 109, Discount: 0},
	}

	var total int64
	for _, p := range products {
		total += gross(p)
	}

	fmt.Println("Annual gross profit:", formatIDR(total)) // expected: Rp 479.445.247
}