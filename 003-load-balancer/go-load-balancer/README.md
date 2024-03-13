# go-load-balancer
Load balancer
- Keeps a list of servers
- Picks an available server on a round robin basis
- Creates a reverse proxy which redirects to the available server

```bash
go run main.go
serving requests at 'localhost:8000'
forwarding request to address "https://www.facebook.com"
forwarding request to address "http://www.bing.com"
forwarding request to address "http://www.duckduckgo.com"
forwarding request to address "https://www.facebook.com"
```
