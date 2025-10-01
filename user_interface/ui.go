package ui

import (
	"fmt"
	"sandwich-shop/utils"
)

func DisplayHome() int {
	fmt.Println("\n-----OPTIONS-----")
	fmt.Println("1 - Order\n0 - Exit")

	userChoice := utils.GetValidatedNumber("Enter option:", 0, 1)
	return userChoice
}
