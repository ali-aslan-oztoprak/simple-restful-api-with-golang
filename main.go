package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

func main() {
    fmt.Println("Rest API v1.0 - Simple RESTful API")
	handleRequests()
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", defaultPage)
    router.HandleFunc("/users", getUsers)
    router.HandleFunc("/user/{id}", getUser)
	log.Fatal(http.ListenAndServe(port, router))
}

func defaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Users)
}

func getUser(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
    intKey, err := strconv.Atoi(key)

    if (err != nil) {
        fmt.Println(err)
    }

    for _, user := range Users {
            if user.Id == intKey {
                json.NewEncoder(w).Encode(user)
            }
        }
}

type User struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Surname string `json:"surname"`
}

var port string = ":8000";
var Users []User = []User{
    User{Id: 1, Name: "John", Surname: "Doe"},
    User{Id: 2, Name: "Jane", Surname: "Doe"},
}
