// Package osrm contains client code for OSRM service
package osrm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IntegrationError struct {
	err error
}

func (err IntegrationError) Error() string {
	return fmt.Sprintf("OSRM integration: %s", err.err.Error())
}

// GetRoute function sends a GET request to the OSRM routing API and returns the fetched route
// information.
func GetRoute(from string, to string) (*APIResponse, error) {

	uri := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%s;%s?overview=false", from, to)
	res, err := http.Get(uri)
	if err != nil {
		return nil, IntegrationError{err}
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, IntegrationError{fmt.Errorf("GET %s returned %s", uri, res.Status)}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, IntegrationError{err}
	}

	var route APIResponse
	err = json.Unmarshal(body, &route)
	if err != nil {
		return nil, IntegrationError{err}
	}

	if route.Code != "Ok" {
		return nil, IntegrationError{fmt.Errorf("response code was not ok")}
	}

	if len(route.Routes) == 0 {
		return nil, IntegrationError{fmt.Errorf("expected to recive 1 route, recived %d", len(route.Routes))}
	}

	return &route, nil
}
