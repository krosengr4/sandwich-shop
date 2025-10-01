package main

import (
	"fmt"
	ui "sandwich-shop/user_interface"
	"sandwich-shop/utils"
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

func sandwichLogic() {
	userSandwich := Sandwich{}
	userSandwich = createSandwich(userSandwich)

}

func createSandwich(newSandwich Sandwich) Sandwich {
	userSize := ""
	userBread := ""
	userMeat := ""
	userCheese := ""
	userSauce := ""
	userToppings := []string{}
	extraMeat := false
	extraCheese := false

	sizeNum := ui.SandwichSizes()
	switch sizeNum {
	case 1:
		userSize = "small"
	case 2:
		userSize = "medium"
	case 3:
		userSize = "large"
	}

	breadNum := ui.SandwichBreads()
	switch breadNum {
	case 1:
		userBread = "white"
	case 2:
		userBread = "wheat"
	case 3:
		userBread = "sourdough"
	case 4:
		userBread = "lettuce wrap"
	}

	meatNum := ui.SandwichMeats()
	switch meatNum {
	case 1:
		userMeat = "turkey"
	case 2:
		userMeat = "ham"
	case 3:
		userMeat = "roast beef"
	case 4:
		userMeat = "steak"
	case 0:
		userMeat = "none"
	}
	if userMeat != "none" {
		extraMeatNum := ui.ExtraMeatOption(userSize)
		if extraMeatNum == 1 {
			extraMeat = true
		}
	}

	cheeseNum := ui.SandwichCheese()
	switch cheeseNum {
	case 1:
		userCheese = "american"
	case 2:
		userCheese = "cheddar"
	case 3:
		userCheese = "provolone"
	case 4:
		userCheese = "swiss"
	case 0:
		userCheese = "none"
	}
	if userCheese != "none" {
		extraCheeseNum := ui.ExtraCheeseOption(userSize)
		if extraCheeseNum == 1 {
			extraCheese = true
		}
	}

	sauceNum := ui.SandwichSauce()
	switch sauceNum {
	case 1:
		userSauce = "mayo"
	case 2:
		userSauce = "mustard"
	case 3:
		userSauce = "ketchup"
	case 4:
		userSauce = "none"
	}

	for {
		toppingNum := ui.SandwichTopping()
		switch toppingNum {
		case 1:
			userToppings = append(userToppings, "onions")
		case 2:
			userToppings = append(userToppings, "lettuce")
		case 3:
			userToppings = append(userToppings, "peppers")
		case 4:
			userToppings = append(userToppings, "pickles")
		case 5:
			userToppings = append(userToppings, "tomatoes")
		case 0:
			userToppings = []string{"none"}
		}

		fmt.Println("Would you like to add another topping?\n1 - Yes\n2 - No")
		moreToppings := utils.GetValidatedNumber("Enter option: ", 1, 2)
		if moreToppings == 2 {
			break
		}
	}

	newSandwich = Sandwich{
		Size:        userSize,
		Bread:       userBread,
		Meat:        userMeat,
		Cheese:      userCheese,
		Sauce:       userSauce,
		Toppings:    userToppings,
		ExtraMeat:   extraMeat,
		ExtraCheese: extraCheese,
	}

	return newSandwich
}

func chipLogic() {}

func drinkLogic() {}

func checkoutLogic() {}
