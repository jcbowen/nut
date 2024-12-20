package example

import (
	"fmt"
	"github.com/jcbowen/nut"
)

// This example connects to NUT, authenticates and returns the first UPS listed.
func getUPS() {
	client, connectErr := nut.Connect("127.0.0.1")
	if connectErr != nil {
		fmt.Print(connectErr)
	}
	_, authenticationError := client.Authenticate("username", "password")
	if authenticationError != nil {
		fmt.Print(authenticationError)
	}
	upsList, listErr := client.GetUPSList()
	if listErr != nil {
		fmt.Print(listErr)
	}
	fmt.Print("First UPS", upsList[0])

	// Let's say you only need to get the first UPS
	ups := upsList[0]
	fmt.Println("First UPS Information:")

	// Get the variables for UPS
	variables, err := ups.GetVariables()
	if err != nil {
		fmt.Println("Error getting UPS variables:", err)
		return
	}

	// Find and print the value of a specific variable
	for _, v := range variables {
		// Print the battery level
		if v.Name == "battery.charge" {
			fmt.Printf("Battery Charge: %v\n", v.Value)
		}
		// Print device load
		if v.Name == "ups.load" {
			fmt.Printf("UPS Load: %v\n", v.Value)
		}
		// Print the UPS status
		if v.Name == "ups.status" {
			fmt.Printf("UPS Status: %v\n", v.Value)
		}
	}
}
