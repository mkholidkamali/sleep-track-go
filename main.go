package main

import (
	"SleepTrack/seeders"
	"SleepTrack/template"
	"fmt"
)

func main() {
	// Initialize variable
	var action int

	// Seed users
	users := seeders.SeedUsers()

	// Show user login

	// Show landing page
	template.PrintLandingPage(&action, &users[0])
	fmt.Println(users)

	//
}
