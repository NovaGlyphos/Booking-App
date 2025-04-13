package main

import (
	// "bufio"
	. "fmt"
	// "os"
)

func main(){
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var numberOfTickets int
	Printf("Enter your firstName: ")
	Scan(&firstName)
	Printf("Enter your lastName:")
	Scan(&lastName)
	Printf("Enter your email:")
	Scan(&email)
	Printf("Enter the number of tickets you want to book:")
	Scan(&numberOfTickets)
	Printf("%v %v having email:%v has booked %v tickets for the conference\n",firstName,lastName,email,numberOfTickets)

	remainingTickets = uint(conferenceTickets)-uint(numberOfTickets)
	Printf("There are %v tickets remaining\n",remainingTickets)

	bookings[0] = firstName+" "+lastName
	Printf("Whole array: %v\n",bookings)
	Printf("First element of array is %v\n",bookings[0])
	Printf("Type of array is: %T",bookings)
	Printf("Type of element in array: %T",bookings[0])
}
