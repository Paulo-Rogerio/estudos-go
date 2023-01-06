package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			Id:   1,
			Name: "Paulo",
		},
		{
			Id:   2,
			Name: "Camilla",
		},
	})
}

func main() {

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/users", getUsers).Methods("GET")

	fmt.Println("Listen in :3000")

	log.Fatal(http.ListenAndServe(":3000", muxRouter))

}
