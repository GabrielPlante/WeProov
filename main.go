package main

import (
	database "WeProov/Database"
	weproovuser "WeProov/User"
	view "WeProov/View"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//First the functions handling each page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, view.MainPage())
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, view.CreateUser())
}

func postUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, view.PostUser())
	user := weproovuser.User{
		Username:     r.FormValue("username"),
		PasswordHash: weproovuser.HashPassword(r.FormValue("password")),
		Email:        r.FormValue("email"),
	}
	database.AddUser(user)
}

//No view for this one because too time consuming
func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of every user:\n")
	users := database.GetAllUsers()
	for i := 0; i != len(users); i++ {
		fmt.Fprintf(w, "User "+strconv.Itoa(i)+":\n")
		fmt.Fprintf(w, "Username: "+users[i].Username+":\n")
		fmt.Fprintf(w, "Email: "+users[i].Email+":\n")
	}
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, view.RemoveUser())
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, view.DeleteUser())
	database.DeleteUser(r.FormValue("username"))
}

//Function handling the request of the user, redirecting to previous functions
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/createuser", createUser)
	router.HandleFunc("/removeuser", removeUser)

	router.HandleFunc("/user", getUser).Methods("GET")
	router.HandleFunc("/user", postUser).Methods("POST")
	router.HandleFunc("/deleteuser", deleteUser) //Delete method won't work
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	//database.Connect()
	//Just launch the server
	handleRequests()
}
