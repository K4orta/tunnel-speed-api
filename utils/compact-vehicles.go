package utils

import (
	"time"

	"github.com/k4orta/muni"
	"github.com/k4orta/tunnel-watch/models"
)

// CompactVehicles transforms an array of Vehicles into MultiVehicle and merges the ones with the same ID
func CompactVehicles(vehicles []*muni.Vehicle) []*models.MultiVehicle {
	cache := map[string]*models.MultiVehicle{}
	out := []*models.MultiVehicle{}
	for _, v := range vehicles {
		if _, exists := cache[v.ID]; !exists {
			cache[v.ID] = &models.MultiVehicle{
				ID:               v.ID,
				Route:            v.RouteTag,
				LeadingVehicleID: v.LeadingVehicleID,
			}
			out = append(out, cache[v.ID])
		}
		cache[v.ID].Stats = append(cache[v.ID].Stats, &models.VehicleStats{
			SpeedKmHr:   v.SpeedKmHr,
			Predictable: v.Predictable,
			Heading:     v.Heading,
			Time:        v.TimeRecieved.Add(-time.Second * time.Duration(v.SecsSinceReport)),
			DirTag:      v.DirTag,
			Position: models.LatLng{
				Lat: v.Lat,
				Lng: v.Lng,
			},
		})
	}
	return out
}
