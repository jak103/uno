package db

import (
	"os"
	"strings"
)

// GetDb factory method to get the Database connection
func GetDb() UnoDB {
	switch environment := strings.ToUpper(os.Getenv("DB_TYPE")); environment {
	case "MOCK":
		return newMockDB()
	case "MONGO":
		return newMongoDB()
	case "FIREBASE":
		return newFirebaseDB()
	}
	return newMockDB()
}
