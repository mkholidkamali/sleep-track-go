package template

import (
	"SleepTrack/helpers"
	"SleepTrack/model"
	"SleepTrack/services/user"
	"fmt"
)

func PrintIntroduction(action *int, users *[model.MaxUser]model.User, nUsers *int) int {
	var loggedUserIndex int
	var isSuccess bool

	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Aplikasi Pemantau Tidur Dan Kesehatan")
	fmt.Println("=============================")

	for !isSuccess {
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
			if loggedUserIndex != -1 {
				isSuccess = true
			}
		case 2:
			PrintRegister(users, nUsers)
			*action = 1
		case 3:
			helpers.PrintExitMessage()
			isSuccess = true
		default:
			helpers.PrintErrorMessage()
		}
	}

	return loggedUserIndex
}

func PrintLogin(users *[model.MaxUser]model.User) int {
	var tempUser model.User
	var loggedUserId int
	var isSuccess bool

	for loggedUserId != -1 && !isSuccess {
		fmt.Println("")
		fmt.Println("=============================")
		fmt.Println("Login Page")
		fmt.Println("=============================")
		fmt.Println("Masukan username kamu")
		fmt.Print("> ")
		fmt.Scan(&tempUser.Username)
		fmt.Println("Masukan password kamu")
		fmt.Print("> ")
		fmt.Scan(&tempUser.Password)

		loggedUserId = user.AttemptLogin(&tempUser, users)

		if loggedUserId == -1 {
			helpers.PrintFailedLoginMessage()
		} else {
			isSuccess = true
		}
	}

	fmt.Println("=============================")

	return loggedUserId
}

func PrintRegister(users *[model.MaxUser]model.User, nUsers *int) bool {
	var tempUser model.User
	var isRegistered bool

	isCancel := 1

	for !isRegistered && isCancel == 1 {
		fmt.Println("")
		fmt.Println("=============================")
		fmt.Println("Register Page")
		fmt.Println("=============================")
		fmt.Println("Masukan username kamu")
		fmt.Print("> ")
		fmt.Scan(&tempUser.Username)
		fmt.Println("Masukan password kamu")
		fmt.Print("> ")
		fmt.Scan(&tempUser.Password)

		isRegistered = user.RegisterUser(&tempUser, users, nUsers)

		if !isRegistered {
			helpers.PrintFailedRegisterMessage()

			fmt.Println("=============================")
			fmt.Println("Lanjut mendaftar?")
			fmt.Print("> ")
			fmt.Println("1. Iya")
			fmt.Println("2. Tidak")
			fmt.Scan(&isCancel)
		} else {
			*nUsers += 1
			fmt.Println("")
			fmt.Println("User berhasil diregistrasi!")
		}
	}

	fmt.Println("=============================")

	return isRegistered
}
