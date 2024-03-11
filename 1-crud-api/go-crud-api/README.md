# go-crud-api
CRUD API Using Gorilla Mux

# How to use
```
# go run .
```

# Example
```
CREATE http://localhost:8888/movies
GET    http://localhost:8888/movies/1
UPDATE http://localhost:8888/movies/1
DELETE http://localhost:8888/movies/1

RAW BODY
{
    "isbn": "123123123",
    "title": "movie2",
    "director": {
        "firstname": "Jackie",
        "lastname": "Chan"        
    }

}
```
