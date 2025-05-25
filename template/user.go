package template

import (
	"SleepTrack/helpers"
	"SleepTrack/model"
	"SleepTrack/services/user"
	"fmt"
)

func PrintIntroduction(action *int, users *[]model.User) int {
	var loggedUserIndex int
	var isSuccess bool

	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Aplikasi Pemantau Tidur Dan Kesehatan")
	fmt.Println("=============================")

	for loggedUserIndex != -1 && !isSuccess {
		fmt.Println("")
		fmt.Println("Apakah kamu sudah punya akun?")
		fmt.Println("1. Sudah")
		fmt.Println("2. Belum, aku ingin mendaftar")
		fmt.Println("3. Keluar")
		fmt.Print("> ")
		fmt.Scan(action)

		switch *action {
		case 1:
			loggedUserIndex = PrintLogin(users)
			isSuccess = true
		case 2:
			PrintRegister()
			loggedUserIndex = PrintLogin(users)
			isSuccess = true
		case 3:
			helpers.PrintExitMessage()
			loggedUserIndex = -1
		default:
			helpers.PrintErrorMessage()
			loggedUserIndex = -1
		}
	}

	return loggedUserIndex
}

func PrintLogin(users *[]model.User) int {
	var tempUser model.User
	var loggedUserId int
	var isSuccess bool

	for loggedUserId != -1 && !isSuccess {
		fmt.Println("")
		fmt.Println("=============================")
		fmt.Println("Masukan username kamu")
		fmt.Print("> ")
		fmt.Scan(&tempUser.Username)
		fmt.Println("Masukan password kamu")
		fmt.Print("> ")
		fmt.Scan(&tempUser.Password)

		loggedUserId = user.AttemptLogin(&tempUser, []model.User{})

		if loggedUserId == -1 {
			helpers.PrintFailedLoginMessage()
		} else {
			isSuccess = true
		}
	}

	fmt.Println("=============================")

	return loggedUserId
}

func PrintRegister() int {
	var tempUser model.User
	var loggedUserId int
	var isSuccess bool

	for loggedUserId != -1 && !isSuccess {
		fmt.Println("")
		fmt.Println("=============================")
		fmt.Println("Masukan username kamu")
		fmt.Print("> ")
		fmt.Scan(&tempUser.Username)
		fmt.Println("Masukan password kamu")
		fmt.Print("> ")
		fmt.Scan(&tempUser.Password)

		loggedUserId = user.RegisterUser(&tempUser, []model.User{})

		if loggedUserId == -1 {
			helpers.PrintFailedRegisterMessage()
		} else {
			isSuccess = true
		}
	}

	fmt.Println("=============================")

	return loggedUserId
}
