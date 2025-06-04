package model

type User struct {
	Username    string
	Password    string
	Sleeps      [MaxSleep]SleepRecord
	TotalSleeps int
}
