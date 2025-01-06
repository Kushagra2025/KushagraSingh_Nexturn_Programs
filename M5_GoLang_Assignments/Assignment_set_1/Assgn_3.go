package main

import (
	"errors"
	"fmt"
	"sort"
)

// Product represents a product in the inventory
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

// AddProduct adds a new product to the inventory
func AddProduct(id int, name string, priceInput string, stock int) error {
	// Convert price from string to float64
	price, err := stringToFloat64(priceInput)
	if err != nil {
		return err
	}

	// Ensure stock is non-negative
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	// Create new product and add to inventory
	product := Product{ID: id, Name: name, Price: price, Stock: stock}
	inventory = append(inventory, product)

	return nil
}

// stringToFloat64 converts a string input to float64
func stringToFloat64(str string) (float64, error) {
	var price float64
	_, err := fmt.Sscanf(str, "%f", &price)
	if err != nil {
		return 0, errors.New("invalid price format")
	}
	return price, nil
}

// UpdateStock updates the stock of an existing product
func UpdateStock(id int, stock int) error {
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	// Find the product by ID and update its stock
	for i := range inventory {
		if inventory[i].ID == id {
			inventory[i].Stock = stock
			return nil
		}
	}
	return fmt.Errorf("product with ID %d not found", id)
}

// SearchProduct searches for a product by ID or Name
func SearchProduct(id int, name string) (*Product, error) {
	for i := range inventory {
		if inventory[i].ID == id || inventory[i].Name == name {
			return &inventory[i], nil
		}
	}
	return nil, fmt.Errorf("product not found with ID: %d or Name: %s", id, name)
}

// DisplayInventory displays the inventory of all products
func DisplayInventory() {
	if len(inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}
	fmt.Println("ID\tName\t\tPrice\tStock")
	fmt.Println("---------------------------------------")
	for _, product := range inventory {
		fmt.Printf("%d\t%s\t%.2f\t%d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// SortInventoryByPrice sorts the inventory by price in ascending order
func SortInventoryByPrice() {
	sort.SliceStable(inventory, func(i, j int) bool {
		return inventory[i].Price < inventory[j].Price
	})
}

// SortInventoryByStock sorts the inventory by stock in ascending order
func SortInventoryByStock() {
	sort.SliceStable(inventory, func(i, j int) bool {
		return inventory[i].Stock < inventory[j].Stock
	})
}

func main_3() {
	// Adding products to inventory
	err := AddProduct(1, "Product1", "10.50", 100)
	if err != nil {
		fmt.Println("Error adding product:", err)
	}
	err = AddProduct(2, "Product2", "20.30", 50)
	if err != nil {
		fmt.Println("Error adding product:", err)
	}
	err = AddProduct(3, "Product3", "15.40", 70)
	if err != nil {
		fmt.Println("Error adding product:", err)
	}

	// Display all products
	fmt.Println("\nInventory before sorting:")
	DisplayInventory()

	// Sorting products by price
	SortInventoryByPrice()
	fmt.Println("\nInventory sorted by price:")
	DisplayInventory()

	// Sorting products by stock
	SortInventoryByStock()
	fmt.Println("\nInventory sorted by stock:")
	DisplayInventory()

	// Search for a product by ID
	product, err := SearchProduct(2, "")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nFound product by ID: %v\n", product)
	}

	// Search for a product by Name
	product, err = SearchProduct(0, "Product3")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nFound product by Name: %v\n", product)
	}

	// Update stock of a product
	err = UpdateStock(2, 60)
	if err != nil {
		fmt.Println("Error updating stock:", err)
	} else {
		fmt.Println("\nStock updated for Product ID 2")
		DisplayInventory()
	}
}
