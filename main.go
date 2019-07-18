package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Orang struct (Model)//ghjklaxaa
type Orang struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Pesan struct {
	KD           string `json:"kd"`
	OrderDate    string `json:"orderdate"`
	EstimateDate string `json:"estimateDate"`
	Status       string `json:"status"`
}

type Item struct {
	ID          string `json:"id"`
	ProductName string `json:"productname"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
}
type Desk struct {
	Router string `json:"router"`
	URL    string `json:"URL"`
}

var customers []Orang
var order []Pesan
var products []Item
var information []Desk

// Get all customers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
func getInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(information)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	customers = append(customers, Orang{ID: "001", Name: "Malik", Age: 23, Address: "Meruya Selatan"})
	customers = append(customers, Orang{ID: "002", Name: "Agesta", Age: 23, Address: "Ciledug"})

	order = append(order, Pesan{KD: "001", OrderDate: "17-07-2019", EstimateDate: "19-07-2019", Status: "Delivered"})
	order = append(order, Pesan{KD: "002", OrderDate: "19-08-2019", EstimateDate: "23-08-2019", Status: "Pending"})

	products = append(products, Item{ID: "001", ProductName: "Kemeja Biru Polos", Stock: 24, Price: 300000})
	products = append(products, Item{ID: "002", ProductName: "Selendang Merah Muda", Stock: 4, Price: 150000})

	information = append(information, Desk{Router: "gorilla/mux", URL: "/api/info"})

	// Router Handlers / Endpoints
	r.HandleFunc("/api/customers", getCustomers).Methods("GET")
	r.HandleFunc("/api/order", getOrder).Methods("GET")
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/info", getInformation).Methods("GET")

	log.Fatal(http.ListenAndServe(":9001", r))
}

// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// //Book struct (Model)
// type Book struct {
// 	ID     string  `json:"id"`
// 	Isbn   string  `json:"isbn"`
// 	Title  string  `json:"title"`
// 	Author *Author `json:"author"`
// }

// //Author struct
// type Author struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

// // Init books var as a slice book struct
// var books []Book

// // Get all books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// // Get single Book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) //Get params
// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

// // Create a new book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(rand.Intn(10000000))
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)

// }

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// func main() {
// 	// Init Router
// 	r := mux.NewRouter()

// 	// Mock Data - @todo - implement DB
// 	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
// 	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
// 	books = append(books, Book{ID: "3", Isbn: "454687", Title: "Book Three", Author: &Author{Firstname: "Joe", Lastname: "Inoe"}})

// 	// Router Handlers / Endpoints
// 	r.HandleFunc("/api/books", getBooks).Methods("GET")
// 	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
// 	r.HandleFunc("/api/books", createBook).Methods("POST")
// 	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
// 	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe(":8000", r))
// }
