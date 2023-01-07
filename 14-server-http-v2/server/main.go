package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Monitor Struct
type Monitor struct {
	Id         string `json:id`
	Monitor_Id string `json:monitor_id`
	Profile    string `json:profile`
}

// Init Monitors ( Dados Mock )
var monitors []Monitor

// Get All Monitors
func getMonitors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monitors)
}

// Get Single Monitor
func getMonitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range monitors {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Monitor{})
}

// Create Single Monitor
func createMonitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var monitor Monitor
	_ = json.NewDecoder(r.Body).Decode(&monitor)
	monitor.Id = strconv.Itoa(rand.Intn(10000))
	monitors = append(monitors, monitor)
	json.NewEncoder(w).Encode(monitor)
}

func main() {

	// Init Router
	r := mux.NewRouter()

	// Mock Data
	monitors = append(monitors, Monitor{Id: "1", Monitor_Id: "10", Profile: "monitor-teste-123"})

	// Router Handlers / Endpoints
	r.HandleFunc("/api/monitors", getMonitors).Methods("GET")
	r.HandleFunc("/api/monitors/{id}", getMonitor).Methods("GET")
	r.HandleFunc("/api/monitors", createMonitor).Methods("POST")

	// Listen Server
	fmt.Println("Listen in :3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
