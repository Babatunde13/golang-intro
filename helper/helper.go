package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) bool {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	if !isValidName {
		fmt.Println("First name or last name you entered is too short")
	}
	if !isValidEmail {
		fmt.Println("Please enter a valid email address")
	}
	if !isValidTicketNumber {
		fmt.Println("Please enter a valid number of tickets")
	}
	if isValidName && isValidEmail && isValidTicketNumber {
		return true
	}
	return false
}