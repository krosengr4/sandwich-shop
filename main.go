package main

import (
	"fmt"
	ui "sandwich-shop/user_interface"
	"strings"
	"time"
)

type Order struct {
	orderId      int
	customerName string
	totalPrice   float32
	timeOfOrder  time.Time
}

func main() {
	fmt.Println("\n\t\t======WELCOME TO THE SANDWICH SHOP!=====")
	fmt.Println(strings.Repeat("_", 80))

	mainMenuLogic()

	fmt.Println("\n\t\t=====GOODBYE!======")
	fmt.Println(strings.Repeat("_", 80))
}

func mainMenuLogic() {
	ifContinue := true

	for ifContinue {
		userChoice := ui.DisplayHome()

		switch userChoice {
		case 1:
			orderScreenLogic()
		case 0:
			ifContinue = false
		}
	}
}

func orderScreenLogic() {
	ifContinue := true

	for ifContinue {
		userChoice := ui.DisplayOrderScreen()

		switch userChoice {
		case 1:
			sandwichLogic()
		case 2:
			chipLogic()
		case 3:
			drinkLogic()
		case 4:
			checkoutLogic()
		}
	}
}

func sandwichLogic() {}

func chipLogic() {}

func drinkLogic() {}

func checkoutLogic() {}
