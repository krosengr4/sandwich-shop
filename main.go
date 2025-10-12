package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sandwich-shop/config"
	"sandwich-shop/database"
	"sandwich-shop/models"
	ui "sandwich-shop/user_interface"
	"sandwich-shop/utils"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var userOrder models.Order

func main() {
	fmt.Println("\n\t\t======WELCOME TO THE SANDWICH SHOP!=====")
	fmt.Println(strings.Repeat("_", 80))

	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to get a connection to the database: %v", err)

	}
	defer db.Close()

	mainMenuLogic(db)

	fmt.Println("\n\t\t=====GOODBYE!======")
	fmt.Println(strings.Repeat("_", 80))
}

func initDB() (*database.Database, error) {
	// Load env variables
	if err := config.LoadEnv(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get DB configurations with env variables
	dbConfig := config.GetDatabaseConfig()

	return database.GetConnection(dbConfig)
}

func mainMenuLogic(db *database.Database) {
	ifContinue := true

	for ifContinue {
		userChoice := ui.HomeScreen()

		switch userChoice {
		case 1:
			orderScreenLogic(db)
		case 0:
			ifContinue = false
		}
	}
}

func orderScreenLogic(db *database.Database) {
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
			checkoutLogic(db)
		case 0:
			userOrder = models.Order{}
			fmt.Println("Order cancelled.")
			ifContinue = false
		}
	}
}

// #region Sandwich Logic
func sandwichLogic() {
	newSandwich := models.Sandwich{}
	newSandwich = createSandwich(newSandwich)
	newSandwich = calculateSandwichPrice(newSandwich)

	userValidation := validateSandwich(newSandwich)
	switch userValidation {
	case 1:
		userOrder.ItemsOrdered = append(userOrder.ItemsOrdered, newSandwich)
		userOrder.Quantity += 1
	case 2:
		fmt.Println("My apologies... you can retry from the order menu.")
	}

}

func createSandwich(s models.Sandwich) models.Sandwich {

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

func calculateSandwichPrice(s models.Sandwich) models.Sandwich {
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

func validateSandwich(s models.Sandwich) int {

	s.PrintData()

	fmt.Println("\nIs this sandwich correct?\n1 - Yes\n2 - No")
	return utils.GetValidatedNumber("Enter option: ", 1, 2)
}

// #endregion

// #region Chip Logic
func chipLogic() {
	chipType := getChipType()
	if chipType == "" {
		return
	}

	chipSize := getChipSize()
	chipPrice := calculateChipPrice(chipSize)

	newChip := models.Chip{
		Type:  chipType,
		Size:  chipSize,
		Price: chipPrice,
	}

	newChip.PrintData()

	userOrder.ItemsOrdered = append(userOrder.ItemsOrdered, newChip)
	userOrder.Quantity += 1
	fmt.Println("Chips added to order!")
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
	userChoice := ui.ChipSizes()

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

// #endregion

// #region Drink Logic
func drinkLogic() {
	drinkType := getDrinkType()
	if drinkType == "" {
		return
	}

	drinkSize := getDrinkSize()
	drinkPrice := calculateDrinkPrice(drinkSize)

	newDrink := models.Drink{
		Type:  drinkType,
		Size:  drinkSize,
		Price: drinkPrice,
	}

	newDrink.PrintData()

	userOrder.ItemsOrdered = append(userOrder.ItemsOrdered, newDrink)
	userOrder.Quantity += 1
	fmt.Println("Drink has been added to your order!")
}

func getDrinkType() string {
	var drinkType string
	userChoice := ui.DrinkTypes()

	switch userChoice {
	case 1:
		drinkType = "soda"
	case 2:
		drinkType = "lemonade"
	case 3:
		drinkType = "milkshake"
	case 0:
		return ""
	}

	return drinkType
}

func getDrinkSize() string {
	var drinkSize string
	userChoice := ui.DrinkSizes()

	switch userChoice {
	case 1:
		drinkSize = "small"
	case 2:
		drinkSize = "medium"
	case 3:
		drinkSize = "large"
	}

	return drinkSize
}

func calculateDrinkPrice(size string) float32 {
	var drinkPrice float32

	switch size {
	case "small":
		drinkPrice = 2.00
	case "medium":
		drinkPrice = 2.50
	case "large":
		drinkPrice = 3.00
	}

	return drinkPrice
}

// #endregion

func checkoutLogic(db *database.Database) {
	scanner := bufio.NewScanner(os.Stdin)

	if len(userOrder.ItemsOrdered) == 0 {
		fmt.Println("Your order is empty! Please add items before checking out.")
		return
	}

	fmt.Println("\n---CHECKOUT---")
	fmt.Println("Please enter a name for your order: ")
	scanner.Scan()
	userOrder.CustomerName = utils.CapitalizeFirstLetter(strings.TrimSpace(scanner.Text()))

	userOrder.TotalPrice = calculateTotalPrice()
	userOrder.TimeOfOrder = time.Now()

	userValidate := validateOrder()
	switch userValidate {
	case 1:
		fmt.Println("Success! Your order will be right out!")
		db.AddOrder(&userOrder)
	case 2:
		fmt.Println("Oh no... Perhaps we can try again? My apologies.")
	}

	userOrder = models.Order{}
}

func calculateTotalPrice() float32 {
	var totalPrice float32

	for _, item := range userOrder.ItemsOrdered {
		totalPrice += item.GetPrice()
	}

	return totalPrice
}

func validateOrder() int {
	fmt.Println("\n" + strings.Repeat("*", 50))
	fmt.Println("ORDER FOR:", strings.ToUpper(userOrder.CustomerName))
	fmt.Println(strings.Repeat("*", 50))

	for _, item := range userOrder.ItemsOrdered {
		item.PrintData()
	}

	fmt.Println(strings.Repeat("*", 50))
	fmt.Printf("\nTotal Items: %d\n", userOrder.Quantity)
	fmt.Printf("Total Price: $%.2f\n\n", userOrder.TotalPrice)

	fmt.Println("\nIs this order correct?\n1 - Yes\n2 - No")
	return utils.GetValidatedNumber("Enter option: ", 1, 2)
}
