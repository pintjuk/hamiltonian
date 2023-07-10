package http_resources

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pintjuk/routemaster/api"
	"github.com/pintjuk/routemaster/src/integrations/osrm"
	"github.com/pintjuk/routemaster/src/route"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// oSRMGetDistance implements the getDistances function in domain model using the OSRM client
func oSRMGetDistance(src route.Coord, dist route.Coord) (duration float64, distance float64, err error) {
	osrmRoute, err := osrm.GetRoute(
		fmt.Sprintf("%f,%f", src.Long, src.Lat),
		fmt.Sprintf("%f,%f", dist.Long, dist.Lat),
	)
	if err != nil {
		log.Printf("Error: %s", err)
		return 0, 0, err
	}
	return osrmRoute.Routes[0].Duration, osrmRoute.Routes[0].Distance, nil
}

// getRoutes is handler for http request GET <service>/routes?src=13.388860,52.517037&dst=13.397634,52.529407
func getRoutes(c echo.Context) error {

	// Validate src param
	src := c.QueryParam("src")

	// I know we technically do not need to parse coords to implement the specification,
	// but doing it allows us to catch formatting errors and provide better feedback to the caller
	srcCordDTO, err := api.MakeCoordFromString(src)
	if src == "" {
		return c.String(http.StatusBadRequest, "Missing src parameters")
	}

	srcCord := route.Coord{
		Long: srcCordDTO.Long,
		Lat:  srcCordDTO.Lat,
	}

	// Validate dist params
	queryValues, _ := url.ParseQuery(c.QueryString())
	dstParams, ok := queryValues["dst"]
	if !ok {
		return c.String(http.StatusBadRequest, "Missing dst parameters")
	}
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("parameter src had invalid format: %s", err))
	}

	var destCords []route.Coord

	for _, dst := range dstParams {
		destCordDTO, err := api.MakeCoordFromString(strings.TrimSpace(dst))

		if err != nil {
			log.Printf("Error: %s", err)
			return c.String(http.StatusBadRequest, fmt.Sprintf("Parameter dst had invalid format: %s", err))
		}

		destCords = append(destCords, route.Coord{
			Long: destCordDTO.Long,
			Lat:  destCordDTO.Lat,
		})
	}

	// Calculate Results
	routs := route.GetClosestRouteWithDurationAndDistance(
		srcCord,
		destCords,
		oSRMGetDistance,
	)

	// Map results to DTO and return

	var routeDTOs []api.Route

	for _, r := range routs {
		routeDTOs = append(routeDTOs, api.Route{
			Destination: api.Coord{
				Long: r.Destination.Long,
				Lat:  r.Destination.Lat,
			},
			Duration: r.Duration,
			Distance: r.Distance,
		})
	}

	data := api.GetRoutesResData{
		Source: srcCordDTO,
		Routes: routeDTOs,
	}

	return c.JSON(http.StatusOK, data)
}
