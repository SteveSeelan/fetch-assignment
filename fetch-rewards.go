package main

import (
	"fetch-rewards/resources"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello, World! YO")
	e := echo.New()

	receiptResource := resources.NewReceiptResource()
	receiptResource.SetupRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
