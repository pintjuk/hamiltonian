// Package api is the public interface if routemaster service
package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Cord represents a point in geographic coordinate system
type Cord struct {
	Long float64
	Lat  float64
}

func (cord Cord) String() string {
	return fmt.Sprintf("%f,%f", cord.Long, cord.Lat)
}

func (cord Cord) MarshalJSON() ([]byte, error) {
	return json.Marshal(cord.String())
}

// Route is information about the rout from some source to a destination
type Route struct {
	Destination Cord    `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

// GetRoutesResData Response data to GET routes/
type GetRoutesResData struct {
	Source Cord    `json:"source"`
	Routes []Route `json:"routes"`
}

// MakeLongLatFromString parses Cord from a String
func MakeLongLatFromString(input string) (Cord, error) {

	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return Cord{}, fmt.Errorf("longlat had %d parts", len(parts))
	}

	long, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return Cord{}, fmt.Errorf("failed to parse longitude: %s", err)
	}

	lat, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return Cord{}, fmt.Errorf("failed to parse latitude: %s", err)
	}

	return Cord{
		Long: long,
		Lat:  lat,
	}, nil
}
