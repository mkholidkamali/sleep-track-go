package main

import (
	"SleepTrack/seeders"
	"SleepTrack/template"
	"fmt"
)

func main() {
	// Initialize variable
	var action int
	var loggedUserIndex int

	// Seed users
	users := seeders.SeedUsers()

	// Show user login
	loggedUserIndex = template.PrintIntroduction(&action, &users)

	// Show landing page
	if loggedUserIndex != -1 {
		template.PrintLandingPage(&action, &users[loggedUserIndex])
		fmt.Println(users)
	}
}
