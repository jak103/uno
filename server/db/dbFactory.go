package db

import (
	"os"
	"strings"
)

// GetDb factory method to get the Database connection
func GetDb() *UnoDB {
	switch environment := strings.ToUpper(os.Getenv("DB_TYPE")); environment {
	case "MOCK":
		return newMockDb()
	case "MONGO":
		return newMongoDb()
	case "FIREBASE":
		return newFirebaseDB()
	}
}
