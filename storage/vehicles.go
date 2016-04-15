package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/k4orta/muni"
)

// InsertVehicle adds a vehicle to the database
func InsertVehicle(db *sqlx.DB, v *muni.Vehicle) error {
	rows, err := db.NamedQuery(`
    INSERT INTO vehicles (
      route_tag, vehicle_id, time_recieved, heading,
      dir_tag, lat, lng, leading_vehicle_id, predictable,
      secs_since_report, speed_km_hr
    )
    VALUES (
      :route_tag, :vehicle_id, :time_recieved, :heading,
      :dir_tag, :lat, :lng, :leading_vehicle_id, :predictable,
      :secs_since_report, :speed_km_hr
    );`, v)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

// GetVehiclesAfterTime returns all vehicles logged after limit
func GetVehiclesAfterTime(db *sqlx.DB, limit time.Time) ([]*muni.Vehicle, error) {
	v := []*muni.Vehicle{}
	err := db.Select(&v, `SELECT * FROM vehicles WHERE time_recieved >= $1`, limit)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetVehiclesBeforeTime returns all vehicles logged after limit
func GetVehiclesBeforeTime(db *sqlx.DB, limit time.Time) ([]*muni.Vehicle, error) {
	v := []*muni.Vehicle{}
	err := db.Select(&v, `SELECT * FROM vehicles WHERE time_recieved < $1`, limit)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetVehiclesByID returns all vehicles with a given ID
func GetVehiclesByID(db *sqlx.DB, id string) ([]*muni.Vehicle, error) {
	v := []*muni.Vehicle{}
	err := db.Select(&v, `SELECT * FROM vehicles WHERE vehicle_id=$1`, id)
	if err != nil {
		return nil, err
	}
	return v, nil
}
