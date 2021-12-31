package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets int = 50
var bookings = make([]Booking, 0)

type Booking struct {
	firstName string
	lastName  string
	email     string
	numberOfTickets   int
	// isOptedInForNewsletter bool
}

var wg = sync.WaitGroup {}

func main() {
	routines()
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		// validate inputs
		isValidInput := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidInput {
			booking := bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(booking)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
			fmt.Printf("We have %v tickets remaining\n", remainingTickets)
		} else {
			continue
		}
	}
	firstNames := printFirstNames()
	fmt.Println(bookings)
	fmt.Printf("The first name of our bookings %v\n", firstNames)
	// // switch - case
	// city := "London"
	// switch city {
	// 	case "London", "london":
	// 		fmt.Println("London is the capital of Great Britain")
	// 	case "Paris":
	// 		fmt.Println("Paris is the capital of France")
	// 	case "New York":
	// 		fmt.Println("New York is the capital of USA")
	// 	default:
	// 		fmt.Println("I don't know that city")
	// }
		wg.Wait()
}

// if - else if - else

func greetUsers() {
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T and conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have a total number of %v tickets and we have %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	// fmt.Printf("The first name of our bookings %v\n", firstNames)
	return firstNames
}

func bookTicket(userTickets int, firstName string, lastName string, email string) Booking {
	remainingTickets = remainingTickets - userTickets
	var booking = Booking {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, booking)
	return booking
}

func sendTicket(booking Booking) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v\n", booking.numberOfTickets, booking.email)
	fmt.Println("###########")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, booking.email)
	fmt.Println("###########")
	wg.Done()
}

// map - slice - struct
// map()\\
