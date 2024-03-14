package main


import (
	"encoding/json"
	"log"
	"net/http"
	tollbooth "github.com/didip/tollbooth/v7"
)

type Message struct {
	Status string `json:"status"`
	Body string `json:"body"`
}

func endPointHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	message := Message{
		Status: "Successful",
		Body: "Hi, you've reached the api",
	}
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func main() {
	message := Message {
		Status: "Request failed",
		Body: "The API is at capacity, try again later",
	}
	jsonMessage, _ := json.Marshal(message)
	/*
	Limit 1 request per second and no expiry options
	*/
	tlbthLimiter := tollbooth.NewLimiter(1, nil)
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonMessage))
	http.Handle("/ping", tollbooth.LimitFuncHandler(tlbthLimiter, endPointHandler))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("There was an error listening on port :8000", err)
	}
}
