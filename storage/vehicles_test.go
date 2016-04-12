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

		v, err := GetVehiclesByID(db, "1234")

		if err != nil {
			t.Error(err)
		}

		if v[0].Id != "1234" {
			t.Error("Didn't insert vehicle with correct ID")
		}
	})
}

func TestGetVehiclesByTime(t *testing.T) {
	RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		var newVehicle = muni.Vehicle{Id: "1234", TimeRecieved: time.Now()}
		var newVehicle2 = muni.Vehicle{Id: "1234", TimeRecieved: time.Now().Add(-time.Minute)}
		var oldVehicle = muni.Vehicle{Id: "1235", TimeRecieved: time.Unix(1420919252102, 0)}

		InsertVehicle(db, &newVehicle)
		InsertVehicle(db, &newVehicle2)
		InsertVehicle(db, &oldVehicle)

		v, err := GetVehiclesByTime(db, time.Minute*5)

		if err != nil {
			t.Error(err)
		}

		if len(v) != 2 {
			t.Error("Wrong number of items returned")
		}

	})
}
