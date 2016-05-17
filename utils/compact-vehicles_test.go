package utils

import (
	"testing"
	"time"

	"github.com/k4orta/muni"
)

func TestCompactVehicles(t *testing.T) {
	vehicles := CompactVehicles(createStubs())

	if len(vehicles) != 2 {
		t.Error("Did not combine vehicles with same ID")
	}

	v := vehicles[0]

	if v.ID != "1545" || v.Route != "N" || v.LeadingVehicleID != "1510" {
		t.Error("Failed to populate vehicle static properties")
	}

	if len(v.Stats) == 0 {
		t.Error("Failed to populate stats")
	}

	if len(v.Stats) != 2 {
		t.Error("Failed to populate all stats")
	}

	stats := v.Stats[0]

	if stats.Position.Lat != 37.77693 || stats.Position.Lng != -122.41684 {
		t.Error("Failed to populate the LatLng stat")
	}

	if stats.Time != time.Unix(1460775707000/1000, 0).Add(-time.Second*44) {
		t.Error("Failed to subtract SecsSinceReport from TimeReceived")
	}
}

func createStubs() []*muni.Vehicle {
	return []*muni.Vehicle{
		&muni.Vehicle{
			ID:               "1545",
			RouteTag:         "N",
			Lat:              37.77693,
			Lng:              -122.41684,
			TimeReceived:     time.Unix(1460775707000/1000, 0),
			LeadingVehicleID: "1510",
			SpeedKmHr:        66,
			DirTag:           "N____I_E30",
			SecsSinceReport:  44,
		},
		&muni.Vehicle{
			ID:               "1545",
			RouteTag:         "N",
			Lat:              37.78032,
			Lng:              -122.41257,
			TimeReceived:     time.Unix(1460775767000/1000, 0),
			LeadingVehicleID: "1510",
			SpeedKmHr:        0,
			DirTag:           "N____I_E30",
			SecsSinceReport:  44,
		},
		&muni.Vehicle{
			ID:               "1111",
			RouteTag:         "N",
			Lat:              37.78032,
			Lng:              -122.41257,
			TimeReceived:     time.Unix(1460775767000/1000, 0),
			LeadingVehicleID: "1510",
			SpeedKmHr:        0,
			DirTag:           "N____I_E30",
			SecsSinceReport:  44,
		},
	}
}
