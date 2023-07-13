package main

import (
	"fmt"
	"strings"
)

func main() {

	confereneceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceName type is %T\nconfereneceTickets type is %T\nremainingTickets type is %T\n", confereneceName, conferenceTickets, remainingTickets)

	fmt.Printf("WELCOME TO OUR %v BOOKING APPLICATION\n", confereneceName)
	fmt.Println("We hava total of", conferenceTickets, "tickets and remaining", remainingTickets, "tickets are still available")
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email : ")
		fmt.Scan(&email)

		fmt.Println("Enter your ticket count : ")
		fmt.Scan(&userTickets)

		if userTickets <= int(remainingTickets) {
			remainingTickets = remainingTickets - uint(userTickets)

			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("Whole slice : %v\n", bookings)
			// fmt.Printf("slice first element : %v\n", bookings[0])
			// fmt.Printf("slice type : %T\n", bookings)
			// fmt.Printf("slice length : %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remianing for the %v\n", remainingTickets, confereneceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, all tickets are sold out. Comeback next year!")
				break
			}
		} else {
			fmt.Printf("Sorry, we only have %v tickets remaining. So, you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

}
