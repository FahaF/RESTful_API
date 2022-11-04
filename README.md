# RESTful_API

REST API using [GoLang](https://github.com/golang/go), [gorilla mux](https://github.com/gorilla/mux), [Basic Auth](https://golangbyexample.com/http-basic-auth-golang/), [JWT Auth](https://github.com/dgrijalva/jwt-go)

### Installation
  - [Postman](https://www.postman.com/downloads/)</li>
  - [RESTful_API](https://github.com/FahaF/RESTful_API)</li>
    - go install github.com/FahaF/RESTful_API
  
### Server Run

### API Endpoints

| Endpoint    |  Function   |Method      |StatusCode  |Auth       |   
| ----------- | ----------- | ----------- |-----------|-----------|
| /login  | LogIn	         |   POST       | Success - StatusOK, Failure - StatusUnauthorized  |  Basic          |
| /books  | GetBooks      |    GET        | Success - StatusOK  |  JWT         |
| /books/{id} | GetBook   |    GET     | Success - StatusOK, Failure - StatusNoContent  |  JWT         |
| /booksCreate | CreateBook   |    POST     | Success - StatusCreated, Failure - StatusConflict  |  JWT         |
| /booksUpdate/{id} | UpdateBook   |    PUT    | Success - StatusCreated, Failure - StatusNoContent  |  JWT         |
| /booksDelete/{id} | DeleteBook   |    DELETE    | Success - StatusOK, Failure - StatusNoContent  |  JWT         |

