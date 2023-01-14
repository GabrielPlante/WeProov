package main

import (
	view "WeProov/View"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//A standard user model
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordhash"`
	Email        string `json:"email"`
}

//The base page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, view.CreateUser())
}

//Function handling the request of the user
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/createuser", createUser)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	/*users := []User{
		User{Username: "Olivier",
			PasswordHash: "sdflkmqsdvnqsrlqsm",
			Email:        "oliver.du69@gmail.com"},
	}*/
	handleRequests()
}
