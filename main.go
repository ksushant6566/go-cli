package main

import (
	"fmt"
	"go-cli/helper"
)

var bookings = []string{}
const conferenceName = "Golang Conference"
const conferenceTickets uint = 50
var remainingTickets = conferenceTickets

func main () {
  
	fmt.Printf("Welcome to %v booking system!\n", conferenceName)
	fmt.Printf("We currently have %v tickets remaining!\n", remainingTickets)
	fmt.Println("Get your tickets here to attend the conference!")

	for remainingTickets > 0 {
		
		username, email, ticketQuantity := helper.GetUserInput()
		isUserInputValid := helper.ValidateUserInput(username, email, ticketQuantity, remainingTickets)

		if(isUserInputValid) {
			remainingTickets = bookTickets(username, email, ticketQuantity)
		} else {
			continue
		}

	}

}

func bookTickets(username, email string, numberOfTickets uint) uint {
	
	bookings = append(bookings, username)
	remainingTickets -= numberOfTickets

	fmt.Printf("\nThankyou %v, you have booked %v tickets!\nYou will recieve an email confirmation at %v\n", username, numberOfTickets, email)

	if remainingTickets == 0 {
		fmt.Println("All tickets for this conference are sold out!")
	} else {
		fmt.Printf("%v tickets remaining\n", remainingTickets)
	}

	return remainingTickets
}