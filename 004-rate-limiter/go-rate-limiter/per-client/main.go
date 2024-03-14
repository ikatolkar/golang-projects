package main

import (
	"log"
	"net/http"
	"time"
	"net"
	"golang.org/x/time/rate"
	"encoding/json"
	"sync"
)

type Message struct {
	Status string `json:"status"`
	Body string `json:"body"`
}
func perClientRateLimiter(next func(writer http.ResponseWriter, request *http.Request)) http.Handler {

	type client struct {
		limiter *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu sync.Mutex
		clients = make(map[string]*client)
	)
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, client := range clients{
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
		}
	}()
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = &client{limiter: rate.NewLimiter(2, 4)}
		}
		clients[ip].lastSeen = time.Now()
		if ! clients[ip].limiter.Allow() {
			mu.Unlock()
			message := Message{
				Status: "Request failed",
				Body: "The API is at capacity, try again later",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		}
	mu.Unlock()
	next(w, req)
	})
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	message := Message {
		Status: "Successful",
		Body: "Hi! You've reached the API",
	}
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func main() {
	http.Handle("/ping", perClientRateLimiter(endpointHandler))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("Failed to listen at port 8000")
	}
}
