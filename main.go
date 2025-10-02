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
	itemsOrdered []interface{}
	totalPrice   float32
	timeOfOrder  time.Time
}

// type MenuItem struct {
// 	sandwich Sandwich
// }

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
		case 0:
			ifContinue = false
		}
	}
}

// #region Sandwich Logic
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

func sandwichLogic() {
	newSandwich := Sandwich{}
	newSandwich = createSandwich(newSandwich)
	newSandwich = calculateSandwichPrice(newSandwich)

	userValidation := validateSandwich(newSandwich)
	switch userValidation {
	case 1:
		fmt.Println("Your sandwich will be added to your order!")
	case 2:
		fmt.Println("My apologies... you can retry from the order menu.")
	}

}

func createSandwich(s Sandwich) Sandwich {

	s.ExtraMeat = false
	s.ExtraCheese = false

	sizeNum := ui.SandwichSizes()
	switch sizeNum {
	case 1:
		s.Size = "small"
	case 2:
		s.Size = "medium"
	case 3:
		s.Size = "large"
	}

	breadNum := ui.SandwichBreads()
	switch breadNum {
	case 1:
		s.Bread = "white"
	case 2:
		s.Bread = "wheat"
	case 3:
		s.Bread = "sourdough"
	case 4:
		s.Bread = "lettuce wrap"
	}

	meatNum := ui.SandwichMeats()
	switch meatNum {
	case 1:
		s.Meat = "turkey"
	case 2:
		s.Meat = "ham"
	case 3:
		s.Meat = "roast beef"
	case 4:
		s.Meat = "steak"
	case 0:
		s.Meat = "none"
	}
	if s.Meat != "none" {
		extraMeatNum := ui.ExtraMeatOption(s.Size)
		if extraMeatNum == 1 {
			s.ExtraMeat = true
		}
	}

	cheeseNum := ui.SandwichCheese()
	switch cheeseNum {
	case 1:
		s.Cheese = "american"
	case 2:
		s.Cheese = "cheddar"
	case 3:
		s.Cheese = "provolone"
	case 4:
		s.Cheese = "swiss"
	case 0:
		s.Cheese = "none"
	}
	if s.Cheese != "none" {
		extraCheeseNum := ui.ExtraCheeseOption(s.Size)
		if extraCheeseNum == 1 {
			s.ExtraCheese = true
		}
	}

	sauceNum := ui.SandwichSauce()
	switch sauceNum {
	case 1:
		s.Sauce = "mayo"
	case 2:
		s.Sauce = "mustard"
	case 3:
		s.Sauce = "ketchup"
	case 4:
		s.Sauce = "none"
	}

	for {
		toppingNum := ui.SandwichTopping()
		switch toppingNum {
		case 1:
			s.Toppings = append(s.Toppings, "onions")
		case 2:
			s.Toppings = append(s.Toppings, "lettuce")
		case 3:
			s.Toppings = append(s.Toppings, "peppers")
		case 4:
			s.Toppings = append(s.Toppings, "pickles")
		case 5:
			s.Toppings = append(s.Toppings, "tomatoes")
		case 0:
			s.Toppings = []string{"none"}
		}

		if toppingNum != 0 {
			fmt.Println("Would you like to add another topping?\n1 - Yes\n2 - No")
			moreToppings := utils.GetValidatedNumber("Enter option: ", 1, 2)
			if moreToppings == 2 {
				break
			}
		} else {
			break
		}
	}

	return s
}

func calculateSandwichPrice(s Sandwich) Sandwich {
	s.Price = 0.0

	switch s.Size {
	case "small":
		s.Price = 5.50

		if s.Meat != "none" {
			s.Price += 1.00
		}
		if s.ExtraMeat {
			s.Price += .50
		}
		if s.Cheese != "none" {
			s.Price += .75
		}
		if s.ExtraCheese {
			s.Price += .30
		}

	case "medium":
		s.Price = 7.00

		if s.Meat != "none" {
			s.Price += 2.00
		}
		if s.ExtraMeat {
			s.Price += 1.00
		}
		if s.Cheese != "none" {
			s.Price += 1.50
		}
		if s.ExtraCheese {
			s.Price += .60
		}

	case "large":
		s.Price = 8.50

		if s.Meat != "none" {
			s.Price += 3.00
		}
		if s.ExtraMeat {
			s.Price += 1.50
		}
		if s.Cheese != "none" {
			s.Price += 2.25
		}
		if s.ExtraCheese {
			s.Price += .90
		}
	}

	return s
}

func validateSandwich(s Sandwich) int {
	fmt.Println("\n---YOUR SANDWICH---")
	fmt.Println("Size:", s.Size)
	fmt.Println("Bread:", s.Bread)
	fmt.Println("Meat:", s.Meat)
	if s.ExtraMeat {
		fmt.Println("Extra Meat!")
	}
	fmt.Println("Cheese:", s.Cheese)
	if s.ExtraCheese {
		fmt.Println("Extra Cheese!")
	}
	fmt.Println("Sauce:", s.Sauce)

	for _, topping := range s.Toppings {
		fmt.Println("Topping:", topping)
	}
	fmt.Printf("\nTotal Price: $%.2f", s.Price)

	fmt.Println("\nIs this sandwich correct?\n1 - Yes\n2 - No")
	return utils.GetValidatedNumber("Enter option: ", 1, 2)
}

// #endregion

// #Chip Logic
type Chip struct {
	Type  string
	Size  string
	Price float32
}

func chipLogic() {
	chipType := getChipType()
	if chipType == "" {
		return
	}

	chipSize := getChipSize()
	chipPrice := calculateChipPrice(chipSize)

	newChip := Chip{
		Type:  chipType,
		Size:  chipSize,
		Price: chipPrice,
	}

}

func getChipType() string {
	var chipType string
	userChoice := ui.ChipTypes()

	switch userChoice {
	case 1:
		chipType = "doritos"
	case 2:
		chipType = "lays original"
	case 3:
		chipType = "cheetos"
	case 0:
		return ""
	}

	return chipType
}

func getChipSize() string {
	var chipSize string
	userChoice := ui.ChipTypes()

	switch userChoice {
	case 1:
		chipSize = "small"
	case 2:
		chipSize = "medium"
	case 3:
		chipSize = "large"
	}

	return chipSize
}

func calculateChipPrice(size string) float32 {
	var chipPrice float32

	switch size {
	case "small":
		chipPrice = 1.25
	case "medium":
		chipPrice = 2.00
	case "large":
		chipPrice = 2.70
	}

	return chipPrice
}

func drinkLogic() {}

func checkoutLogic() {}
