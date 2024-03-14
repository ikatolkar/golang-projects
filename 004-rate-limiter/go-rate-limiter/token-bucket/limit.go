package main

import (
	"encoding/json"
	"net/http"
	"golang.org/x/time/rate"
)

/*
limits the number of messages on an endpoint to R message per seconds with a total burst of B
*/
func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	/*
	NewLimiter returns a new Limiter that allows events up to rate r 
	and permits bursts of at most b tokens.
	rate r means r messages per second
	*/
	limiter := rate.NewLimiter(2, 4)
	/*
	return an http.Handler type of callback function which will be accepted by ServeAndListen
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ! limiter.Allow() {
			message := Message{
				Status: "Request Failed",
				Body: "The API is at capacity, try later",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(message)
			return
		} else {
			/*
			if limiter is OK
			return the handler that we got as input
			in this case EndpointHandler
			*/
			next(w, r)
		}
	})
}
