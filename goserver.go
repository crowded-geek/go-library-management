package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func bookAvailable(w http.ResponseWriter, r *http.Request){
	//doSmth to get the data and display
	fmt.Fprintf(w, "bookAvailable?")
	fmt.Println("Endpoint Hit: bookAvailable")
}

func bookIssued(w http.ResponseWriter, r *http.Request){
	//doSmth to get the data and display
	fmt.Fprintf(w, "bookIssued?")
	fmt.Println("Endpoint Hit: bookIssued")
}

func mostIssuedBook(w http.ResponseWriter, r *http.Request){
	//doSmth to get the data and display
	fmt.Fprintf(w, "someRandomBookIdon'tKnowIfExists")
	fmt.Println("Endpoint Hit: mostIssuedBook")
}

func topTrendingBook(w http.ResponseWriter, r *http.Request){
	//doSmth to get the data and display
	fmt.Fprintf(w, "topTrendingBook?Idon'tKnowWhich")
	fmt.Println("Endpoint Hit: topTrendingBook")
}

func issuedTo(w http.ResponseWriter, r *http.Request){
	//doSmth to get the data and display
	fmt.Fprintf(w, "book is given to user 1")
	fmt.Println("Endpoint Hit: issuedTo")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/bookAvailable", bookAvailable)
    http.HandleFunc("/bookIssued", bookIssued)
    http.HandleFunc("/mostIssuedBook", mostIssuedBook)
    http.HandleFunc("/topTrendingBook", topTrendingBook)
    http.HandleFunc("/issuedTo", issuedTo)
    log.Fatal(http.ListenAndServe(":8282", nil))
}

func main() {
    handleRequests()
}