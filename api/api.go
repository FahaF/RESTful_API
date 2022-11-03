package api

import (
	"GoLangProject/auth"
	"GoLangProject/data"
	"GoLangProject/model"
	"encoding/json" // core package
	"log"           // log errors
	"math/rand"     // to add id as a randum number
	"net/http"      // to work with http
	"strconv"       // for string conversion

	"github.com/gorilla/mux" // router
)

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("getUsers")
	log.Println("Authentication successful!")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data.Books); err != nil {
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
	for _, item := range data.Books {
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
	book := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Println(err.Error())
		return
	}
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	for _, item := range data.Books {
		if book.ID == item.ID {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	data.Books = append(data.Books, book)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data.Books); err != nil {
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
	for index, item := range data.Books {
		if item.ID == params["id"] {
			data.Books = append(data.Books[:index], data.Books[index+1:]...)
			book := model.Book{}
			book.ID = params["id"]
			data.Books = append(data.Books, book)
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(data.Books); err != nil {
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

	for index, item := range data.Books {
		if item.ID == params["id"] {
			data.Books = append(data.Books[:index], data.Books[index+1:]...)
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(data.Books); err != nil {
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
