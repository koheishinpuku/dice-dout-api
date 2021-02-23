package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	roomController "github.com/koheishinpuku/dice-dout-api/controllers/room"
	userController "github.com/koheishinpuku/dice-dout-api/controllers/user"
)

var port = os.Getenv("PORT")

// Init :
func Init() {
	e := echo.New()

	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	v1 := e.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/:id", userController.Get)
		}
		rooms := v1.Group("/rooms")
		{
			rooms.GET("/:id", roomController.Get)
		}
	}

	fmt.Println("Port:", port)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
