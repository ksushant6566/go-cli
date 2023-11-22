package helper

import (
	"fmt"
	"strings"
)

func GetUserInput() (string, string, uint) {
		var username string
		var ticketQuantity uint
		var email string

		fmt.Println("Enter your name: ")
		fmt.Scan(&username)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets you want to book: ")
		fmt.Scan(&ticketQuantity)

		return username, email, ticketQuantity
}

func ValidateUserInput(username, email string, ticketQuantity, remainingTickets uint,) bool {

	isUserNameValid := len(username) >= 3
	isEmailValid := strings.Contains(email, "@") && len(email) > 4
	isTicketQuantityValid := ticketQuantity <= remainingTickets

	if !isUserNameValid {
		fmt.Println("Invalid username, username must be atleast 3 chars long")
	}
	if !isEmailValid {
		fmt.Println("Invalid email, email must be atleast 5 chars long and must contain '@'")
	}
	if !isTicketQuantityValid {
		fmt.Println("Invalid number of tickets, number of tickets must be less than remaining tickets")
	}
	
	return isUserNameValid && isEmailValid && isTicketQuantityValid
}