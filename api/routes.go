package api

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type LongLat struct {
	Long float64
	Lat  float64
}

func (cord LongLat) MarshalJSON() ([]byte, error) {
	log.Print(fmt.Sprintf("%f,%f", cord.Long, cord.Lat))
	return json.Marshal(fmt.Sprintf("%f,%f", cord.Long, cord.Lat))
}

type Route struct {
	Destination LongLat `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

type Data struct {
	Source LongLat `json:"source"`
	Routes []Route `json:"routes"`
}

func MakeLongLatFromString(input string) (LongLat, error) {

	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return LongLat{}, fmt.Errorf("longlat had %d parts", len(parts))
	}

	long, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return LongLat{}, fmt.Errorf("failed to parse longitude: %s", err)
	}

	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return LongLat{}, fmt.Errorf("failed to parse latitude: %s", err)
	}

	return LongLat{
		Long: long,
		Lat:  lat,
	}, nil
}
