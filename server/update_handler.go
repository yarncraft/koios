package main

import (
	"fmt"
	"net/http"
  	"github.com/labstack/echo/v4"
)

func UpdateHandler(c echo.Context) error {
	update(10)
	return c.String(http.StatusOK, "OK")
}

func update(max int) {
	fmt.Printf("Updating DB\n")
	err := rr.BatchUpdateSimilarUsers(max)
	checkErrorAndExit(err)
	fmt.Printf("DB Updated!\n")
}