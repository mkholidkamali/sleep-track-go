package user

import (
	"SleepTrack/model"
)

func AttemptLogin(attemptUser *model.User, users *[model.MaxUser]model.User) int {
	// Check user availability
	userIndex := searchUser(*attemptUser, *users, true)

	return userIndex
}

func RegisterUser(attemptUser *model.User, users *[model.MaxUser]model.User, nUser *int) bool {
	// Check user availability
	userIndex := searchUser(*attemptUser, *users, false)
	if userIndex != -1 {
		return false
	}

	// Create new user
	users[*nUser] = *attemptUser

	// Append total users
	*nUser += 1

	return true
}

/**
* Private Method
* Algorithm : Sequential
 */
func searchUser(attemptUser model.User, users [model.MaxUser]model.User, isCheckPassword bool) int {
	// Initialize user index
	userIndex := -1

	// Initalize variable
	isUsernameFound := false
	isPasswordCorrect := true

	// Search throught all users
	for i, user := range users {
		// Check username
		isUsernameFound = attemptUser.Username == user.Username

		// Check password
		if isCheckPassword {
			isPasswordCorrect = attemptUser.Password == user.Password
		}

		// Set user index
		if isUsernameFound && isPasswordCorrect && userIndex == -1 {
			userIndex = i
		}
	}

	return userIndex
}
