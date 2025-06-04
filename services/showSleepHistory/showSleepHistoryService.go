package showSleepHistory

import (
	"SleepTrack/helpers"
	"SleepTrack/model"
	"SleepTrack/services/landingPage"
	"fmt"
)

func InputSelectedAction(action *int, user *model.User) {
	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Sleep History")
	fmt.Println("=============================")
	fmt.Println("Apa yang ingin kamu lakukan")
	fmt.Println("1. Ubah urutan")
	fmt.Println("2. Cari dan edit jadwal tidur")
	fmt.Println("3. Laporan hasil jadwal 7 hari terakhir")
	fmt.Println("4. Rekomendasi jadwal tidur")
	fmt.Println("5. Exit")
	fmt.Print("> ")
	fmt.Scan(action)
	fmt.Println("=============================")

	if *action == 1 {
		landingPage.InputSortAction(action, user)
		// panggil input sort
	} else if *action == 2 {
		// panggil search
	} else if *action == 3 {
		lastWeekSleeps(&user.Sleeps, user.TotalSleeps)

	} else if *action == 4 {
		// apan enih?
	} else if *action == 5 {
		// exit
	} else {
		helpers.PrintErrorMessage()
		InputSelectedAction(action, user)
	}

	// if *action < 1 || *action > 5 {
	// 	helpers.PrintErrorMessage()
	// 	InputSelectedAction(action)
	// }
}

func lastWeekSleeps(sleeps *[model.MaxSleep]model.SleepRecord, total int) {
	// ntar di sort pake tanggal

	var averageDuration float64

	if total < 7 {
		total = 7
	}
	for i := 0; i < total; i++ {
		averageDuration += sleeps[i].Duration
	}

	// Averagenya harus di adjut lagi supaya nilai di blkg koma (0,) disesuain sm format menit
	fmt.Printf("%.2f", averageDuration/float64(total))
}
