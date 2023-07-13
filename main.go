package main
import "fmt"

func main(){

	confereneceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName type is %T\nconfereneceTickets type is %T\nremainingTickets type is %T\n",confereneceName, conferenceTickets, remainingTickets)

	fmt.Printf("WELCOME TO OUR %v BOOKING APPLICATION\n",confereneceName)
	fmt.Println("We hava total of",conferenceTickets,"tickets and remaining",remainingTickets,"tickets are still available")
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n",firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remianing for the %v\n", remainingTickets, confereneceName)

}