package inputSleepHistory

import (
	"SleepTrack/model"
	"fmt"
)

func InputSelectedAction(sleepRecord *model.SleepRecord) {
	// for

	fmt.Println("--- Start : Input Data ---")
	fmt.Println("Tanggal (Format : ddmmyy, contoh : 230525)")
	fmt.Print("> ")
	fmt.Scan(&sleepRecord.Date)
	// !TODO : Validate

	fmt.Println()
	fmt.Println("Jam Tidur (Format : 24 Jam, contoh : 2200)")
	fmt.Print("> ")
	fmt.Scan(&sleepRecord.SleepStart)
	// !TODO : Validate

	fmt.Println()
	fmt.Println("Jam Bangun (Format : 24 Jam, contoh : 0500)")
	fmt.Print("> ")
	fmt.Scan(&sleepRecord.SleepEnd)
	// !TODO : Validate
	fmt.Println("--- End : Input Data ---")

	calculateTotalSleepHour(sleepRecord)
}

/**
* Private Method
**/
func calculateTotalSleepHour(sleepRecord *model.SleepRecord) {
	// !TODO : Calculate here

	sleepRecord.Duration = 8
}

// func validate() {
// 	//
// }
