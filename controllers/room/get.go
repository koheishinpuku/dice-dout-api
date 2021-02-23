package room

import (
	"net/http"

	"github.com/koheishinpuku/dice-dout-api/models"
	"github.com/labstack/echo/v4"
)

// Get :
func Get(c echo.Context) error {
	roomID := c.Param("id")

	room := models.GetRoom(roomID)

	return c.JSON(http.StatusInternalServerError, room)
}
