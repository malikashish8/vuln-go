package main

import (
	"errors"
	"net/http"
)

func getAlbums(h httpHelper) {
	id, idExists := h.GetQueryParam("id")
	if idExists {
		getAlbumByID(h, id)
		return
	} else {
		albumsFromDB := readAlbumsFromDB()
		h.WriteJson(albumsFromDB)
	}
}

func getAlbumByID(h httpHelper, id string) {
	result, err := readByIDFromDB(id)
	if err != nil {
		h.WriteError(httpError{http.StatusNotFound, "album not found"})
	} else {
		h.WriteJson(result)
	}
}

func postAlbums(h httpHelper) {
	var newAlbum album
	var err error
	err = h.BindJSON(&newAlbum)
	if err == nil {
		success := writeAlbumToDB(newAlbum)
		if success {
			h.WriteJson(newAlbum)
		} else {
			err = errors.New("write failed")
		}
	}
	if err != nil {
		h.WriteError(httpError{http.StatusInternalServerError, err.Error()})
	}
}
