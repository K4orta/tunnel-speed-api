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

	if v.ID != "1545" || v.Route != "N" || v.LeadingVehicleID != "1510" || v.DirTag != "N____I_E30" {
		t.Error("Did not populate vehicle static properties")
	}

	if len(v.Stats) == 0 {
		t.Error("Did not populate stats")
	}

	if v.Stats[0].Position.Lat != 37.77693 || v.Stats[0].Position.Lng != -122.41684 {
		t.Error("Did not populate the LatLng stat")
	}

	if v.Stats[0].Time != time.Unix(1460775707000/1000, 0).Add(-time.Second*44) {
		t.Error("Did not subtract SecsSinceReport from TimeRecieved")
	}
}

func createStubs() []*muni.Vehicle {
	return []*muni.Vehicle{
		&muni.Vehicle{
			ID:               "1545",
			RouteTag:         "N",
			Lat:              37.77693,
			Lng:              -122.41684,
			TimeRecieved:     time.Unix(1460775707000/1000, 0),
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
			TimeRecieved:     time.Unix(1460775767000/1000, 0),
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
			TimeRecieved:     time.Unix(1460775767000/1000, 0),
			LeadingVehicleID: "1510",
			SpeedKmHr:        0,
			DirTag:           "N____I_E30",
			SecsSinceReport:  44,
		},
	}
}
