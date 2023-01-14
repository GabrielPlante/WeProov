package database

import (
	weproovuser "WeProov/User"
)

var database []weproovuser.User

func AddUser(user weproovuser.User) {
	database = append(database, user)
}

func GetAllUsers() []weproovuser.User {
	return database
}

//Remove a user at a given index, order isn't kept
func removeUser(s []weproovuser.User, i int) []weproovuser.User {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

//Remove a user from the database, return true if a user has been deleted
func DeleteUser(username string) bool {
	for i := 0; i != len(database); i++ {
		if database[i].Username == username {
			database = removeUser(database, i)
			return true
		}
	}
	return false
}
