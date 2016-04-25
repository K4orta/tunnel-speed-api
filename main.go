package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/k4orta/tunnel-speed-api/api"
	"github.com/k4orta/tunnel-speed-api/jobs"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/routes/{route}", api.Stops)
	router.HandleFunc("/routes", api.AllStops)
	router.HandleFunc("/vehicles/{route}", api.Vehicles)
	router.HandleFunc("/vehicles", api.AllVehicles)

	n := negroni.New()

	go jobs.RunFetch()
	go jobs.RunExpire()

	n.UseHandler(cors.Default().Handler(router))

	n.Run(":8048")
}
