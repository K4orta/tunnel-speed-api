package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k4orta/muni"
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

	vd, _ := muni.GetMultiVehicleData([]string{"N", "L", "J", "KT", "M"})

	out, err := json.Marshal(vd)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(out))
}
