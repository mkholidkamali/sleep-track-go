package inputSleepHistory

import (
	"SleepTrack/helpers"
	"SleepTrack/model"
	"SleepTrack/services/showSleepHistory"
	"fmt"
	"strconv"
)

func InputSelectedAction(sleepRecord *model.SleepRecord, existingSleeps *[model.MaxSleep]model.SleepRecord, totalSleep int) {
	fmt.Println("--- Start : Input Data ---")
	fmt.Println("")

	fmt.Println("Tanggal (Format : ddmmyy, contoh : 230525)")
	fmt.Print("> ")
	fmt.Scan(&sleepRecord.Date)
	if !ValidateDate(sleepRecord.Date) {
		helpers.PrintErrorMessage()
		InputSelectedAction(sleepRecord, existingSleeps, totalSleep)
		return
	}
	if !ValidateExistingDate(sleepRecord.Date, existingSleeps, totalSleep) {
		helpers.PrintExistMessage()
		InputSelectedAction(sleepRecord, existingSleeps, totalSleep)
		return
	}

	fmt.Println()
	fmt.Println("Jam Tidur (Format : 24 Jam, contoh : 2200)")
	fmt.Print("> ")
	fmt.Scan(&sleepRecord.SleepStart)
	if !ValidateTime(sleepRecord.SleepStart) {
		helpers.PrintErrorMessage()
		InputSelectedAction(sleepRecord, existingSleeps, totalSleep)
		return
	}

	fmt.Println()
	fmt.Println("Jam Bangun (Format : 24 Jam, contoh : 0500)")
	fmt.Print("> ")
	fmt.Scan(&sleepRecord.SleepEnd)
	if !ValidateTime(sleepRecord.SleepEnd) {
		helpers.PrintErrorMessage()
		InputSelectedAction(sleepRecord, existingSleeps, totalSleep)
		return
	}

	fmt.Println("")
	fmt.Println("--- End : Input Data ---")

	calculateTotalSleepHour(sleepRecord)
}

func ValidateDate(dateString string) bool {
	if len(dateString) != 6 {
		return false
	}

	dateInt, err := strconv.Atoi(dateString)
	if err != nil {
		return false
	}

	dd := helpers.ExtractDay(dateInt)
	mm := helpers.ExtractMonth(dateInt)
	yy := dateInt % 100

	if dd < 1 || dd > 30 {
		return false
	}

	if mm < 1 || mm > 12 {
		return false
	}

	if yy < 1 || yy > 25 {
		return false
	}

	return true
}

func ValidateExistingDate(dateString string, sleeps *[model.MaxSleep]model.SleepRecord, totalSleep int) bool {
	return showSleepHistory.SearchByDate(dateString, sleeps, totalSleep) == -1
}

func ValidateTime(timeString string) bool {
	if len(timeString) != 4 {
		return false
	}

	timeInt, err := strconv.Atoi(timeString)
	if err != nil {
		return false
	}

	hh := timeInt / 100
	mm := timeInt % 100

	if hh < 0 || hh > 23 {
		return false
	}

	if mm < 0 || mm > 59 {
		return false
	}

	return true
}

func calculateTotalSleepHour(sleepRecord *model.SleepRecord) {
	startInt, _ := strconv.Atoi(sleepRecord.SleepStart)
	endInt, _ := strconv.Atoi(sleepRecord.SleepEnd)

	startHour := startInt / 100
	startMinute := startInt % 100

	endHour := endInt / 100
	endMinute := endInt % 100

	startTotalMinutes := startHour*60 + startMinute
	endTotalMinutes := endHour*60 + endMinute

	if endTotalMinutes <= startTotalMinutes {
		endTotalMinutes += 24 * 60
	}

	durationMinutes := endTotalMinutes - startTotalMinutes
	durationHours := float64(durationMinutes) / 60.0

	sleepRecord.Duration = durationHours
}
