package route

import (
	"errors"
	"testing"
)

type testDistance struct {
	Duration    float64
	Distance    float64
	ShouldError bool
}

func (td testDistance) calculateDistance(_ Cord, _ Cord) (duration float64, distance float64, err error) {
	if td.ShouldError {
		return 0, 0, errors.New("unable to calculate distance")
	}
	return td.Duration, td.Distance, nil
}

func TestGetClosestRouteWithDurationAndDistance(t *testing.T) {

	source := Cord{Long: 1.1, Lat: 2.2}
	destinations := []Cord{
		{Long: 3.3, Lat: 4.4},
		{Long: 5.5, Lat: 6.6},
	}

	t.Run("should return sorted routes based on order given by distance function", func(t *testing.T) {
		distance := testDistance{Duration: 1.0, Distance: 1.1}
		routes := GetClosestRouteWithDurationAndDistance(source, destinations, distance.calculateDistance)

		if len(routes) != len(destinations) {
			t.Fatalf("expected %d routes, got %d", len(destinations), len(routes))
		}

		for i, route := range routes {
			if route.Duration != distance.Duration || route.Distance != distance.Distance || route.Destination != destinations[i] {
				t.Errorf("expected route with duration %f, distance %f and destination %+v, got route with duration %f, distance %f and destination %+v",
					distance.Duration, distance.Distance, destinations[i], route.Duration, route.Distance, route.Destination)
			}
		}
	})

	t.Run("should skip any routes that cause distance to error", func(t *testing.T) {
		erroneousDistance := testDistance{ShouldError: true}
		routes := GetClosestRouteWithDurationAndDistance(source, destinations, erroneousDistance.calculateDistance)

		if len(routes) != 0 {
			t.Fatalf("expected %d routes, got %d", 0, len(routes))
		}
	})

	t.Run("should maintain the order of two routes with identical times and distances based", func(t *testing.T) {
		distance := testDistance{Duration: 1.0, Distance: 1.1}
		routes := GetClosestRouteWithDurationAndDistance(source, destinations, distance.calculateDistance)

		if len(routes) != len(destinations) {
			t.Fatalf("expected %d routes, got %d", len(destinations), len(routes))
		}

		if routes[0].Destination != destinations[0] || routes[1].Destination != destinations[1] {
			t.Errorf("expected first route with destination %+v and second route with destination %+v, got first route with destination %+v and second route with destination %+v",
				destinations[0], destinations[1], routes[0].Destination, routes[1].Destination)
		}

		for _, route := range routes {
			if route.Duration != distance.Duration || route.Distance != distance.Distance {
				t.Errorf("expected route with duration %f, distance %f, got route with duration %f, distance %f",
					distance.Duration, distance.Distance, route.Duration, route.Distance)
			}
		}
	})
}
