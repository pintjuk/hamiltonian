package http_resources

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// getHealth http handler for GET <service>/health
func getHealth(c echo.Context) error {
	response := map[string]interface{}{
		"status": "ok",
	}
	return c.JSON(http.StatusOK, response)
}
