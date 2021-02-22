package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var user = os.Getenv("GALLERY_TESTDB_GO_USER")
var pass = os.Getenv("GALLERY_TESTDB_GO_PASSWORD")
var dbIP = os.Getenv("GALLERY_TESTDB_GO_IP")
var name = os.Getenv("GALLERY_TESTDB_GO_DBNAME")

var url = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, dbIP, name)

// DBConn is db connector function
func DBConn() *sql.DB {
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("Cannot connect to database: %s/%s\n", dbIP, name)
		log.Fatal("Database connection error: ", err)
	} else {
		fmt.Printf("Connected to: %s/%s\n", dbIP, name)
	}
	return db
}
