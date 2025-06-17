package helpers

import "fmt"

func PrintErrorMessage() {
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
	fmt.Println("Input anda tidak sesuai! Harap Masukan lagi")
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
}

func PrintFailedLoginMessage() {
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
	fmt.Println("Username atau password anda salah!")
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
}

func PrintFailedRegisterMessage() {
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
	fmt.Println("Username sudah digunakan!")
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
}

func PrintFailedSearchMessage() {
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
	fmt.Println("Data tidak ditemukan!")
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
}

func PrintExistMessage() {
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
	fmt.Println("Tanggal yang diinput sudah ada!")
	fmt.Println()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
}
