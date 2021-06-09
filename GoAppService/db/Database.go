package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB

func SetupDatabase() {
	var err error
	connectionString := fmt.Sprintf("%s:%s@tcp(remotemysql.com:3306)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_USER"))
	log.Println("connectionString ", connectionString)
	dbConn, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *sql.DB {
	return dbConn
}
