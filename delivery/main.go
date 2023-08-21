// main.go
package main

import (
	_httpDeliver "theapp/delivery/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	_httpDeliver.ConfigureLoanHTTP(e)
	e.Start(":8080")
}
