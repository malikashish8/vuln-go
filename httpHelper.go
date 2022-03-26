package main

import (
	"encoding/json"
	"net/http"
	"os"
)

// create custom helper functions since some SAST tools might not support
// the specific http library such as gin that we use
// this ensure that we use basic constucts of the language so that SAST tools
// that do not actually build the application can still find vulnerabilities
type httpHelper struct {
	w http.ResponseWriter
	r *http.Request
}

func (h httpHelper) GetQueryParam(paramName string) (paramValue string, exists bool) {
	q := h.r.URL.Query()
	paramValue = q.Get(paramName)
	if len(paramValue) > 0 {
		exists = true
	}
	return
}

func (h httpHelper) WriteJson(input interface{}) (err error) {
	output, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return
	}
	h.w.Header().Add("content-type", "application/json")
	_, err = h.w.Write(output)
	return
}

func (h httpHelper) WriteError(inputError httpError) (err error) {
	httpCode := inputError.httpCode
	message := inputError.message
	h.w.Header().Add("content-type", "application/json")
	h.w.WriteHeader(httpCode)
	s := map[string]string{"error": message}
	errorStr, _ := json.MarshalIndent(s, "", "  ")
	_, err = h.w.Write([]byte(string(errorStr)))
	return
}

func (h httpHelper) WriteFile(filePath string) (err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	_, err = h.w.Write(data)
	return
}

func (h httpHelper) BindJSON(obj interface{}) (err error) {
	err = json.NewDecoder(h.r.Body).Decode(&obj)
	return
}
