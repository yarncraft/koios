package main

import (
	"fmt"
	"net/http"
	"strconv"
  	"github.com/labstack/echo/v4"
)

func RateHandler(c echo.Context) error {
	user := c.QueryParam("user")
	item := c.QueryParam("item")
	rating_string := c.QueryParam("rating")

	rating, _ := strconv.ParseFloat(rating_string, 64)

	rate(user, item, rating)

	return c.String(http.StatusOK, "Rating submitted")
}


func rate(user string, item string, score float64) {
	fmt.Printf("User %s ranked item %s with %.2f\n", user, item, score)
	err := rr.Rate(item, user, score)
	checkErrorAndExit(err)
}