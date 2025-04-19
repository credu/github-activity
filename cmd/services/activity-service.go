package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/credu/github-activity/cmd/models"
)

func GetUserActivityByUsername(username string) []models.Activity {
	// Make request to https://api.github.com/users/<username>/events
	res, err := http.Get("https://api.github.com/users/" + username + "/events")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		fmt.Println("The user @" + username + " not found")
		os.Exit(1)
	}

	var activities []models.Activity
	json.NewDecoder(res.Body).Decode(&activities)

	return activities
}
