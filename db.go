package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	// vulnerability - hardcoded secret - db password
	db, err = sql.Open("mysql", "user:password@tcp(db)/vulngo")

	if err != nil {
		panic(err.Error())
	}
}

func readAlbumsFromDB() []album {
	var albums []album
	rows, err := db.Query("SELECT * FROM ALBUM")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var album album
		rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		albums = append(albums, album)
	}
	return albums
}

func writeAlbumToDB(album album) (status bool) {
	// vulnerability - sqli
	query := fmt.Sprintf("INSERT INTO ALBUM(title, artist, price) VALUES('%s','%s',%f)", album.Title, album.Artist, album.Price)
	fmt.Println("query = " + query)
	_, err := db.Exec(query)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func readByIDFromDB(id string) (album album, err error) {
	// vulnerability - sqli
	query := fmt.Sprintf("SELECT * from ALBUM WHERE ID = %v", id)
	fmt.Println("query = " + query)
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err = errors.New("no album for this id")
		fmt.Println(err.Error())
		return
	}

	rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	return

}
