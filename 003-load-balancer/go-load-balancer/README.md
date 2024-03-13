# go-load-balancer
Creates a reverse proxy which redirects to available servers picked on a round robin basis

```bash
go run main.go
serving requests at 'localhost:8000'
forwarding request to address "https://www.facebook.com"
forwarding request to address "http://www.bing.com"
forwarding request to address "http://www.duckduckgo.com"
forwarding request to address "https://www.facebook.com"
```
