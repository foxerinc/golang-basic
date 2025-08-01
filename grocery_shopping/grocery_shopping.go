package main

import "fmt"

type GroceryItem struct {
	Name     string
	Quantity int
	Price    float64
}

func NewGroceryItem(name string, quantity int, price float64) *GroceryItem {
	return &GroceryItem{
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
}

func (g *GroceryItem) AddStock(amount int) {
	if amount <= 0 {
		return
	}

	g.Quantity += amount
}

func (g *GroceryItem) SellItem(amount int) {
	if amount <= 0 {
		return
	}

	if amount > g.Quantity {
		fmt.Printf("Oh no! Not enough %s for the hungry squirrels!\n", g.Name)
		return
	}

	g.Quantity -= amount
}

func main() {
	fmt.Println("Welcome to Choki's Grocery Store")

	inventory := make(map[string]*GroceryItem)

	item := NewGroceryItem("Acorn Crunchies", 50, 0.50)
	inventory[item.Name] = item

	fmt.Printf(
		"Item Details: Name: %s, Quantity: %d, Price: $%.2f\n",
		item.Name,
		item.Quantity,
		item.Price,
	)

	fmt.Println("A fresh batch of Acorn Crunchies has arrived! Adding 20 to stock...")
	item.AddStock(20)
	fmt.Printf("New Quantity: %d\n", item.Quantity)

	fmt.Println("Attempting to sell 100 Acorn Crunchies...")
	item.SellItem(100)

	fmt.Println("Selling 30 Acorn Crunchies...")
	item.SellItem(30)
	fmt.Printf("New Quantity: %d\n", item.Quantity)

	fmt.Println("Hooray! The forest creatures are happy and full!")

	_ = inventory
}
