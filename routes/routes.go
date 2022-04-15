package routes

import (
	"crud-go/controller"
	"crud-go/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/waifu", controller.FetchAllWaifu, middleware.IsAuthenticated)
	e.POST("/waifu", controller.StoreWaifu)
	e.PUT("/waifu", controller.UpdateWaifu)
	e.DELETE("/waifu", controller.DeleteWaifu)

	e.GET("/generate-hash/:password", controller.GenerateHashPassword)
	e.POST("/login", controller.CheckLogin)

	return e
}
