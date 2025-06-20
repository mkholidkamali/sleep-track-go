package seeders

import (
	"SleepTrack/model"
)

func SeedUsers(nUser *int) [model.MaxUser]model.User {
	// Update total users
	*nUser = 2

	// Return users
	return [model.MaxUser]model.User{
		{
			Username: "andi",
			Password: "123",
			Sleeps: [model.MaxSleep]model.SleepRecord{
				{Date: "220525", SleepStart: "2000", SleepEnd: "0500", Duration: 9},
				{Date: "230525", SleepStart: "2200", SleepEnd: "0500", Duration: 7},
			},
			TotalSleeps: 2,
		},
		{
			Username: "budi",
			Password: "456",
			Sleeps: [model.MaxSleep]model.SleepRecord{
				{Date: "200525", SleepStart: "2330", SleepEnd: "0500", Duration: 5.5},
				{Date: "210525", SleepStart: "2100", SleepEnd: "0500", Duration: 8},
			},
			TotalSleeps: 2,
		},
	}
}
