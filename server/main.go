package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/yarncraft/koios/recommender"  
  	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "golang.org/x/crypto/acme/autocert"
)

var err error
var rr *recommender.Redrec

func main() {

	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

	rr, err = recommender.New(os.Getenv("REDIS_URL"))
	checkErrorAndExit(err)

	// Initiate HTTP Server
  	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
  		return key == os.Getenv("API_SECRET"), nil
	}))

	// Routes
	e.GET("/api/rate", RateHandler)
	e.GET("/api/recommend", RecommendationHandler)
	e.GET("/api/update", UpdateHandler)

	// Start server
	// e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	// e.Logger.Fatal(e.StartAutoTLS(":443"))
	e.Logger.Fatal(e.Start(":1323"))
}


func checkErrorAndExit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		rr.CloseConn()
		os.Exit(1)
	}
}