# Token Bucket Rate Limiter
Rate limits for the entire endpoint
Each endpoint handler will invoke the golang.org/x/time/rate limiter
Which will limit messages at R rate, ie R messages per second with an allowed burst of 4

```bash
# go run .
{"status":"Successful","body":"Hi! You've reached the API"}
{"status":"Successful","body":"Hi! You've reached the API"}
{"status":"Successful","body":"Hi! You've reached the API"}
{"status":"Successful","body":"Hi! You've reached the API"}
{"status":"Request Failed","body":"The API is at capacity, try later"}
```
