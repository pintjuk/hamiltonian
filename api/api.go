// Package api is the public interface if routemaster service
//
// I like to put types that are parts of the services public api here,
// It makes it simpler to manage integration with down stream services, we may publish is as a go mod
package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Coord represents a point in geographic coordinate system
type Coord struct {
	Long float64
	Lat  float64
}

func (coord Coord) String() string {
	return fmt.Sprintf("%f,%f", coord.Long, coord.Lat)
}

func (coord Coord) MarshalJSON() ([]byte, error) {
	return json.Marshal(coord.String())
}

// Route is information about the route from some source to a destination
type Route struct {
	Destination Coord   `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

// GetRoutesResData Response data to GET routes/
type GetRoutesResData struct {
	Source Coord   `json:"source"`
	Routes []Route `json:"routes"`
}

// MakeCoordFromString parses Coord from a String
func MakeCoordFromString(input string) (Coord, error) {

	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return Coord{}, fmt.Errorf("longlat had %d parts", len(parts))
	}

	long, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return Coord{}, fmt.Errorf("failed to parse longitude: %s", err)
	}

	lat, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return Coord{}, fmt.Errorf("failed to parse latitude: %s", err)
	}

	return Coord{
		Long: long,
		Lat:  lat,
	}, nil
}
