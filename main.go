package main

import (
	"SleepTrack/model"
	"SleepTrack/seeders"
	"SleepTrack/template"
	"fmt"
)

func main() {
	// Initialize variable
	var action int
	var loggedUserIndex int
	var nUsers int
	var users [model.MaxUser]model.User

	// Seed users
	users = seeders.SeedUsers(&nUsers)

	// Show user login
	// loggedUserIndex = template.PrintIntroduction(&action, &users, &nUsers)
	loggedUserIndex = 0

	// Show landing page
	if loggedUserIndex != -1 {
		template.PrintLandingPage(&action, &users[loggedUserIndex])
		fmt.Println(users)
	}
}
