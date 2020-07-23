package db

import (
	"fmt"
	"os"
	"strings"
)

// GetDb factory method to get the Database connection
func GetDb() UnoDB {
	fmt.Println("Connecting to the database...")
	switch environment := strings.ToUpper(os.Getenv("DB_TYPE")); environment {
	case "MOCK":
		fmt.Println("Using Mock database.")
		return newMockDB()
	case "MONGO":
		fmt.Println("Using Mongo database.")
		return newMongoDB()
	case "FIREBASE":
		fmt.Println("Using Firebase database.")
		return newFirebaseDB()
	}
	return newMockDB()
}
