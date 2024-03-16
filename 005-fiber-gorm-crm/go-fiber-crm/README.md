# NOTE on Go Fiber
go fiber is based on fasthttp, which is not compatible with real http.
It doesn't even support http2 or http3.
Whenever possible use mux or even easier net/http.

# go-fiber-crm
CRM using gorm+sqlite, and go fiber
```bash

$ curl -X POST -H "Accept: application/json" http://localhost:3000/api/v1/lead -d '{"name":"John Doe", "email":"johndoe@example.com", "company":"ACME", "phone":8888888881}'
{"ID":1,"CreatedAt":"2024-03-17T00:33:34.7083264+05:30","UpdatedAt":"2024-03-17T00:33:34.7083264+05:30","DeletedAt":null,"name":"","company":"","email":"","phone":0}

$ curl -X GET http://localhost:3000/api/v1/lead
[{"ID":1,"CreatedAt":"2024-03-17T00:33:34.7083264+05:30","UpdatedAt":"2024-03-17T00:33:34.7083264+05:30","DeletedAt":null,"name":"","company":"","email":"","phone":0}]root@PF2Z4T5C-inl:.../005-fiber-crm/go-fiber-crm

$ curl -X GET http://localhost:3000/api/v1/lead/1
{"ID":1,"CreatedAt":"2024-03-17T00:33:34.7083264+05:30","UpdatedAt":"2024-03-17T00:33:34.7083264+05:30","DeletedAt":null,"name":"","company":"","email":"","
```
