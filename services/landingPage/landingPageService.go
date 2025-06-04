package landingPage

import (
	"SleepTrack/helpers"
	"SleepTrack/model"

	// "SleepTrack/services/inputSleepHistory"
	"fmt"
)

func InputSelectedAction(action *int, user *model.User) {
	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Landing Page")
	fmt.Println("=============================")
	fmt.Println("Apa yang ingin kamu lakukan")
	fmt.Println("1. Masukan waktu tidur kamu")
	fmt.Println("2. Lihat jadwal tidur")
	fmt.Println("3. Tentang aplikasi ini")
	fmt.Println("4. Exit")
	fmt.Print("> ")
	fmt.Print(user)
	fmt.Scan(action)
	fmt.Println("=============================")

	if *action < 1 || *action > 4 {
		helpers.PrintErrorMessage()
		InputSelectedAction(action, user)
	}
}

func InputSortAction(action *int, user *model.User) {

	// fmt.Print(user.TotalSleeps)

	// Sleeps := user.Sleeps

	fmt.Println("")
	// !TODO : Buat fungsi untuk memilih urutan
	fmt.Println("> Urutan sekarang : ")
	// fmt.Print("> ")
	fmt.Println("=============================")
	fmt.Println("=== Silahkan pilih urutan ===")
	fmt.Println("1. Berdasarkan Durasi ") //asc - insertion
	fmt.Println("2. Berdasarkan Tanggal") // desc - selection
	fmt.Println("=============================")
	fmt.Scan(action)

	if *action == 1 {
		sortbyDuration(&user.Sleeps, user.TotalSleeps)
		fmt.Print(user.Sleeps)
		// fmt.Print(user.Sleeps)
	} else if *action == 2 {

		// panggil sort by tanggal ye
	} else {
		helpers.PrintErrorMessage()
		// InputSortAction(action)
	}
}

// Enih pake insertion ya
func sortbyDuration(sleeps *[model.MaxSleep]model.SleepRecord, total int) {
	// fmt.Print(sleeps)
	fmt.Print(total)
	for i := 0; i < total; i++ {
		fmt.Print(sleeps[i])
		key := sleeps[i]
		j := i - 1

		for j >= 0 && sleeps[j].Duration > key.Duration {
			sleeps[j+1] = sleeps[j]
			j--
		}
		sleeps[j+1] = key
	}

}
