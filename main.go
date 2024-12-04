package main

import (
	"fmt"
)

// currently the user can able to book more number of tickets but they need only to book limited number of tickets use some logic to make this possible and store ticket details in MongoDB
// single user limit is 6 Tickets implement this user booking logic before bookTickets function
func bookTickets(userTickets uint, remainingTickets uint) uint {
	fmt.Println("Processing....")
	remainingTickets = remainingTickets - userTickets
	fmt.Println("Booking Completed...")
	return remainingTickets
}
func getCost(userTickets uint, price float32) float32 {
	//get Ticket Cost for differenet type of Ticket booking
	price = price * float32(userTickets)
	return price
}

// implement features in features branch
func main() {

	conferenceName := "GoLang Training Conference"
	//user data
	var firstName string
	var lastName string
	var email string
	var zipcode int
	var mobile uint
	var userTickets uint
	var ticketCost float32
	var seatType string
	var validBooking bool
	var taxAmount float32
	// Conference data and collection of bookings
	var totatEarnings float32
	// Use Slices to handle the user details to store everyone's details
	const totalTickets = 10
	var remainingTickets uint = totalTickets
	var bookingUsers []string //Dynamic array is called as slices
	var counter int = 0
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
		fmt.Printf("\n%v Welcome to ticket booking!\n", bookingUsers[counter])
		fmt.Printf("Total number of tickets %v remaining are %v\n", totalTickets, remainingTickets)
		fmt.Println("Grab your tickets before it solds out")
		//ticketBooking Logic
		fmt.Println("Enter number of ticket to be booked:")
		fmt.Scan(&userTickets) //use if conditionals to handle 0 to outOfRange tickets to book
		var limitofUserTicket uint = 6
		if userTickets > limitofUserTicket || userTickets <= 0 {
			fmt.Println("You entered wrong number of Ticket")
			validBooking = false
			fmt.Printf("The lower limit is 0 and upper limit of ticket booking is %v\n", limitofUserTicket)
			fmt.Println("Booking failed...Try Again")
		} else {
			validBooking = true
		}
		if userTickets <= remainingTickets && validBooking {
			//Try user to select a normal seat or premium seats based on their need.
			fmt.Println("Enter N for Normal Seat  or P for Premium Seat :")
			fmt.Scan(&seatType)
			for {
				if seatType == "N" || seatType == "P" {
					break
				} else {
					fmt.Println("Invalid Seat Selection")
					fmt.Println("Enter N for Normal Seat - No Tax  or P for Premium Seat Includes Tax (Case-Sensitive) :")
					fmt.Scan(&seatType)
				}
			}
			switch seatType {
			case "N":
				fmt.Println("Normal Seat Selected")
				ticketCost = 199
				ticketCost = getCost(userTickets, ticketCost)
				totatEarnings += ticketCost
			case "P":
				fmt.Println("Premium Seat Selected")
				ticketCost = 299
				taxAmount = (2.5 / ticketCost) * 100
				fmt.Printf("Tax For Selceted Seat : %v\n", taxAmount)
				ticketCost = getCost(userTickets, ticketCost+taxAmount)
				totatEarnings += ticketCost
			default:
				fmt.Println("Invalid Seat Selection")
			}
			remainingTickets = bookTickets(userTickets, remainingTickets)
			fmt.Println("-------------------------------------Ticket Summary-------------------------------------")
			fmt.Printf("\n\tCustomer Name :%v \n \tTickets Booked : %v \n \tAvailable tickets after booking  :%v\n \tTickets sent to e-mail : %v\n \tContact-No : %v\n \tPostal-Code %v\n \tTotalCost : %v\n", bookingUsers[counter], userTickets, remainingTickets, email, mobile, zipcode, ticketCost)
			counter += 1
			fmt.Println("Happy Ticket Bokking...")
			if remainingTickets == 0 {
				fmt.Println("Conference is full,Try next year...")
				break
			}
		} else if userTickets == totalTickets {
			fmt.Printf("You cannot book all tickets once ...\n")
			fmt.Print("Booking Failed...Try Again\n")

		} else if userTickets > remainingTickets && !validBooking {
			fmt.Printf("We only have %v remaining,you can't book %v tickets...\n", remainingTickets, userTickets)
		}
	}
	fmt.Printf("Total Earning's after Bookings :%v\n", totatEarnings)
	fmt.Println("Total User Bookings:", len(bookingUsers))
	fmt.Println("Thank You...Mikka Nandri...Arigato Gosaimasu")
}
