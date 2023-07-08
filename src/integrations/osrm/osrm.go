package osrm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetRoute function sends a GET request to the OSRM routing API and returns the fetched route
// information.
func GetRoute() (*APIResponse, error) {
	res, err := http.Get("http://router.project-osrm.org/route/v1/driving/13.388860,52.517037;13.397634,52.529407?overview=false")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var route APIResponse
	err = json.Unmarshal(body, &route)
	if err != nil {
		return nil, err
	}

	if route.Code != "Ok" {
		return nil, fmt.Errorf("OSRM integration: ")
	}

	return &route, nil
}
