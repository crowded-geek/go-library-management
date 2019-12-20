package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

var books [4]string
var issuedBooks [2]string
var m map[string]string
var mostIssued string="Atlanta"
var topTrend string="Where magic ends"

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func bookAvailable(w http.ResponseWriter, r *http.Request){
	params := r.URL.Query()
        link := params.Get("book")
	fmt.Fprintf(w, strconv.FormatBool(contains(books, link)))
	fmt.Println("Endpoint Hit: bookAvailable")
}

func bookIssued(w http.ResponseWriter, r *http.Request){
	params := r.URL.Query()
        link := params.Get("book")
	fmt.Fprintf(w, strconv.FormatBool(containss(issuedBooks, link)))
	fmt.Println("Endpoint Hit: bookIssued")
}

func mostIssuedBook(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, mostIssued)
	fmt.Println("Endpoint Hit: mostIssuedBook")
}

func topTrendingBook(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, topTrend)
	fmt.Println("Endpoint Hit: topTrendingBook")
}

func issuedTo(w http.ResponseWriter, r *http.Request){
	params := r.URL.Query()
        link := params.Get("book")
	fmt.Fprintf(w, m[link])
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

func contains(arr [4]string, str string) bool {
   for _, a := range arr {
      if a == str {
         return true
      }
   }
   return false
}

func containss(arr [2]string, str string) bool {
   for _, a := range arr {
      if a == str {
         return true
      }
   }
   return false
}

func main() {

    m = make(map[string]string)
books[0]="Harry Potter Philosopher's stone"
books[1]="Atlanta"
books[2]="Where magic ends"
books[3]="When i'll fight"
issuedBooks[0]="Atlanta"
issuedBooks[1]="Where magic ends"
m["Atlanta"]="aayushman"
m["Where magic ends"]="kamihe"
m["When i'll fight"]="Nokomata"
m["Harry Potter Philosopher's stone"]="Natamo"
    handleRequests()
}
