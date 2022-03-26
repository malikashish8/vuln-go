package main

import (
	"log"
	"net/http"
	"os"
)

var licenseFolder string = "./licenses"

type album struct {
	ID     int     `json:"id,omitempty"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type httpError struct {
	httpCode int
	message  string
}

// call the right func based on http method
func handleAlbum(w http.ResponseWriter, r *http.Request) {
	h := httpHelper{w, r}
	if r.Method == "GET" {
		getAlbums(h)
	} else if r.Method == "POST" {
		postAlbums(h)
	}
}

func main() {
	mux := http.NewServeMux()
	// configure handlers
	mux.HandleFunc("/album", handleAlbum)
	mux.HandleFunc("/systemtime", getTime)
	mux.HandleFunc("/license", getLicenseNames)

	listenerIP, present := os.LookupEnv("LISTENER_IP")
	serverIP := "127.0.0.1"
	if present {
		serverIP = listenerIP
	}
	listenAddress := serverIP + ":8080"
	log.Println("listening on " + listenAddress)
	http.ListenAndServe(listenAddress, mux)
}
