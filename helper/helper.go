package helper

import "strings"

func CheckValidInputs(firstName string, lastName string, email string, noOfTicket int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNoOfTicket := noOfTicket >= 0 && noOfTicket <= remainingTickets
	return isValidName, isValidEmail, isValidNoOfTicket
}
