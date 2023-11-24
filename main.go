package main

import (
	"fmt"
	"go-cli/helper"
	"sync"
	"time"
)

type UserData struct {
	userName string
	email string
	numberOfTickets uint
	confirmationSent bool
}

var wg = sync.WaitGroup{}

var bookings = make([]UserData, 0)
const conferenceName = "Golang Conference"
const conferenceTickets uint = 50
var remainingTickets = conferenceTickets



func main () {
  
	fmt.Printf("Welcome to %v booking system!\n", conferenceName)
	fmt.Printf("We currently have %v tickets remaining!\n", remainingTickets)
	fmt.Println("Get your tickets here to attend the conference!")

	// for remainingTickets > 0 {
		username, email, ticketQuantity := helper.GetUserInput()
		isUserInputValid := helper.ValidateUserInput(username, email, ticketQuantity, remainingTickets)

		if(isUserInputValid) {
			remainingTickets = bookTickets(username, email, ticketQuantity)
		}
	// }
	wg.Wait()
}

func bookTickets(username, email string, numberOfTickets uint) uint {

	userData := new(UserData)

	userData.userName = username
	userData.email = email
	userData.numberOfTickets = numberOfTickets
	
	bookings = append(bookings, *userData)
	remainingTickets -= numberOfTickets

	fmt.Printf("\nThankyou %v, you have booked %v tickets!\nYou will recieve an email confirmation at %v\n", username, numberOfTickets, email)

	if remainingTickets == 0 {
		fmt.Println("All tickets for this conference are sold out!")
	} else {
		printBookingSummary()
	}
	wg.Add(1)
	go sendTicketConfirmation(userData)

	return remainingTickets
}

func sendTicketConfirmation(booking *UserData) {
	msg:= fmt.Sprintf("Hi %v, \n Your %v tickets have been booked successfully!", booking.userName, booking.numberOfTickets)

	time.Sleep(time.Second * 10)
	fmt.Println("####################")
	fmt.Println(msg)
	fmt.Println("####################")

	booking.confirmationSent = true
	wg.Done()
}

func printBookingSummary() {
	fmt.Println("Booking summary")
	for _, booking := range bookings {
		fmt.Printf("\nusername: %v, email: %v, tickets: %v, confirmationSent %v.", booking.userName, booking.email, booking.numberOfTickets, booking.confirmationSent)
	}
	fmt.Println("")
	fmt.Println("Booking summary end")
}