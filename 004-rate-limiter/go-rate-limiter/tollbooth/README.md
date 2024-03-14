# Tollbooth
Tollbooth : "github.com/didip/tollbooth/limiter" 
Tollbooth provides a lot of ratelimiting methods, the one used here is again the token bucket method
limits 1 request per second

```bash
$ for x in [ 1 ... 80 ]; do curl -X GET http://localhost:8000/ping; sleep .51; done
{"status":"Successful","body":"Hi, you've reached the api"}
{"status":"Request failed","body":"The API is at capacity, try again later"}
{"status":"Successful","body":"Hi, you've reached the api"}
{"status":"Request failed","body":"The API is at capacity, try again later"}
{"status":"Successful","body":"Hi, you've reached the api"}
```
