package main

import (
	"fmt"
	"ticket-app/helper"
)

type userData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets int
}

var firstNames = []string{}

func main() {
	var remainingTickets int = 50

	bookings := make([]userData, 0)

	for remainingTickets > 0 && len(bookings) <= 50 {

		firstName, lastName, email, noOfTickets := takeInput()

		isValidName, isValidEmail, isValidNoOfTicket := helper.CheckValidInputs(firstName, lastName, email, noOfTickets, remainingTickets)

		if isValidName && isValidEmail && isValidNoOfTicket {
			processTicket(firstName, lastName, email, bookings, noOfTickets, remainingTickets)
		} else {
			if !isValidName {
				fmt.Printf("Name is not valid. Atleast 2 characters are needed: %v", firstName)
			}
			if !isValidEmail {
				fmt.Printf("Email is not valid. Need to have @: %v", email)
			}
			if !isValidNoOfTicket {
				fmt.Printf("Number of ticket is not valid : %v", noOfTickets)
			}
		}

	}
}

func takeInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var noOfTickets int

	fmt.Println("Enter the first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter the last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter the email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&noOfTickets)

	return firstName, lastName, email, noOfTickets
}

func getBookingsFirstNamesOnly(bookings []userData) []string {
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func processTicket(firstName string, lastName string, email string, bookings []userData, noOfTickets int, remainingTickets int) {
	remainingTickets = remainingTickets - noOfTickets
	bookings = append(bookings, userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: noOfTickets,
	})

	if remainingTickets == 0 {
		fmt.Printf("We don've have anymore tickets left!\n")
	}
	getBookingsFirstNamesOnly(bookings)
	fmt.Printf("Bookings are %v \n", firstNames)
}
