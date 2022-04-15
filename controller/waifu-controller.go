package controller

import (
	"crud-go/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllWaifu(c echo.Context) error {
	result, err := models.FetchAllWaifu()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreWaifu(c echo.Context) error {
	name := c.FormValue("name")
	title := c.FormValue("title")
	full_name := c.FormValue("full_name")

	result, err := models.StoreWaifu(name, title, full_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateWaifu(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	title := c.FormValue("title")
	full_name := c.FormValue("full_name")

	// convert from formvalue string to int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateWaifu(name, title, full_name, conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteWaifu(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteWaifu(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
