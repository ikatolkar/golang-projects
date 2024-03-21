# URL Shortener

Service that accepts URL and returns short URL. Also accepts saved short URL and redirects to original URL.
- Serves over HTTP
- http api built using go-fiber
- uses redis database, connects using go-redis
- ratelimiting is done using redis' key TTL feature
    - the api access counter will have a TTL, and will decrease on every succesful api request
    - will be reset if there is no entry (the entry will get cleared after 30 minutes automatically)
- using godotenv for reading env file
- deployed using docker-compose

## Shorten using a custom short
```bash
curl -X POST http://172.29.87.67:32778/api/v1 -H "Content-Type: application/json" -d '{"url":"youtube.com","short":"utube" }'
{"url":"http://youtube.com","short":"localhost:3000/utube","expiry":24,"rate_limit":8,"rate_limit_reset":27}
```

## Redirects if short is found
```bash
curl -X GET http://172.29.87.67:32778/utube -I
HTTP/1.1 301 Moved Permanently
Date: Thu, 21 Mar 2024 08:09:50 GMT
Content-Length: 0
Location: http://youtube.com
```

## Error if custom short is already in use
```bash
curl -X POST http://172.29.87.67:32778/api/v1 -H "Content-Type: application/json" -d '{"url":"youtube.com","short":"utube" }'
{"error":"URL custom short is already in use"}
```

## Get random short, with expiry
```bash
curl -X POST http://172.29.87.67:32778/api/v1 -H "Content-Type: application/json" -d '{"url":"youtube.com","expiry":3 }'
{"url":"http://youtube.com","short":"localhost:3000/c1c583","expiry":3,"rate_limit":7,"rate_limit_reset":23}

#---
curl -X POST http://172.29.87.67:32778/api/v1 -H "Content-Type: application/json" -d '{"url":"youtube.com","short":"utube" }'
{
    "error": "rate limit exceeded",
    "rate_limit_reset": 11
}
```
