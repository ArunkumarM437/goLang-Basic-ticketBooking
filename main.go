package main

import (
	"fmt"
)

func bookTickets(userTickets uint, remainingTickets int) int {
	fmt.Println("Processing....")
	remainingTickets = remainingTickets - int(userTickets)
	fmt.Println("Booking Completed...")
	return remainingTickets
}

func main() {
	conferenceName := "GoLang Training Conference"

	var firstName string
	var lastName string
	var email string
	var zipcode int
	var mobile uint
	var userTickets uint
	// Use Slices to handle the user details to store everyone's details
	const totalTickets = 100
	var remainingTickets int = 100
	var bookingUsers []string //Dynamic array is called as slices

	fmt.Printf("Hi Greeting From %v Team :\n", conferenceName) //Greet Customer
	//Get User Details
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
	//Array slices usage
	bookingUsers = append(bookingUsers, firstName+" "+lastName)
	// Notifying user about about ticket details
	fmt.Printf("\n%v Welcome to ticket booking!\n", bookingUsers[0])
	fmt.Printf("Total number of tickets %v remaining are %v\n", totalTickets, remainingTickets)
	fmt.Println("Grab your tickets before it solds out")
	//ticketBooking Logic
	fmt.Println("Enter number of ticket to be booked:")
	fmt.Scan(&userTickets) //use if conditionals to handle 0 tickets to booked
	remainingTickets = bookTickets(userTickets, remainingTickets)
	fmt.Println("-------------------------------------Ticket Summary-------------------------------------")
	fmt.Printf("\n\tCustomer Name :%v \n \tTickets Booked : %v \n \tAvailable tickets after booking  :%v\n \tTickets sent to e-mail : %v\n \tContact-No : %v\n \tPostal-Code %v\n", bookingUsers[0], userTickets, remainingTickets, email, mobile, zipcode)
}
