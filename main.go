package main

import (
	"SleepTrack/seeders"
	"SleepTrack/template"
)

func main() {
	// Initialize variable
	var action int
	var loggedUserIndex int
	var nUsers int

	// Seed users
	users := seeders.SeedUsers(&nUsers)

	// Show user login
	loggedUserIndex = template.PrintIntroduction(&action, &users, &nUsers)

	// Show landing page
	if loggedUserIndex != -1 {
		template.PrintLandingPage(&action, &users[loggedUserIndex])
	}
}
