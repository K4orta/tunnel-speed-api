package jobs

import (
	"log"
	"time"

	"github.com/k4orta/tunnel-watch/storage"
)

// Deletes all vehicles older than a day
func ExpireOldVehicles() {
	db, err := storage.CreateConnection()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	storage.RemoveVehiclesOlderThan(db, time.Now().Add(time.Hour*-24))
}
