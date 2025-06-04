package template

import (
	"SleepTrack/helpers"
	"SleepTrack/model"
	"SleepTrack/services/showSleepHistory"
	"fmt"
)

func PrintShowSleepHistory(action *int, user *model.User) {
	// Show all data
	showAllHistory(user)

	for *action != 5 {
		// Input selected action
		showSleepHistory.InputSelectedAction(action, user)

		switch *action {
		case 1:
			//
		case 2:
			//
		case 3:
			//
		case 4:
			//
		case 5:
			//
		}
	}
}

/**
* Private Method
**/
func showAllHistory(user *model.User) {
	fmt.Println()
	fmt.Println("--- Start : Riwayat Tidur ---")
	fmt.Println("")
	for i, sleepRecord := range user.Sleeps {
		if (sleepRecord.Date != "") && (sleepRecord.SleepStart != "") && (sleepRecord.SleepEnd != "") {
			fmt.Printf("%d) \n", i+1)
			fmt.Println("Tanggal    : ", helpers.PrintFormattedDate(sleepRecord.Date))
			fmt.Println("Jam Tidur  : ", helpers.PrintFormattedHour(sleepRecord.SleepStart))
			fmt.Println("Jam Bangun : ", helpers.PrintFormattedHour(sleepRecord.SleepEnd))
			// INI benerin yakk hehe
			fmt.Printf("\nDurasi Tidur : %.f", (sleepRecord.Duration))

			fmt.Println()
		}
	}
	fmt.Println("---- End : Riwayat Tidur ----")
}
