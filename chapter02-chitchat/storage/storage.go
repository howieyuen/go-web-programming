package storage

import (
	"database/sql"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Db, err = sql.Open("postgres", psqlInfo)
	Db, err = sql.Open("mysql", "root:!QAZ2wsx@tcp(127.0.0.1:3306)/chitchat?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	
	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Successfully connected!")
}
