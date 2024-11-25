package main

import	(
	"fmt"
	"strings"
	"github.com/tuoreste/Go/book-app/sources"
)

const conferenceTickets = 50
var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, secondName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, secondName, email, userTickets, remainingTickets)

		if (!isValidTicketNumber) {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		} else if (isValidName && isValidEmail) {
				
			bookTicket(userTickets, firstName, secondName, email)
				
			// call print first names
			firstNames := printFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)
				
			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is fully booked out. Come back next year.")
				break
			}
		} else {
			fmt.Println("Invalid Name or Email Input")
			break
		}
	}
}

func	greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your ticket to attend")
}

func	printFirstNames() [] string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func	getUserInput()(string, string, string, uint){
	var firstName string
	var secondName string
	var email string
	var userTickets uint
	
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your second name: ")
	fmt.Scan(&secondName)

	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, secondName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, secondName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + secondName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, secondName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
