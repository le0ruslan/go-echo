package handlers

import (
	"github.com/labstack/echo"
	"github.com/le0ruslan/go-echo/cmd/api/service"
	"net/http"
	"strconv"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, "Unadle to process Data")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process Data")
	}

	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process Data")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}
