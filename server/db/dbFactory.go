package db

import (
	"fmt"
	"os"
	"strings"
)

// Registry of initialized DB COnnections
var dbRegistry = make(map[string]*DB)

var connectedDB *DB

// registerDB is designed to make a singleton registry for consistent access to the database
func registerDB(d *DB) {
	dbRegistry[d.name] = d
}

// DB is a struct wrapper for UnoDB that allows it to be registered. It also provides a consistent return type for the user
type DB struct {
	name        string
	description string
	UnoDB       // By not declaring a local variable, DB automatically assumes the functions of this instantiated member
}

// GetDb factory method to get the Database connection
func GetDb() (*DB, error) {
	if connectedDB == nil {
		environment := strings.ToUpper(os.Getenv("DB_TYPE"))

		if environment == "" {
			environment = "MOCK"
		}

		if db, ok := dbRegistry[environment]; ok {
			db.connect()
			connectedDB = db
		} else {
			return nil, fmt.Errorf("%s is not a registered DB_TYPE", environment)
		}
	}

	return connectedDB, nil
}

// Disconnect disconnects the database
func Disconnect() {
	if connectedDB != nil {
		connectedDB.disconnect()
	}
}
