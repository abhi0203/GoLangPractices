package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferanceName string = "Go Conferance"
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var ticketsWanted uint
	var email string

	var firstNameList = []string{}

	fmt.Printf("Welcome to our Ticket Selling App for %v !!! \n", conferanceName)
	fmt.Printf("We have %v remaining tickets \n", remainingTickets)

	for {

		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Please enter the number of tickets you want to purchase.")
		fmt.Scan(&ticketsWanted)

		fmt.Println("Please enter your email")
		fmt.Scan(&email)

		// Validate User Inputs
		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketCount := ticketsWanted > 0 && ticketsWanted <= remainingTickets

		if isValidName && isValidEmail && isValidTicketCount {

			firstNameList = append(firstNameList, firstName)

			if ticketsWanted > remainingTickets {
				fmt.Printf("Only %v tickets are available and you are asking for %v tickets. Please enter correct number \n", remainingTickets, ticketsWanted)
				continue
			} else {
				remainingTickets = remainingTickets - ticketsWanted
			}

			fmt.Printf("Your order for %v tickets is now completed %v %v. You will receive confirmation on your email at %v \n", ticketsWanted, firstName, lastName, email)
		} else {
			fmt.Println("You have not enterted correct information. Please enter correct data.")
			continue
		}

		fmt.Println("************************")
		fmt.Printf("List of people who have purchased tickets for conferance are %v\n", firstNameList)
		fmt.Printf("We have %v remaining tickets \n", remainingTickets)

		if remainingTickets == 0 {
			break
		}

	}

	fmt.Println("All the tickets are sold. Conferance is full")

}
