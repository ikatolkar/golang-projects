# PER CLIENT

Keep track of clients by their IP. IP can be taken from the request.
Clients are kepts in a map as map[IP] = Client
Each client represents an Ip and the last time it tried to connect
If its been more than 3 minutes remove client from the map. This means if client exists in the map, it must have been seen within last 3 minutes.
To delete the clients from the map, call a go routine.
Each client also contains a separate limiter, which can be called on handling its request.

NOTES:
Notice how:
- functions are nested
- functions are defined as callbacks
- different files are there in same package


```bash
$ for x in [ 1 ... 30 ]; do curl http://localhost:8000/ping; done
{"status":"Successful","body":"Hi! You've reached the API"}
{"status":"Successful","body":"Hi! You've reached the API"}
{"status":"Successful","body":"Hi! You've reached the API"}
{"status":"Successful","body":"Hi! You've reached the API"}
{"status":"Request failed","body":"The API is at capacity, try again later"}
```
