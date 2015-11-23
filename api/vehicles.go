package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k4orta/muni"
)

func Vehicles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	vd, _ := muni.GetVehiclesData(vars["route"])

	out, _ := json.Marshal(vd)
	fmt.Fprint(w, string(out))
}

func AllVehicles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vd, _ := muni.GetMultiVehicleData([]string{"N", "L", "J"})

	out, _ := json.Marshal(vd)
	fmt.Fprint(w, string(out))
}
