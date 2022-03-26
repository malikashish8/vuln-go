package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	h := httpHelper{w, r}
	format, isFormatSet := h.GetQueryParam("format")
	var out []byte
	var err error
	if isFormatSet {
		log.Println("format: " + format)
		// vulnerability - RCE - OS Command Injection
		out, err = exec.Command("sh", "-c", "date "+format).Output()
	} else {
		out, err = exec.Command("sh", "-c", "date").Output()
	}
	if err != nil {
		h.WriteError(httpError{http.StatusInternalServerError, err.Error()})
	} else {
		outString := strings.TrimSpace(string(out))
		h.WriteJson(map[string]interface{}{"time": outString})
	}
}
