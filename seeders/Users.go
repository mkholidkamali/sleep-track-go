package seeders

import (
	"SleepTrack/model"
)

func SeedUsers() []model.User {
	return []model.User{
		{
			Username: "andi",
			Password: "123",
			Sleeps: []model.SleepRecord{
				{Date: 220525, SleepStart: 2200, SleepEnd: 0500, Duration: 7},
				{Date: 230525, SleepStart: 2200, SleepEnd: 0500, Duration: 8},
			},
		},
		{
			Username: "budi",
			Password: "456",
			Sleeps: []model.SleepRecord{
				{Date: 200525, SleepStart: 2330, SleepEnd: 0500, Duration: 5.5},
				{Date: 210525, SleepStart: 2100, SleepEnd: 0500, Duration: 8},
			},
		},
	}
}
