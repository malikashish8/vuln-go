package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/gin-gonic/gin"
)

var licenseFolder string = "./licenses"

type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func getAlbums(c *gin.Context) {
	id, idExists := c.GetQuery("id")
	if idExists {
		getAlbumByID(c, id)
	} else {
		albumsFromDB := readAlbumsFromDB()
		c.IndentedJSON(http.StatusOK, albumsFromDB)
	}
}

func main() {
	router := gin.Default()
	router.GET("/album", getAlbums)
	router.POST("/album", postAlbums)

	router.GET("/system-time", getTime)

	router.GET("/license", getLicenseNames)

	listenerIP, present := os.LookupEnv("LISTENER_IP")
	serverIP := "127.0.0.1"
	if present {
		serverIP = listenerIP
	}
	router.Run(serverIP + ":8080")
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		print(err)
		return
	}
	success := writeAlbumToDB(newAlbum)
	if success {
		c.IndentedJSON(http.StatusCreated, newAlbum)
	} else {
		c.Status(http.StatusBadRequest)
	}

}

func getTime(c *gin.Context) {
	format, isFormatSet := c.GetQuery("format")
	var out []byte
	var err error
	if isFormatSet {
		log.Println("format: " + format)
		out, err = exec.Command("date", format).Output()
		// out, err = exec.Command("date", "+%Y").Output()

	} else {
		out, err = exec.Command("date").Output()
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"time": string(out)})
	}
}

func getAlbumByID(c *gin.Context, id string) {
	result, err := readByIDFromDB(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func getLicenseNames(c *gin.Context) {
	filename, hasFilename := c.GetQuery("filename")
	if hasFilename {
		getLicenseText(c, filename)
	} else {
		files, err := ioutil.ReadDir(licenseFolder)
		if err != nil {
			log.Fatal(err)
		}
		var fileNames []string
		for _, f := range files {
			fileNames = append(fileNames, f.Name())
		}
		c.IndentedJSON(http.StatusOK, gin.H{"files": fileNames})
	}
}

func getLicenseText(c *gin.Context, filename string) {
	// vulnerability - LFI
	filePath := path.Join(licenseFolder, filename)
	log.Println("filePath: " + filePath)
	c.File(filePath)
}
