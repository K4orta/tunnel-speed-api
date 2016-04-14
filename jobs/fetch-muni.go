package jobs

import (
	"log"
	"time"

	"github.com/k4orta/muni"
	"github.com/k4orta/tunnel-watch/storage"
)

// RunFetch loops forever and collects transit data
func RunFetch() {
	for {
		log.Println("Starting Fetch")
		err := fetchTransit([]string{"N"})
		if err != nil {
			log.Println("ERROR:", err)
		}
		time.Sleep(time.Minute)
	}
}

func fetchTransit(lines []string) error {
	vehicles, err := muni.GetMultiVehicleData(lines)
	if err != nil {
		return err
	}
	db, err := storage.CreateConnection()
	if err != nil {
		log.Println(db)
	}
	for _, v := range vehicles {
		storage.InsertVehicle(db, v)
	}
	return nil
}
