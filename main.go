package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/k4orta/tunnel-speed-api/api"
	"github.com/k4orta/tunnel-speed-api/jobs"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/routes/{route}", api.Stops)
	router.HandleFunc("/routes", api.AllStops)
	router.HandleFunc("/vehicles/{route}", api.Vehicles)
	router.HandleFunc("/vehicles", api.AllVehicles)

	n := negroni.New()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(cors.Default().Handler(router))

	go jobs.RunFetch()
	go jobs.RunExpire()

	n.Run(":8048")
}
