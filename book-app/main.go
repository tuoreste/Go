package main

import	(
	"fmt"
	"time"
	"sync"
	"github.com/tuoreste/Go/book-app/sources"
)

const conferenceTickets = 50
var conferenceName = "Go conference"
var remainingTickets uint = 50
// var bookings = make([](map[string]string), 0) //empty list of maps in go
var bookings = make([]UserData, 0) //empty list of maps in go

type UserData struct {
	firstName string
	secondName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// for {
	firstName, secondName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, secondName, email, userTickets, remainingTickets)

	if (!isValidTicketNumber) {
		fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
	} else if (isValidName && isValidEmail) {
			
		bookTicket(userTickets, firstName, secondName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, secondName, email)
			
		// call print first names
		firstNames := printFirstNames()
		fmt.Printf("The first names of the bookings are: %v\n", firstNames)
			
		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our conference is fully booked out. Come back next year.")
			// break
		}
	} else {
		fmt.Println("Invalid Name or Email Input")
		// break
	}
	// }
	wg.Wait()
}

func	greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your ticket to attend")
}

func	printFirstNames() [] string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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

	var userData = UserData {
		firstName: firstName,
		secondName: secondName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, secondName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("######################################")
	wg.Done()
}
