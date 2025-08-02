# Grocery Shopping — Choki's Inventory Helper

## Objectives

- Use `map` to manage data.
- Use `struct` to model real-world entities.
- Use pointers to update data in place.
- Use method receivers on structs.

## Story

Choki has just opened a new grocery store in the heart of the Enchanted Forest.  
The shelves are full of magical snacks and forest-friendly food.  
To keep everything under control, Choki needs a small Go program to help manage inventory.

Your job is to build a single item model plus simple operations to add stock and sell items.

## Requirements

### 1. Define a struct

Create a struct called `GroceryItem` with these fields:

- `Name` (string)
- `Quantity` (int)
- `Price` (float64)

```go
type GroceryItem struct {
    Name     string
    Quantity int
    Price    float64
}
