package main

import "testing"

func TestNewGroceryItem(t *testing.T) {
	item := NewGroceryItem("Acorn Crunchies", 50, 0.50)

	if item.Name != "Acorn Crunchies" {
		t.Errorf("Name = %s, want %s", item.Name, "Acorn Crunchies")
	}

	if item.Quantity != 50 {
		t.Errorf("Quantity = %d, want %d", item.Quantity, 50)
	}

	if item.Price != 0.50 {
		t.Errorf("Price = %f, want %f", item.Price, 0.50)
	}
}

func TestAddStock_IncreaseQuantity(t *testing.T) {
	item := NewGroceryItem("Acorn Crunchies", 50, 0.50)

	item.AddStock(20)

	if item.Quantity != 70 {
		t.Errorf("Quantity after AddStock = %d, want %d", item.Quantity, 70)
	}
}

func TestAddStock_IgnoreNonPositiveAmount(t *testing.T) {
	item := NewGroceryItem("Acorn Crunchies", 50, 0.50)

	item.AddStock(0)
	item.AddStock(-10)

	if item.Quantity != 50 {
		t.Errorf("Quantity should stay 50, got %d", item.Quantity)
	}
}

func TestSellItem_NotEnoughStock(t *testing.T) {
	item := NewGroceryItem("Acorn Crunchies", 50, 0.50)

	item.SellItem(100)

	if item.Quantity != 50 {
		t.Errorf("Quantity should stay 50 when stock is not enough, got %d", item.Quantity)
	}
}

func TestSellItem_EnoughStock(t *testing.T) {
	item := NewGroceryItem("Acorn Crunchies", 50, 0.50)

	item.SellItem(30)

	if item.Quantity != 20 {
		t.Errorf("Quantity after selling 30 = %d, want %d", item.Quantity, 20)
	}
}

func TestSellItem_IgnoreNonPositiveAmount(t *testing.T) {
	item := NewGroceryItem("Acorn Crunchies", 50, 0.50)

	item.SellItem(0)
	item.SellItem(-5)

	if item.Quantity != 50 {
		t.Errorf("Quantity should stay 50 when amount <= 0, got %d", item.Quantity)
	}
}
