package template

import (
	"SleepTrack/helpers"
	"SleepTrack/model"
	"SleepTrack/services/landingPage"
	"fmt"
)

func PrintLandingPage(action *int, user *model.User) {
	fmt.Println()
	fmt.Println("=============================")
	fmt.Println("Zzz...")
	fmt.Printf("Hoam... hm?, ehhh. Halo %s selamat pagi !!\n", user.Username)
	fmt.Println("Gimana tidur kamu hari ini?")
	fmt.Println("=============================")

	for *action != 4 {
		landingPage.InputSelectedAction(action, user)

		switch *action {
		case 1:
			user.Sleeps[user.TotalSleeps] = PrintInputNewHistory(&user.Sleeps, user.TotalSleeps)
			user.TotalSleeps++
		case 2:
			PrintShowSleepHistory(action, user)
		case 3:
			PrintAboutPage()
		default:
			helpers.PrintExitMessage()
		}
	}
}
