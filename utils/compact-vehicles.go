package utils

import (
	"strings"
	"time"

	"github.com/k4orta/muni"
	"github.com/k4orta/tunnel-speed-api/models"
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
				Predictable:      true,
			}
			out = append(out, cache[v.ID])
		}

		cache[v.ID].Stats = append(cache[v.ID].Stats, &models.VehicleStats{
			SpeedKmHr:   v.SpeedKmHr,
			Predictable: v.Predictable,
			Heading:     v.Heading,
			Time:        v.TimeReceived.Add(-time.Second * time.Duration(v.SecsSinceReport)),
			DirTag:      v.DirTag,
			Position: models.LatLng{
				Lat: v.Lat,
				Lng: v.Lng,
			},
		})
	}

	// Loop over created MultiVehicle objects and buble up certain stats
	for _, v := range out {
		if !v.Stats[0].Predictable {
			v.Predictable = false
		}

		v.Direction = parseDirection(v.Stats[0].DirTag)
	}

	return out
}

func parseDirection(dirTag string) string {
	if strings.Contains(dirTag, "_I_") {
		return "inbound"
	}
	return "outbound"
}
