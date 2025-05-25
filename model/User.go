package model

type User struct {
	username string
	password string
}

type Users struct {
	users  []User
	sleeps []SleepRecord
}
