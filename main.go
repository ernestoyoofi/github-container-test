package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	addrlisten := ":3500"

	http.HandleFunc("/", Test)

	log.Printf("Running at port http://%v\n", addrlisten)
	http.ListenAndServe(addrlisten, nil)
}

type TestHTTPData struct {
	Message string `json:"message"`
}
type TestHTTP struct {
	Status int          `json:"status"`
	Data   TestHTTPData `json:"data"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	writeResponseData := TestHTTPData{
		Message: "Success!",
	}
	writeResponse := TestHTTP{
		Status: 200,
		Data:   writeResponseData,
	}

	json.NewEncoder(w).Encode(writeResponse)
}
