package showSleepHistory

import (
	"SleepTrack/helpers"
	"SleepTrack/model"
	"fmt"
)

func InputSelectedAction(action *int) {
	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Sleep History")
	fmt.Println("=============================")
	fmt.Println("Apa yang ingin kamu lakukan")
	fmt.Println("1. Ubah urutan")
	fmt.Println("2. Cari dan edit jadwal tidur")
	fmt.Println("3. Laporan hasil jadwal 7 hari terakhir")
	fmt.Println("4. Exit")
	fmt.Print("> ")
	fmt.Scan(action)
	fmt.Println("=============================")
}

func InputSortAction(action *int, user *model.User, sortedBy *string, sortedType *string) {
	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Silahkan pilih urutan")
	fmt.Println("1. Berdasarkan Durasi (Ascending)")
	fmt.Println("2. Berdasarkan Tanggal (Descending)")
	fmt.Print("> ")
	fmt.Scan(action)
	fmt.Println("=============================")

	if *action == 1 {
		SortbyDuration(&user.Sleeps, user.TotalSleeps)
		*sortedBy = "Duration"
		*sortedType = "Ascending"
	} else if *action == 2 {
		SortbyDate(&user.Sleeps, user.TotalSleeps)
		*sortedBy = "Date"
		*sortedType = "Descending"
	} else {
		helpers.PrintErrorMessage()
		InputSortAction(action, user, sortedBy, sortedType)
	}
}

func InputSearcAndEditAction(action *int, user *model.User) {
	var searchedDate string
	*action = 0

	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Cari jadwal tidur kamu berdasar tanggal")
	fmt.Print("> ")
	fmt.Scan(&searchedDate)
	fmt.Println("=============================")

	searchedIndex := SearchByDate(searchedDate, &user.Sleeps, user.TotalSleeps)

	if searchedIndex != -1 {
		for *action != 2 {
			fmt.Println("")
			fmt.Println("=============================")
			fmt.Println("Jadwal Tidur Kamu")
			fmt.Printf("%d) \n", searchedIndex+1)
			fmt.Println("Tanggal      :", helpers.PrintFormattedDate(user.Sleeps[searchedIndex].Date))
			fmt.Println("Jam Tidur    :", helpers.PrintFormattedHour(user.Sleeps[searchedIndex].SleepStart))
			fmt.Println("Jam Bangun   :", helpers.PrintFormattedHour(user.Sleeps[searchedIndex].SleepEnd))
			fmt.Printf("Durasi Tidur : %.f\n\n", (user.Sleeps[searchedIndex].Duration))

			fmt.Println("Apa yang mau kamu lakukan?")
			fmt.Println("1. Edit data")
			fmt.Println("2. Kembali")
			fmt.Print("> ")
			fmt.Scan(action)

			if *action == 1 {
				EditSleepHistory(action, &user.Sleeps[searchedIndex])
			}
		}
	} else {
		helpers.PrintFailedSearchMessage()
	}
}

// Algorithm Sort : Insertion
func SortbyDuration(sleeps *[model.MaxSleep]model.SleepRecord, totalSleep int) {
	for i := 1; i < totalSleep; i++ {
		key := sleeps[i]
		j := i - 1

		for j >= 0 && sleeps[j].Duration > key.Duration {
			sleeps[j+1] = sleeps[j]
			j--
		}
		sleeps[j+1] = key
	}
}

// Algorithm Sort : Selection
func SortbyDate(sleeps *[model.MaxSleep]model.SleepRecord, totalSleep int) {
	for i := 0; i < totalSleep-1; i++ {
		maxIdx := i
		for j := i + 1; j < totalSleep; j++ {
			if sleeps[j].Date > sleeps[maxIdx].Date {
				maxIdx = j
			}
		}

		if maxIdx != i {
			sleeps[i], sleeps[maxIdx] = sleeps[maxIdx], sleeps[i]
		}
	}
}

// Algorithm Search: Binary
func SearchByDate(searchedDate string, sleeps *[model.MaxSleep]model.SleepRecord, totalSleep int) int {
	// Initialize
	low := 0
	high := totalSleep - 1

	// Must be sorted first
	SortbyDate(sleeps, totalSleep)

	// Sort
	for low <= high {
		mid := low + (high-low)/2

		if sleeps[mid].Date == searchedDate {
			return mid
		}

		if sleeps[mid].Date < searchedDate {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func EditSleepHistory(action *int, sleep *model.SleepRecord) {
	var editedCategory, before, changed string

	for *action != 4 {
		fmt.Println("")
		fmt.Println("=============================")
		fmt.Println("Apa yang mau kamu edit?")
		fmt.Println("1. Tanggal")
		fmt.Println("2. Jam Tidur")
		fmt.Println("3. Jam Bangun")
		fmt.Println("4. Kembali")
		fmt.Print("> ")
		fmt.Scan(action)
		fmt.Println("")

		if *action == 1 {
			editedCategory = "Tanggal"
			before = helpers.PrintFormattedDate(sleep.Date)
		} else if *action == 2 {
			editedCategory = "Jam Tidur"
			before = helpers.PrintFormattedHour(sleep.SleepStart)
		} else if *action == 3 {
			editedCategory = "Jam Bangun"
			before = helpers.PrintFormattedHour(sleep.SleepEnd)
		} else if *action == 4 {
			return
		} else {
			helpers.PrintErrorMessage()
			EditSleepHistory(action, sleep)
			return
		}

		fmt.Println("Silahkan masukan data", editedCategory)
		fmt.Println("Sebelum :", before)
		fmt.Print("Sesudah : ")
		fmt.Scan(&changed)

		if *action == 1 {
			sleep.Date = changed
		} else if *action == 2 {
			sleep.SleepStart = changed
		} else if *action == 3 {
			sleep.SleepEnd = changed
		}
	}

	fmt.Println("> Data berhasil diupdate!")
	fmt.Println("=============================")
}

func CalculateTotalSleepInWeek(sleepRecord *[model.MaxSleep]model.SleepRecord) int {
	totalSleep := CountTotalSleep(sleepRecord)

	startIndex := 0
	if totalSleep > 7 {
		startIndex = totalSleep - 7
	}

	totalDays := 0
	lastDate := ""

	for i := startIndex; i < totalSleep; i++ {
		if sleepRecord[i].Date != lastDate {
			totalDays++
			lastDate = sleepRecord[i].Date
		}
	}

	return totalDays
}

func CalculateTotalSleepHour(sleepRecord *[model.MaxSleep]model.SleepRecord) float64 {
	totalSleep := CountTotalSleep(sleepRecord)

	startIndex := 0
	if totalSleep > 7 {
		startIndex = totalSleep - 7
	}

	var totalHours float64

	for i := startIndex; i < totalSleep; i++ {
		totalHours += sleepRecord[i].Duration
	}

	return totalHours
}

func CalculateAverageSleepHour(sleepRecord *[model.MaxSleep]model.SleepRecord) float64 {
	totalDays := CalculateTotalSleepInWeek(sleepRecord)
	if totalDays == 0 {
		return 0
	}

	totalHours := CalculateTotalSleepHour(sleepRecord)
	return totalHours / float64(totalDays)
}

func CountTotalSleep(sleeps *[model.MaxSleep]model.SleepRecord) int {
	totalSleep := 0
	for i := 0; i < 7; i++ {
		if sleeps[i].Date != "" {
			totalSleep++
		}
	}
	return totalSleep
}
