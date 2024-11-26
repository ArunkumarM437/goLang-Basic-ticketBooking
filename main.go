package main

import (
	"fmt"
)

func bookTickets(userTickets uint, remainingTickets uint) uint {
	fmt.Println("Processing....")
	remainingTickets = remainingTickets - userTickets
	fmt.Println("Booking Completed...")
	return remainingTickets
}
func getCost(userTickets int, price int) int {
	//get Ticket Cost for differenet type of Ticket booking
	price = price * userTickets
	return price
}

// implement features in features branch
func main() {
	conferenceName := "GoLang Training Conference"

	var firstName string
	var lastName string
	var email string
	var zipcode int
	var mobile uint
	var userTickets uint
	var ticketCost int
	var seatType int
	// Use Slices to handle the user details to store everyone's details
	const totalTickets = 100
	var remainingTickets uint = 100
	var bookingUsers []string //Dynamic array is called as slices
	for remainingTickets > 0 {
		fmt.Printf("Hi Greeting From %v Team :\n", conferenceName) //Greet Customer
		//Get User Details via function needs to implement
		fmt.Println("Enter your first name : ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name : ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email : ")
		fmt.Scan(&email)
		fmt.Println("Enter your mobile : ")
		fmt.Scan(&mobile)
		fmt.Println("Enter your postal-code : ")
		fmt.Scan(&zipcode)
		//Dynamic Array slices usage
		bookingUsers = append(bookingUsers, firstName+" "+lastName)
		// Notifying user about about ticket details
		fmt.Printf("\n%v Welcome to ticket booking!\n", bookingUsers[0])
		fmt.Printf("Total number of tickets %v remaining are %v\n", totalTickets, remainingTickets)
		fmt.Println("Grab your tickets before it solds out")
		//ticketBooking Logic
		fmt.Println("Enter number of ticket to be booked:")
		fmt.Scan(&userTickets) //use if conditionals to handle 0 to outOfRange tickets to book
		//Try user to select a normal seat or premium seats based on their need.
		fmt.Println("Enter 1 for Normal Seat  or 2 for Premium Seat :")
		fmt.Scan(&seatType)
		switch seatType {
		case 1:
			fmt.Println("Normal Seat Selected")
			ticketCost = 199
			ticketCost = getCost(int(userTickets), ticketCost)
		case 2:
			fmt.Println("Premium Seat Selected")
			ticketCost = 299
			ticketCost = getCost(int(userTickets), ticketCost)
		default:
			fmt.Println("Invalid Seat Selection")
		}
		if userTickets < remainingTickets {
			remainingTickets = bookTickets(userTickets, remainingTickets)
			fmt.Println("-------------------------------------Ticket Summary-------------------------------------")
			fmt.Printf("\n\tCustomer Name :%v \n \tTickets Booked : %v \n \tAvailable tickets after booking  :%v\n \tTickets sent to e-mail : %v\n \tContact-No : %v\n \tPostal-Code %v\n \tTotalCost : %v\n", bookingUsers[0], userTickets, remainingTickets, email, mobile, zipcode, ticketCost)
			fmt.Println("Happy Ticket Bokking...")
			if remainingTickets == 0 {
				fmt.Println("Conference is full,Try next year...")
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Printf("You cannot book all tickets once ...\n")
		} else {
			fmt.Printf("We only have %v remaining,you can't book %v tickets...\n", remainingTickets, userTickets)
			fmt.Print("Booking Failed...Try Again\n")
		}
	}
}
