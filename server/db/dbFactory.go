package db

import (
	"os"
	"strings"
)

func GetDb() *UnoDB {
	switch enviornment := strings.ToUpper(os.Getenv("DB_TYPE")); environment {
	case "MOCK":
		return newMockDb()
	case "MONGO":
		return newMongoDb()
	case "FIREBASE":
		return newFirebaseDb() // TODO convert to prod DB connection
	}
}
