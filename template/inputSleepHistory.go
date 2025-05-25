package template

import (
	"SleepTrack/model"
	"SleepTrack/services/inputSleepHistory"
	"fmt"
)

func PrintInputNewHistory() model.SleepRecord {
	var SleepRecord model.SleepRecord

	fmt.Println()
	fmt.Println("=============================")
	fmt.Println("Masukan waktu tidur kamu!")
	fmt.Println()

	inputSleepHistory.InputSelectedAction(&SleepRecord)

	fmt.Println()
	fmt.Println("Review data kamu")
	fmt.Println("Tanggal    : ", SleepRecord.Date)
	fmt.Println("Jam Tidur  : ", SleepRecord.SleepStart)
	fmt.Println("Jam Bangun : ", SleepRecord.SleepEnd)

	fmt.Println()
	fmt.Println("Data sudah tersimpan!")
	fmt.Println("=============================")

	return SleepRecord
}
