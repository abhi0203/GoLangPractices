package main

import (
	"fmt"
	"strings"
)

type customerInfo struct {
	firstName     string
	lastName      string
	email         string
	ticketsWanted uint
}

func main() {

	var conferanceName string = "Go Conferance"
	var remainingTickets uint = 50

	var firstNameList = []string{}
	var customerInfoList = []customerInfo{}

	fmt.Printf("Welcome to our Ticket Selling App for %v !!! \n", conferanceName)
	fmt.Printf("We have %v remaining tickets \n", remainingTickets)

	for {

		customer := customerInfo{}

		fmt.Println("Please enter your first name")
		fmt.Scan(&customer.firstName)

		fmt.Println("Please enter your last name")
		fmt.Scan(&customer.lastName)

		fmt.Println("Please enter the number of tickets you want to purchase.")
		fmt.Scan(&customer.ticketsWanted)

		fmt.Println("Please enter your email")
		fmt.Scan(&customer.email)

		// Validate User Inputs
		isValidName := len(customer.firstName) > 2 && len(customer.lastName) > 2
		isValidEmail := strings.Contains(customer.email, "@")
		isValidTicketCount := customer.ticketsWanted > 0 && customer.ticketsWanted <= remainingTickets

		if isValidName && isValidEmail && isValidTicketCount {

			firstNameList = append(firstNameList, customer.firstName)

			if customer.ticketsWanted > remainingTickets {
				fmt.Printf("Only %v tickets are available and you are asking for %v tickets. Please enter correct number \n", remainingTickets, customer.ticketsWanted)
				continue
			} else {
				remainingTickets = remainingTickets - customer.ticketsWanted
			}

			fmt.Printf("Your order for %v tickets is now completed %v %v. You will receive confirmation on your email at %v \n", customer.ticketsWanted, customer.firstName, customer.lastName, customer.email)
		} else {
			fmt.Println("You have not enterted correct information. Please enter correct data.")
			continue
		}

		customerInfoList = append(customerInfoList, customer)

		fmt.Println("************************")
		fmt.Printf("Name of people who have purchased tickets for conferance are %v\n", firstNameList)
		fmt.Printf("Complete information of customers who purchased the tickets %v\n", customerInfoList)
		fmt.Printf("We have %v remaining tickets \n", remainingTickets)

		if remainingTickets == 0 {
			break
		}

	}

	fmt.Println("All the tickets are sold. Conferance is full")

}
