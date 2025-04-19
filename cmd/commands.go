package cmd

import (
	"fmt"

	"github.com/credu/github-activity/cmd/services"
	"github.com/credu/github-activity/cmd/utils"
)

func ShowDefaultMessage() {
	fmt.Println("Provide the GitHub username as an argument when running the CLI.")
	fmt.Println("github-activity <username>")
}

func ShowUserActivityByUsername(username string) {
	userActivities := services.GetUserActivityByUsername(username)

	fmt.Println("Output:")
	for _, activity := range userActivities {
		activityMessage := utils.GetMessageFromActivity(activity)
		fmt.Println("-", activityMessage)
	}
}
