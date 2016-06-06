package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/k4orta/muni"
	"github.com/k4orta/tunnel-speed-api/storage"
	"github.com/k4orta/tunnel-speed-api/utils"
)

// Vehicles reutrns the vehicles belonging to a specific line
func Vehicles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	vd, _ := muni.GetVehiclesData(vars["route"])

	out, err := json.Marshal(vd)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(out))
}

// AllVehicles returns all of the vehicles in the system
func AllVehicles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := storage.CreateConnection()
	if err != nil {
		log.Println("Error connecting to DB in AllVehicles")
		return
	}
	defer db.Close()
	vd, err := storage.GetVehiclesAfterTime(db, time.Now().Add(time.Minute*-4))
	compacted := utils.CompactVehicles(vd)
	out, err := json.Marshal(compacted)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(out))
}

//
