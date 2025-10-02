package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Helper function to get a validated number within range
func GetValidatedNumber(prompt string, min, max int) int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if num, err := strconv.Atoi(input); err == nil {
			if num >= min && num <= max {
				return num
			}
			fmt.Printf("Number must be between %d and %d. Try again.\n", min, max)
		} else {
			fmt.Println("Invalid number. Please try again.")
		}
	}
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
