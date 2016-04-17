package jobs

import (
	"log"
	"time"

	"github.com/k4orta/tunnel-watch/storage"
)

// RunExpire loops every hour and clean up entries older than a day
func RunExpire() {
	for {
		expireOldVehicles(time.Now().Add(time.Hour * -24).UTC())
		time.Sleep(time.Hour)
	}
}

// ExpireOldVehicles deletes all vehicles older than a day
func expireOldVehicles(expDate time.Time) {
	db, err := storage.CreateConnection()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	storage.RemoveVehiclesOlderThan(db, expDate)
}
