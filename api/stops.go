package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k4orta/muni"
)

func Stops(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	sd, err := muni.GetStopData(vars["route"])
	if err != nil {
		log.Panic(err)
	}
	out, err := json.Marshal(sd.Routes)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprint(w, string(out))
}
