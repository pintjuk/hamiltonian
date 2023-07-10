package route

import (
	"errors"
	"testing"
)

type expectedResult struct {
	Duration    float64
	Distance    float64
	Destination Coord
}

type testDistanceReturnValue struct {
	Duration    float64
	Distance    float64
	ShouldError bool
}

type testDistancesState map[Coord]testDistanceReturnValue

// fake for calculateDistances
func (state testDistancesState) calculateDistance(src Coord, dest Coord) (duration float64, distance float64, err error) {

	if state[dest].ShouldError {
		return 0, 0, errors.New("unable to calculate distance")
	}
	return state[dest].Duration, state[dest].Distance, nil
}

func TestGetClosestRouteWithDurationAndDistance(t *testing.T) {
	t.Run("should return sorted routes based on order given by distance function", func(t *testing.T) {
		// Arrange
		source := Coord{Long: 1.1, Lat: 2.2}
		destinations := []Coord{
			{Long: 3.3, Lat: 4.4},
			{Long: 5.5, Lat: 6.6},
			{Long: 8.5, Lat: 0.6},
		}
		fakeState := testDistancesState{
			destinations[0]: {Duration: 2.0, Distance: 1.1},
			destinations[1]: {Duration: 1.0, Distance: 1.1},
			destinations[2]: {Duration: 1.0, Distance: 1.0},
		}

		expected := []expectedResult{
			{1.0, 1.0, destinations[2]},
			{1.0, 1.1, destinations[1]},
			{2.0, 1.1, destinations[0]},
		}

		// Act
		routes := GetClosestRouteWithDurationAndDistance(source, destinations, fakeState.calculateDistance)

		// Assert
		if len(routes) != len(expected) {
			t.Fatalf("expected %d routes, got %d", len(expected), len(routes))
		}

		for i, route := range routes {
			if route.Duration != expected[i].Duration || route.Distance != expected[i].Distance || route.Destination != expected[i].Destination {
				t.Errorf("%d: expected route with duration %f, distance %f and destination %+v, got route with duration %f, distance %f and destination %+v",
					i, expected[i].Duration, expected[i].Distance, destinations[i], route.Duration, route.Distance, route.Destination)
			}
		}
	})

	t.Run("should skip any routes that cause distance to error", func(t *testing.T) {
		// Arrange
		source := Coord{Long: 1.1, Lat: 2.2}
		destinations := []Coord{
			{Long: 3.3, Lat: 4.4},
			{Long: 5.5, Lat: 6.6},
			{Long: 8.5, Lat: 0.6},
		}
		fakeState := testDistancesState{
			destinations[0]: {ShouldError: true},
			destinations[1]: {Duration: 1.0, Distance: 1.1},
			destinations[2]: {Duration: 1.0, Distance: 1.0},
		}

		expected := []expectedResult{
			{1.0, 1.0, destinations[2]},
			{1.0, 1.1, destinations[1]},
		}

		// Act
		routes := GetClosestRouteWithDurationAndDistance(source, destinations, fakeState.calculateDistance)

		// Assert
		if len(routes) != len(expected) {
			t.Fatalf("expected %d routes, got %d", len(expected), len(routes))
		}

		for i, route := range routes {
			if route.Duration != expected[i].Duration || route.Distance != expected[i].Distance || route.Destination != expected[i].Destination {
				t.Errorf("%d: expected route with duration %f, distance %f and destination %+v, got route with duration %f, distance %f and destination %+v",
					i, expected[i].Duration, expected[i].Distance, destinations[i], route.Duration, route.Distance, route.Destination)
			}
		}
	})
}

func TestGetClosestRouteWithDurationAndDistanceAsinc(t *testing.T) {

	t.Run("should return sorted routes based on order given by distance function", func(t *testing.T) {
		// Arrange
		source := Coord{Long: 1.1, Lat: 2.2}
		destinations := []Coord{
			{Long: 3.3, Lat: 4.4},
			{Long: 5.5, Lat: 6.6},
			{Long: 8.5, Lat: 0.6},
		}
		fakeState := testDistancesState{
			destinations[0]: {Duration: 2.0, Distance: 1.1},
			destinations[1]: {Duration: 1.0, Distance: 1.1},
			destinations[2]: {Duration: 1.0, Distance: 1.0},
		}

		expected := []expectedResult{
			{1.0, 1.0, destinations[2]},
			{1.0, 1.1, destinations[1]},
			{2.0, 1.1, destinations[0]},
		}

		// Act
		routes := GetClosestRouteWithDurationAndDistanceAsync(source, destinations, fakeState.calculateDistance)

		// Assert
		if len(routes) != len(expected) {
			t.Fatalf("expected %d routes, got %d", len(expected), len(routes))
		}

		for i, route := range routes {
			if route.Duration != expected[i].Duration || route.Distance != expected[i].Distance || route.Destination != expected[i].Destination {
				t.Errorf("%d: expected route with duration %f, distance %f and destination %+v, got route with duration %f, distance %f and destination %+v",
					i, expected[i].Duration, expected[i].Distance, destinations[i], route.Duration, route.Distance, route.Destination)
			}
		}
	})

	t.Run("should skip any routes that cause distance to error", func(t *testing.T) {
		// Arrange
		source := Coord{Long: 1.1, Lat: 2.2}
		destinations := []Coord{
			{Long: 3.3, Lat: 4.4},
			{Long: 5.5, Lat: 6.6},
			{Long: 8.5, Lat: 0.6},
		}
		fakeState := testDistancesState{
			destinations[0]: {ShouldError: true},
			destinations[1]: {Duration: 1.0, Distance: 1.1},
			destinations[2]: {Duration: 1.0, Distance: 1.0},
		}

		expected := []expectedResult{
			{1.0, 1.0, destinations[2]},
			{1.0, 1.1, destinations[1]},
		}

		// Act
		routes := GetClosestRouteWithDurationAndDistanceAsync(source, destinations, fakeState.calculateDistance)

		// Assert
		if len(routes) != len(expected) {
			t.Fatalf("expected %d routes, got %d", len(expected), len(routes))
		}

		for i, route := range routes {
			if route.Duration != expected[i].Duration || route.Distance != expected[i].Distance || route.Destination != expected[i].Destination {
				t.Errorf("%d: expected route with duration %f, distance %f and destination %+v, got route with duration %f, distance %f and destination %+v",
					i, expected[i].Duration, expected[i].Distance, destinations[i], route.Duration, route.Distance, route.Destination)
			}
		}
	})
}
