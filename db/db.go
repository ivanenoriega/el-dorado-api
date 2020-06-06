package db

import (
	"database/sql"
	"os"
)

func dbConn() (db *sql.DB) {
	// Local config
	// dbUser := "root"
	// dbPass := "AsD21158651!"
	// dbName := "el-dorado"
	dbDriver := "mysql"
	dbURL := os.Getenv("JAWSDB_URL")
	db, err := sql.Open(dbDriver, dbURL)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// GetConn - return the db connection
func GetConn() *sql.DB {
	return dbConn()
}
