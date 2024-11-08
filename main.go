package main

import "fmt"



func main() {
	// := is used to declare and assign a value to a variable in one line.
	conferenceName := "Go Conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = conferenceTickets;

	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v available tickets and %v tickets are still available \n", conferenceTickets, remainingTickets)
	
	fmt.Println("Please to book your ",conferenceName,"tickets by filling the form below")

	// Types of variables in Go are:
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T \n", conferenceTickets, remainingTickets)

	var bookings []string // array of strings



	var firstName string
	var lastName string
	var email string
	var userTickets uint8



	// ask user for their name, email and tickets
	fmt.Print("Enter your first name: ")
	// fmt.Scan() is used to read user input from the console. and & is used to pass the value to the variable and used as a pointer.
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName + " " + lastName + " " + email)

	// check if the user is booking more tickets than available
	fmt.Printf("Your name is %v %v and you want to book %v tickets \n", firstName, lastName, userTickets)
	fmt.Printf("You will receive confirmation of your booking via email in %v \n", email)
	fmt.Printf("You have %v tickets remaining  for %v\n", remainingTickets, conferenceName)

	fmt.Printf("Here is the booking information %v \n", bookings )
	fmt.Println("Total Array", len(bookings))

	fmt.Println("Thank you for booking with us")
}