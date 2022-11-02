package api

import (
	"GoLangProject/auth"
	"encoding/json" // core package
	"log"           // log errors
	"math/rand"     // to add id as a randum number
	"net/http"      // tow work with http
	"strconv"       // fort string conversion

	"github.com/gorilla/mux" // router
)

// Book Struct

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var Books []Book

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("getUsers")
	log.Println("Authentication successful!")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Books); err != nil {
		log.Println(err.Error())
		return
	}
}

// Get single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("getUser")
	log.Println("Authentication successful!")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range Books {
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(item); err != nil {
				log.Println(err.Error())
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

// Add new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("addUser")
	log.Println("Authentication successful!")
	var book Book
	//_ = json.NewDecoder(r.Body).Decode(&book)
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Println(err.Error())
		return
	}
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	for _, item := range Books {
		if book.ID == item.ID {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	Books = append(Books, book)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(Books); err != nil {
		log.Println(err.Error())
		return
	}
}

// Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	log.Println("updateUser")
	log.Println("Authentication successful!")

	params := mux.Vars(r)
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			var book Book
			//	_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			Books = append(Books, book)
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(Books); err != nil {
				log.Println(err.Error())
				return
			}
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("updateUser")
	log.Println("Authentication successful!")

	params := mux.Vars(r)
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(Books); err != nil {
				log.Println(err.Error())
				return
			}
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("LogIn")
	log.Println("Authentication successful!")
	log.Println("successfully logged in!")

	token, err := auth.GetToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error generating JWT token: " + err.Error()))
	} else {
		w.Header().Set("Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Token: " + token))
	}
}

func HandleRoutes(port string) {
	log.Println("in HandleRoutes!")

	//init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	Books = append(Books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	Books = append(Books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// rout handlers / endpoint

	r.HandleFunc("/login", LogIn).Methods("POST")
	r.HandleFunc("/books", auth.JWTAuthentication(GetBooks)).Methods("GET") // r.HandleFunc(url,function).httpMethod
	r.HandleFunc("/books/{id}", auth.JWTAuthentication(GetBook)).Methods("GET")
	r.HandleFunc("/books", auth.JWTAuthentication(CreateBook)).Methods("POST")
	r.HandleFunc("/books/{id}", auth.JWTAuthentication(UpdateBook)).Methods("PUT")
	r.HandleFunc("/books/{id}", auth.JWTAuthentication(DeleteBook)).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, r))
}
