package user

import (
	"net/http"

	"github.com/koheishinpuku/dice-dout-api/models"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	userID := c.Param("id")

	user := models.GetUser(userID)

	return c.JSON(http.StatusOK, user)
}
