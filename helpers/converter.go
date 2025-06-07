package helpers

import (
	"fmt"
	"strconv"
)

func ExtractDay(dateInt int) int {
	return dateInt / 10000
}

func ExtractMonth(dateInt int) int {
	return (dateInt / 100) % 100
}

func ExtractYear(dateInt int) int {
	return 2000 + (dateInt % 100)
}

func PrintFormattedDate(dateString string) string {
	// Convert into int
	dateInt, _ := strconv.Atoi(dateString)

	// Explode date
	dd := ExtractDay(dateInt)
	mm := ExtractMonth(dateInt)
	tahun := ExtractYear(dateInt)

	// Ambil nama bulan
	namaBulan := getMonth(mm)

	return fmt.Sprintf("%02d-%s-%d", dd, namaBulan, tahun)
}

func PrintFormattedHour(hourString string) string {
	// Convert into int
	hourInt, _ := strconv.Atoi(hourString)

	// Explode hour
	hh := hourInt / 100
	mm := hourInt % 100

	return fmt.Sprintf("%02d.%02d", hh, mm)
}

/**
* Private Method
**/
func getMonth(mm int) string {
	switch mm {
	case 1:
		return "Jan"
	case 2:
		return "Feb"
	case 3:
		return "Mar"
	case 4:
		return "Apr"
	case 5:
		return "May"
	case 6:
		return "Jun"
	case 7:
		return "Jul"
	case 8:
		return "Aug"
	case 9:
		return "Sep"
	case 10:
		return "Oct"
	case 11:
		return "Nov"
	case 12:
		return "Dec"
	default:
		return ""
	}
}
