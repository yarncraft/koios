package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"encoding/json"
)

func RecommendationHandler(c echo.Context) error {
	user := c.QueryParam("user")
	recommendations := string(suggest(user, 100))
	return c.String(http.StatusOK, fmt.Sprintf("%v", recommendations))
}

func suggest(user string, max int) []byte {
	fmt.Printf("Getting %d results for user %s\n", max, user)
	rr.UpdateSuggestedItems(user, max)
	s, err := rr.GetUserSuggestions(user, max)
	checkErrorAndExit(err)

	suggestions := make(map[string]interface{})

	for idx := 1; idx < len(s); idx += 2  {
		suggestions[s[idx-1]] = s[idx]
	}

	sugJSON, err := json.Marshal(suggestions)   
    if err != nil {
        fmt.Println(err.Error())
    }
	return sugJSON
}