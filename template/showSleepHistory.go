package template

import (
	"SleepTrack/helpers"
	"SleepTrack/model"
	"SleepTrack/services/showSleepHistory"
	"fmt"
)

func PrintShowSleepHistory(action *int, user *model.User) {
	sortedBy := "-"
	sortedType := "-"

	// Show all data
	showAllHistory(user, &sortedBy, &sortedType)

	for *action != 5 {
		// Input selected action
		showSleepHistory.InputSelectedAction(action)

		switch *action {
		case 1:
			showSleepHistory.InputSortAction(action, user, &sortedBy, &sortedType)
			showAllHistory(user, &sortedBy, &sortedType)
		case 2:
			showSleepHistory.InputSearcAndEditAction(action, user)
		case 3:
			showSleepHistory.LastWeekSleeps(&user.Sleeps, user.TotalSleeps)
		case 4:
			//
		case 5:
			//
		default:
			helpers.PrintErrorMessage()
		}
	}
}

/**
* Private Method
**/
func showAllHistory(user *model.User, sortedBy *string, sortedType *string) {
	fmt.Println()
	fmt.Println("--- Start : Riwayat Tidur ---")
	fmt.Println("")

	fmt.Println("Urutan")
	fmt.Println("Berdasar :", *sortedBy)
	fmt.Println("Secara   :", *sortedType)
	fmt.Println("")

	for i, sleepRecord := range user.Sleeps {
		if (sleepRecord.Date != "") && (sleepRecord.SleepStart != "") && (sleepRecord.SleepEnd != "") {
			fmt.Printf("%d) \n", i+1)
			fmt.Println("Tanggal      :", helpers.PrintFormattedDate(sleepRecord.Date))
			fmt.Println("Jam Tidur    :", helpers.PrintFormattedHour(sleepRecord.SleepStart))
			fmt.Println("Jam Bangun   :", helpers.PrintFormattedHour(sleepRecord.SleepEnd))
			fmt.Printf("Durasi Tidur : %.f\n\n", (sleepRecord.Duration))
		}
	}
	fmt.Println("---- End : Riwayat Tidur ----")
}
