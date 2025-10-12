package models

import (
	"fmt"
	"sandwich-shop/utils"
	"time"
)

type Order struct {
	ID           int
	CustomerName string
	ItemsOrdered []MenuItem
	Quantity     int
	TotalPrice   float32
	TimeOfOrder  time.Time
}

// Pricer defines an interface for any item that has a price
type MenuItem interface {
	GetPrice() float32
	PrintData()
}

type Sandwich struct {
	Size        string
	Bread       string
	Meat        string
	Cheese      string
	Sauce       string
	Toppings    []string
	ExtraMeat   bool
	ExtraCheese bool
	Price       float32
}

func (s Sandwich) GetPrice() float32 {
	return s.Price
}

func (s Sandwich) PrintData() {
	fmt.Println("\n---SANDWICH---")
	fmt.Println("Size:", utils.CapitalizeFirstLetter(s.Size))
	fmt.Println("Bread:", utils.CapitalizeFirstLetter(s.Bread))
	fmt.Println("Meat:", utils.CapitalizeFirstLetter(s.Meat))
	if s.ExtraMeat {
		fmt.Println("Extra Meat!")
	}
	fmt.Println("Cheese:", utils.CapitalizeFirstLetter(s.Cheese))
	if s.ExtraCheese {
		fmt.Println("Extra Cheese!")
	}
	fmt.Println("Sauce:", utils.CapitalizeFirstLetter(s.Sauce))

	for _, topping := range s.Toppings {
		fmt.Println("Topping:", utils.CapitalizeFirstLetter(topping))
	}
	fmt.Printf("\nSandwich Price: $%.2f\n", s.Price)
}

type Chip struct {
	Type  string
	Size  string
	Price float32
}

func (c Chip) GetPrice() float32 {
	return c.Price
}

func (c Chip) PrintData() {
	fmt.Println("\n---CHIP---")
	fmt.Println("Chip Type:", c.Type)
	fmt.Println("Chip Size:", c.Size)
	fmt.Printf("\nChip Price: $%.2f\n", c.Price)
}

type Drink struct {
	Type  string
	Size  string
	Price float32
}

func (d Drink) GetPrice() float32 {
	return d.Price
}

func (d Drink) PrintData() {
	fmt.Println("\n---DRINK---")
	fmt.Println("Drink Type:", d.Type)
	fmt.Println("Drink Size:", d.Size)
	fmt.Printf("\nDrink Price: $%.2f\n", d.Price)
}
