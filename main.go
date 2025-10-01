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
		userChoice := ui.HomeScreen()

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
		userChoice := ui.OrderScreen()

		switch userChoice {
		case 1:
			userSandwich := new(Sandwich)
			userSandwich.sandwichLogic()
		case 2:
			chipLogic()
		case 3:
			drinkLogic()
		case 4:
			checkoutLogic()
		}
	}
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
}

func (*Sandwich) sandwichLogic() {}

func chipLogic() {}

func drinkLogic() {}

func checkoutLogic() {}
