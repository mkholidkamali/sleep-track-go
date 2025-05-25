package main

import (
	"SleepTrack/seeders"
	"SleepTrack/template"
)

func main() {
	// Initialize variable
	var action int

	// Seed users
	users := seeders.SeedUsers()

	// Show landing page
	template.PrintLandingPage(&action)

	//
}
