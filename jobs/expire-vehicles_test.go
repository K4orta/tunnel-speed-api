package jobs

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/k4orta/muni"
	"github.com/k4orta/tunnel-watch/storage"
)

func init() {
	storage.SetupDBForTesting()
}

func TestExpireVehicles(t *testing.T) {
	fakeServer := makeFakeServer()
	muni.SetConfig(muni.TransitConfig{DefaultURL: fakeServer.URL + "/"})

	storage.RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		err := fetchTransit([]string{"N"})
		if err != nil {
			t.Error(err)
		}
		ExpireOldVehicles()
		v, _ := storage.GetVehiclesBeforeTime(db, time.Unix(1460498787925/1000, 0).Add(time.Minute))

		if len(v) != 0 {
			t.Error("Failed to remove vehicles from DB")
		}
	})
}
