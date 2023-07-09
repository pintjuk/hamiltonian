// Package route package provides the routing logic for the application
//
// this is the "Domain" layer in layered architecture
package route

import (
	"sort"
)

type Cord struct {
	Long float64
	Lat  float64
}

type Route struct {
	Destination Cord
	Duration    float64
	Distance    float64
}

// ByTimeAndDistance is used for sorting routs,
// primarily by Duration and secondarily by Distance
type ByTimeAndDistance []Route

func (a ByTimeAndDistance) Len() int {
	return len(a)
}

func (a ByTimeAndDistance) Less(i, j int) bool {
	if a[i].Duration == a[j].Duration {
		return a[i].Distance < a[j].Distance
	}
	return a[i].Duration < a[j].Duration
}

func (a ByTimeAndDistance) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// GetClosestRouteWithDurationAndDistance calculates routs from source to a list of destinations, using the provided function getDistance
//
// *getDistance*: a caller supplied function that calculates the distance and duration to travel from src to destination
func GetClosestRouteWithDurationAndDistance(
	source Cord,
	destinations []Cord,
	getDistance func(src Cord, dist Cord) (duration float64, distance float64, err error)) []Route {

	var routs []Route
	for _, destination := range destinations {
		duration, distance, err := getDistance(source, destination)

		if err != nil {
			continue
		}

		routs = append(routs,
			Route{
				Destination: destination,
				Duration:    duration,
				Distance:    distance,
			})

	}

	sort.Sort(ByTimeAndDistance(routs))
	return routs
}
