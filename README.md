# RESTful_API

REST API using [GoLang](https://github.com/golang/go), [gorilla mux](https://github.com/gorilla/mux), [Basic Auth](https://golangbyexample.com/http-basic-auth-golang/), [JWT Auth](https://github.com/dgrijalva/jwt-go)
<hr/>


### Brief Description

A handler function accepts http response and request in json format. Then, the request is decoded and written to response according to the called      function. This handler function is wrapped by the authentication middleware to perform the security check. For login, basic authenticantion  is used and for authorization, Jwt authenticantion is used.

<hr/>

### API Endpoints

| URL    |  Function   |Method      |Description  |Authentication type         |   
| ----------- | ----------- | ----------- |-----------|-----------|
| https://localhost:8080/login  | LogIn	         |   POST       | 	Return jwt token in response for successful authentication  |  Basic          |
| https://localhost:8080/books  | GetBooks      |    GET        | Returns the description of all the books  |  JWT         |
| https://localhost:8080/books/{id} | GetBook   |    GET     | Returns the description of the book with the valid requested book id |  JWT         |
| https://localhost:8080/booksCreate | CreateBook   |    POST     | 	Creates a new book description  |  JWT         |
| https://localhost:8080/booksUpdate/{id} | UpdateBook   |    PUT    | Updates the decription of the requested book id  |  JWT         |
| https://localhost:8080/booksDelete/{id} | DeleteBook   |    DELETE    | 	Deletes the book specified by id |  JWT         |

<hr/>

### Data Model

Author Model

    type Author struct {
    	Firstname string `json:"firstname"`
    	Lastname  string `json:"lastname"`
    }

Book Model

    type Book struct {
    	ID     string  `json:"id"`
    	Isbn   string  `json:"isbn"`
    	Title  string  `json:"title"`
    	Author *Author `json:"author"`
    }
    
<hr/>

### Installation
  - [Postman](https://www.postman.com/downloads/)</li>
  - [RESTful_API](https://github.com/FahaF/RESTful_API)</li>
    - go install github.com/FahaF/RESTful_API
<hr/>

### Environment Variables

    export username="any user name without qoutation"
    export password="any password without qoutation"
<hr/>

### Server Run

    go run . 

<hr/>

