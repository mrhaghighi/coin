package main

import "fmt"

func main() {
	var selectedService string

	// Show greeting message
	fmt.Println("###########################")
	fmt.Println("Welcome to Ultimate Trade")
	fmt.Println("###########################")

	// Get the service that user wants to use it
	fmt.Println("\nSelect the service you want: (ip, coin)")
	fmt.Scanln(&selectedService)

	// Call selected function
	switch selectedService {
		case "ip": getUserIpFromApi()
		case "coin": getPrice()
		default: fmt.Println("The service you selected is wrong!")
	}
}