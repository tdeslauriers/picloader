package dao

import (
	"fmt"
	"log"

	"github.com/picloader/model"
	"github.com/picloader/util"
)

func createAlbum(album model.Album) (id int64, errSQL error) {

	db := util.DBConn()
	defer db.Close()

	query := "INSERT INTO album (album) VALUES (?);"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	r, errSQL := stmt.Exec(album.Album)
	if errSQL != nil {
		log.Fatal(errSQL)
	}

	id, errID := r.LastInsertId()
	if errID != nil {
		log.Fatal(errID)
	}

	fmt.Printf("Created record id: %d\n", id)
	db.Close()

	return id, errSQL
}
