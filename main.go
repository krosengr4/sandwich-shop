package main

import (
	"fmt"
	ui "sandwich-shop/user_interface"
	"strings"
)

func main() {
	fmt.Println("\n\t\t======WELCOME TO THE SANDWICH SHOP!=====")
	fmt.Println(strings.Repeat("_", 80))

	mainMenuLogic()

	fmt.Println("\n\t\t=====GOODBYE!======")
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

func orderScreenLogic() {}
