package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
//third party package for routing
    "github.com/gorilla/mux"
)

// Book - Our struct for all books
type Book struct {
    Id      string    `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Author string `json:"author"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	//printing to the terminal
    fmt.Println("Endpoint Hit: homePage")
	//printing to the route route
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func returnAllBook(w http.ResponseWriter, r *http.Request) {
	//console printing
    fmt.Println("Endpoint Hit: returnAllBook")
	//encoding Books to json and returnig and writing on the route
    json.NewEncoder(w).Encode(Books)
	fmt.Fprint(w, json.NewEncoder(w).Encode(Books))
	
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	//getting id from the url
    vars := mux.Vars(r)
    key := vars["id"]

	//checking same id and returnign for that route
    for _, book := range Books {
        if book.Id == key {
            json.NewEncoder(w).Encode(book)
			fmt.Fprint(w, json.NewEncoder(w).Encode(book))

        }
    }
}


func createNewBook(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Book struct
    // append this to our Book array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var book Book 
    json.Unmarshal(reqBody, &book)
    // update our global Book array to include
    // our new Book
    Books = append(Books, book)

	//returning
    json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

	//reframing book to contain only those indexed without index corresponding to given id 
    for index, book := range Books {
        if book.Id == id {
            Books = append(Books[:index], Books[index+1:]...)
        }
    }

}

func updateBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
    id := vars["id"]
	// fmt.Println("Endpoint Hit: ")
    // get the body of our POST request
    // unmarshal this into a new Book struct
    // append this to our Book array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var book Book 
    json.Unmarshal(reqBody, &book)
    // update our global Books array to include
    // our new Book
    for index, book2 := range Books {
        if book2.Id == id {
			//updating all the fields
            Books[index].Title = book.Title
			Books[index].Desc = book.Desc
			Books[index].Author = book.Author

        }
    }

	//returnig the book added
    json.NewEncoder(w).Encode(book)
}

//to handle requests
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/book/see-all", returnAllBook)
    myRouter.HandleFunc("/book/add", createNewBook).Methods("POST")
    myRouter.HandleFunc("/book/delete/{id}", deleteBook).Methods("DELETE")
    myRouter.HandleFunc("/book/see/{id}", returnSingleBook)
	myRouter.HandleFunc("/book/update/{id}", updateBook).Methods("PUT")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}


func main() {
	//intial records
    Books = []Book{
        Book{Id: "1", Title: "Hello", Desc: "Book Description", Author: "Author 1"},
        Book{Id: "2", Title: "Hello 2", Desc: "Book Description", Author: "Author 2"},
    }
	//making these routes functioning by including in main
    handleRequests()
}