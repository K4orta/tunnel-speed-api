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
		log.Println(err)
	}
	out, err := json.Marshal(sd.Routes)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, string(out))
}

func AllStops(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lines := fetchMultiStop([]string{"N", "J", "KT", "M", "L"})

	out, err := json.Marshal(lines)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, string(out))
}

func fetchMultiStop(stops []string) []*muni.Route {
	lines := []*muni.Route{}

	stopData := make(chan *muni.StopResponse, len(stops))
	for _, s := range stops {
		go func(s string) {
			sd, err := muni.GetStopData(s)
			if err != nil {
				log.Println(err)
			}
			stopData <- sd
		}(s)
	}

	for {
		select {
		case resp := <-stopData:
			lines = append(lines, resp.Routes[0])
			if len(stops) == len(lines) {
				return lines
			}
		}
	}
}
