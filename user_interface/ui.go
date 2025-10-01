package ui

import (
	"fmt"
	"sandwich-shop/utils"
)

func DisplayHome() int {
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - Order\n0 - Exit")

	userChoice := utils.GetValidatedNumber("Enter option: ", 0, 1)
	return userChoice
}

func DisplayOrderScreen() int {
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - Order Sandwich\n2 - Order Chips\n3 - Order Drink\n4 - Checkout\n0 - Cancel")

	userChoice := utils.GetValidatedNumber("Enter option: ", 0, 4)
	return userChoice
}
