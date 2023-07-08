package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pintjuk/routemaster/api"
	osrm "github.com/pintjuk/routemaster/src/integrations/osrm"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func getHealth(c echo.Context) error {
	response := map[string]interface{}{
		"status": "ok",
	}
	return c.JSON(http.StatusOK, response)
}

func getRouts(c echo.Context) error {

	src := c.QueryParam("src")
	queryValues, _ := url.ParseQuery(c.QueryString())
	dsts, ok := queryValues["dst"]
	if !ok {
		return c.String(http.StatusBadRequest, "Missing dst parameters")
	}

	if src == "" {
		return c.String(http.StatusBadRequest, "Missing src parameters")
	}

	srcLongLat, err := api.MakeLongLatFromString(src)

	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("parameter src had invalid format: %s", err))
	}

	var dstLongLat []api.LongLat

	for _, dst := range dsts {
		r, err := api.MakeLongLatFromString(strings.TrimSpace(dst))

		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Parameter dst had invalid format: %s", err))
		}
		log.Printf("%v", r)

		dstLongLat = append(dstLongLat, r)
	}

	data := api.Data{
		Source: srcLongLat,
		Routes: []api.Route{
			{
				Destination: dstLongLat[0],
				Duration:    465.2,
				Distance:    1879.4,
			},
			{
				Destination: dstLongLat[0],
				Duration:    712.6,
				Distance:    4123,
			},
		},
	}

	return c.JSON(http.StatusOK, data)
}

func InitRoutes(e *echo.Echo) {
	e.GET("/routes", getRouts)
	e.GET("/health", getHealth)
}

//func main() {
//	e := echo.New()
//	config := newConfig()
//	// Middleware
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//
//	// Routes
//
//	InitRoutes(e) // Initialize API routes
//
//	// Start server
//	port := fmt.Sprintf(":%s", config.port)
//	log.Printf("Server started on port %s\n", port)
//
//	log.Fatal(e.Start(port))
//
//}

func main() {
	route, err := osrm.GetRoute()

	if err != nil {

		log.Fatal(err)
	}
	v, _ := json.Marshal(route)

	fmt.Println(string(v))
}
