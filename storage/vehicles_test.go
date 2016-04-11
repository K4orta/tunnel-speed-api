package storage

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/k4orta/muni"
)

func TestInsertVehicle(t *testing.T) {
	RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		var vehicle = muni.Vehicle{
			Id:               "1234",
			RouteTag:         "N",
			Lat:              0.0,
			Lng:              0.0,
			DirTag:           "N____I_F00",
			Heading:          1,
			LeadingVehicleId: "1233",
			Predictable:      true,
			SpeedKmHr:        0,
			SecsSinceReport:  0,
			TimeRecieved:     time.Now(),
		}
		err := InsertVehicle(db, &vehicle)
		if err != nil {
			t.Error(err)
		}

		v, err := GetVehiclesById(db, "1234")

		if err != nil {
			t.Error(err)
		}

		if v[0].Id != "1234" {
			t.Error("Didn't insert vehicle with correct ID")
		}
	})
}
