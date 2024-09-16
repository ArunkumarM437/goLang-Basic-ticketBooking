package main

import (
	"fmt"
)

func main() {
	var customerName = "Arunkumar M"
	const totalTickets = 100
	var remainingTickets = 100
	fmt.Printf("Hello %v Welcome to ticket booking!\n", customerName)
	fmt.Printf("Total number of tickets %v remaining are %v\n", totalTickets, remainingTickets)
	fmt.Println("Grab your tickets before it solds out")
}
