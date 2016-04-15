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
			ID:       "1234",
			RouteTag: "N",
		}
		err := InsertVehicle(db, &vehicle)
		if err != nil {
			t.Error(err)
		}

		v, err := GetVehiclesByID(db, "1234")

		if err != nil {
			t.Error(err)
		}

		if v[0].ID != "1234" {
			t.Error("Didn't insert vehicle with correct ID")
		}
	})
}

func TestGetVehiclesAfterTime(t *testing.T) {
	RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		var recentDate int64 = 1460432740083 / 1000
		var oldDate int64 = 1420919252102 / 1000
		var newVehicle = muni.Vehicle{ID: "1234", TimeRecieved: time.Unix(recentDate, 0)}
		var newVehicle2 = muni.Vehicle{ID: "1235", TimeRecieved: time.Unix(recentDate, 0).Add(-time.Minute)}
		var oldVehicle = muni.Vehicle{ID: "1236", TimeRecieved: time.Unix(oldDate, 0)}

		InsertVehicle(db, &newVehicle)
		InsertVehicle(db, &newVehicle2)
		InsertVehicle(db, &oldVehicle)

		v, err := GetVehiclesAfterTime(db, time.Unix(recentDate, 0).Add(time.Minute*-5))

		if err != nil {
			t.Error(err)
		}

		if len(v) != 2 {
			t.Error("Wrong number of items returned")
		}
	})
}

func TestGetVehiclesBeforeTime(t *testing.T) {
	RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		var recentDate int64 = 1460432740083 / 1000
		var oldDate int64 = 1420919252102 / 1000
		var newVehicle = muni.Vehicle{ID: "1234", TimeRecieved: time.Unix(recentDate, 0)}
		var newVehicle2 = muni.Vehicle{ID: "1235", TimeRecieved: time.Unix(recentDate, 0).Add(-time.Minute)}
		var oldVehicle = muni.Vehicle{ID: "1236", TimeRecieved: time.Unix(oldDate, 0)}

		InsertVehicle(db, &newVehicle)
		InsertVehicle(db, &newVehicle2)
		InsertVehicle(db, &oldVehicle)

		v, err := GetVehiclesBeforeTime(db, time.Unix(recentDate, 0).Add(time.Minute*-5))

		if err != nil {
			t.Error(err)
		}

		if len(v) != 1 {
			t.Error("Wrong number of items returned")
		}
	})
}
