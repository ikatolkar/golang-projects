package main


import (
	"log"
	"net/http"
	"encoding/json"
)

/*
The return message to be written in responseWriter
*/
type Message struct {
	Status string `json:"status"`
	Body string `json:"body"`
}

/*
The hander used by http.Handle
This will eventually be called by rateLimiter if limiter is OK
*/
func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	message := Message{
		Status: "Successful",
		Body: "Hi! You've reached the API",
	}
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

/*
Creates an http handler for endpoint /ping on localhost:8000
Starts listening on :8000
*/
func main() {
	http.Handle("/ping", rateLimiter(endpointHandler))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("There was an error listening on port :8000", err)
	} else {
		log.Println("Listening on port 8000")
	}
}
