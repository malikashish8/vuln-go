package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func getLicenseNames(w http.ResponseWriter, r *http.Request) {
	h := httpHelper{w, r}
	filename, hasFilename := h.GetQueryParam("filename")
	if hasFilename {
		getLicenseText(h, filename)
	} else {
		files, err := ioutil.ReadDir(licenseFolder)
		if err != nil {
			log.Fatal(err)
		}
		var fileNames []string
		for _, f := range files {
			fileNames = append(fileNames, f.Name())
		}
		h.WriteJson(map[string]interface{}{"files": fileNames})
	}
}

func getLicenseText(h httpHelper, filename string) {
	// vulnerability - LFI
	filePath := path.Join(licenseFolder, filename)
	log.Println("filePath: " + filePath)
	err := h.WriteFile(filePath)
	if err != nil {
		h.WriteError(httpError{http.StatusBadRequest, err.Error()})
	}
}
