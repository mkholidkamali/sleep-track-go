package template

import "fmt"

func PrintLandingPage(action *int) {
	fmt.Println()
	fmt.Println("=============================")
	fmt.Println()
	fmt.Println("Zzz...")
	fmt.Println("Hoam... Halo selamat pagi !!")
	fmt.Println("Gimana tidur kamu hari ini?")
	fmt.Println("")
	fmt.Println("Apa yang ingin kamu lakukan")
	fmt.Println("1. Masukan waktu tidur kamu")
	fmt.Println("2. Lihat jadwal tidur")
	fmt.Println("3. Tentang aplikasi ini")
	fmt.Println("4. Exit")
	fmt.Println()
	fmt.Print("> ")
	fmt.Scan(action)
	fmt.Println("=============================")
}
