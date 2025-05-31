package model

type User struct {
	Username    string
	Password    string
	Sleeps      []SleepRecord
	TotalSleeps int
}
