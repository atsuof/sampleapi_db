package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbName := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)

	db, dbOpenError := sql.Open("mysql", dbConn)
	if dbOpenError != nil {
		fmt.Println("error occurred", dbOpenError)
	}
	defer db.Close()

	const sqlStr = `select title, contents, username, nice　from articles;`

	rows, err := db.Query(sqlStr)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	articleArray = []Article{}
}
