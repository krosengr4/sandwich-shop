package ui

import (
	"fmt"

	"sandwich-shop/utils"
)

func HomeScreen() int {
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - Order\n2 - Admin\n0 - Exit")

	return utils.GetValidatedNumber("Enter option: ", 0, 2)
}

func OrderScreen() int {
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - Order Sandwich\n2 - Order Chips\n3 - Order Drink\n4 - Checkout\n0 - Cancel")

	return utils.GetValidatedNumber("Enter option: ", 0, 4)
}

// #region Sandwich UI
func SandwichSizes() int {
	fmt.Println("\nPlease choose your size")
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - Small (4\") ... $5.50\n2 - Medium (8\") ... $7.00\n3 - Large (12\") ... $8.50")

	return utils.GetValidatedNumber("Enter option: ", 1, 3)
}

func SandwichBreads() int {
	fmt.Println("\nPlease choose your bread")
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - White\n2 - Wheat\n3 - Sourdough\n4 - Lettuce Wrap")

	return utils.GetValidatedNumber("Enter option: ", 1, 4)
}

func SandwichMeats() int {
	fmt.Println("\nPlease choose your meat. Small: +$1.00 | Medium: +$2.00 | Large: +$3.00")
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - Turkey\n2 - Ham\n3 - Roast Beef\n4 - Steak\n0 - None")

	return utils.GetValidatedNumber("Enter option: ", 0, 4)
}

func ExtraMeatOption(size string) int {
	fmt.Println("\nWould you like extra meat?")
	fmt.Println("\n-----OPTIONS-----")

	var price float32

	switch size {
	case "small":
		price = .50
	case "medium":
		price = 1.00
	case "large":
		price = 1.50
	}

	fmt.Printf("1 - Yes(+$%.2f)\n2 - No\n", price)
	return utils.GetValidatedNumber("Enter option: ", 1, 2)
}

func SandwichCheese() int {
	fmt.Println("\nPlease choose your cheese. Small: +$0.75 | Medium: +$1.50 | Large: +$2.25")
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - American\n2 - Cheddar\n3 - Provolone\n4 - Swiss\n0 - None")

	return utils.GetValidatedNumber("Enter option: ", 0, 4)
}

func ExtraCheeseOption(size string) int {
	fmt.Println("\nWould you like extra cheese?")
	fmt.Println("\n-----OPTIONS-----")

	var price float32

	switch size {
	case "small":
		price = .30
	case "medium":
		price = .60
	case "large":
		price = .90
	}

	fmt.Printf("1 - Yes(+$%.2f)\n2 - No\n", price)
	return utils.GetValidatedNumber("Enter option: ", 1, 2)
}

func SandwichSauce() int {
	fmt.Println("\nPlease choose your sauce")
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - Mayo\n2 - Mustard\n3 - Ketchup\n0 - None")

	return utils.GetValidatedNumber("Enter option: ", 0, 3)
}

func SandwichTopping() int {
	fmt.Println("\nPlease choose a topping")
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - Onions\n2 - Lettuce\n3 - Peppers\n4 - Pickles\n5 - Tomatoes\n0 - None")

	return utils.GetValidatedNumber("Enter option: ", 0, 5)
}

// #endregion

// #region Chips UI
func ChipTypes() int {
	fmt.Println("\n-----CHIP OPTIONS-----")
	fmt.Println("1 - Doritos\n2 - Lays Original\n3 - Cheetos\n0 - Go Back")

	return utils.GetValidatedNumber("Enter option: ", 0, 3)
}

func ChipSizes() int {
	fmt.Println("\n-----CHIP SIZE OPTIONS-----")
	fmt.Println("1 - Small ($1.25)\n2 - Medium ($2.00)\n3 - Large ($2.70)")

	return utils.GetValidatedNumber("Enter option: ", 1, 3)
}

// #endregion

// #region Drink UI
func DrinkTypes() int {
	fmt.Println("\n-----DRINK OPTIONS-----")
	fmt.Println("1 - Soda\n2 - Lemonade\n3 - Milkshake\n0 - Go Back")

	return utils.GetValidatedNumber("Enter opiton: ", 0, 3)
}

func DrinkSizes() int {
	fmt.Println("\n-----DRINK SIZES-----")
	fmt.Println("1 - Small (12oz) ... $2.00\n2 - Medium (24oz) ... $2.50\n3 - Large (32ox) ... $3.00")

	return utils.GetValidatedNumber("Enter option: ", 1, 3)
}

// #endregion

func AdminScreen() int {
	fmt.Println("\n-----DRINK SIZES-----")
	fmt.Println("1 - See all orders\n2 - Edit an order\n3 - Delete an order\n0 - Go Back")

	return utils.GetValidatedNumber("Enter option: ", 0, 3)
}
