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
	dbConn, err = sql.Open("mysql", connectionString)
	log.Println("connectionString " + connectionString)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
}

func GetConnection() *sql.DB {
	return dbConn
}
