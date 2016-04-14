package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/k4orta/tunnel-watch/api"
	"github.com/k4orta/tunnel-watch/jobs"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/stops/{route}", api.Stops)
	router.HandleFunc("/vehicles/{route}", api.Vehicles)
	router.HandleFunc("/vehicles", api.AllVehicles)

	n := negroni.New()

	go jobs.RunFetch()

	n.UseHandler(cors.Default().Handler(router))

	n.Run(":8048")
}
