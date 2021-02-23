package dao

import (
	"fmt"
	"log"

	"github.com/picloader/model"
	"github.com/picloader/util"
)

// ObtainAlbumID used to lookup/create album foreign key for Pic
func ObtainAlbumID(name string) (id int64) {

	db := util.DBConn()
	defer db.Close()

	if a := findAlbumByName(name); a.ID != 0 {
		id = a.ID
		return
	}

	a := model.Album{Album: name}
	id, _ = createAlbum(a)
	return
}

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

func findAlbumByName(name string) (a model.Album) {

	db := util.DBConn()
	defer db.Close()

	query := "SELECT id, album FROM album where album = ?"
	row := db.QueryRow(query, name)
	row.Scan(&a.ID, &a.Album)

	return a
}
