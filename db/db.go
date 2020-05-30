package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "AsD21158651!"
	dbName := "el-dorado"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// GetConn - return the db connection
func GetConn() *sql.DB {
	return dbConn()
}

// Hello - Friendly greeting
func Hello() string {
	return "Hello bitch!"
}
