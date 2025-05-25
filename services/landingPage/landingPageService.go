package landingPage

import (
	"SleepTrack/helpers"
	"SleepTrack/model"
	"fmt"
)

func InputSelectedAction(action *int, user *model.User) {
	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Apa yang ingin kamu lakukan")
	fmt.Println("1. Masukan waktu tidur kamu")
	fmt.Println("2. Lihat jadwal tidur")
	fmt.Println("3. Tentang aplikasi ini")
	fmt.Println("4. Exit")
	fmt.Println()
	fmt.Print("> ")
	fmt.Scan(action)
	fmt.Println("=============================")

	if *action < 1 || *action > 4 {
		helpers.PrintErrorMessage()
		InputSelectedAction(action, user)
	}
}
