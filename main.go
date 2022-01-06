package main

import (
	"fmt"
	"strings"
)

func main() {
	// conferenceTotalTickets := 50
	var remainingTickets int = 50

	bookings := []string{}

	for remainingTickets > 0 && len(bookings) <= 50 {
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

		if noOfTickets > remainingTickets {
			fmt.Printf("Sorry! We have only %v tickets left.", remainingTickets)
			break
		}

		remainingTickets = remainingTickets - noOfTickets
		bookings = append(bookings, firstName+" "+lastName)

		if remainingTickets == 0 {
			fmt.Printf("We don've have anymore tickets left!\n")
		}

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("Bookings are %v \n", firstNames)
	}
}
