package app

import "os"

func IsAdmin(userID string) bool {
	adminID := os.Getenv("ADMIN_USER_ID")
	return userID == adminID
}
