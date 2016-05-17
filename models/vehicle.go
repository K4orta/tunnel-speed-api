package models

import "time"

// MultiVehicle is the basic representation of a train car at a given point in time
type MultiVehicle struct {
	ID               string          `json:"id" db:"vehicle_id"`
	Route            string          `json:"route" db:"route_tag"`
	LeadingVehicleID string          `json:"leadingVehicleId"`
	Stats            []*VehicleStats `json:"stats"`
	Predictable      bool            `json:"predictable"`
}

// VehicleStats contains the properties of a vehicle that change over time
type VehicleStats struct {
	Predictable bool      `json:"predictable"`
	Time        time.Time `json:"time"`
	Position    LatLng    `json:"position"`
	Heading     int       `json:"heading"`
	DirTag      string    `json:"dirTag"`
	SpeedKmHr   float32   `json:"speedKmHr"`
}

// LatLng represents a position
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
