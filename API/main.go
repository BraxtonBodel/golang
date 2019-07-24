package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   string `json:"id,omitempty"`
	Name *Name  `json:"name,omitempty"`
	Mail string `json:"mail,omitempty"`
	Code int    `json:"code,omitempty"`
}

type Name struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
}

var user []User

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range user {
		if item.ID == params["id"] {
			fmt.Println("Serving request", item.ID)
			//Interpreta lo que le mandamos al cliente
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var newUser User
	//Interpreta lo que el cliente manda
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = params["id"]
	user = append(user, newUser)
	json.NewEncoder(w).Encode(user)
	fmt.Println(w)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func main() {

	user = append(user, User{
		ID:   "1",
		Name: &Name{FirstName: "Braxton", LastName: "Bonilla Delgado"},
		Mail: "braxton@yahoo.com", Code: 1996})

	user = append(user, User{
		ID:   "2",
		Name: &Name{FirstName: "Patricia", LastName: "Delgado Delgado"},
		Mail: "patty@yahoo.com", Code: 1231})

	router := mux.NewRouter()

	//Endpoints
	router.HandleFunc("/user", getAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/user/{id}", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))

}
