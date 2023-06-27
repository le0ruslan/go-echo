package main

import (
	"github.com/labstack/echo"
	"github.com/le0ruslan/go-echo/cmd/api/handlers"
)

func main() {
	e := echo.New()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
