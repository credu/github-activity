package utils

import (
	"fmt"

	"github.com/credu/github-activity/cmd/models"
)

// Display the fetched activity in the terminal.
//
// # Output:
//   - Pushed 3 commits to credu/developer-roadmap
//   - Opened a new issue in credu/developer-roadmap
//   - Starred credu/developer-roadmap
//   - ...
func GetMessageFromActivity(activity models.Activity) string {
	// Activity types extracted from https://docs.github.com/en/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28
	switch activity.Type {
	case "CommitCommentEvent":
		return fmt.Sprintf("Created a commit comment in %s", activity.Repo.Name)
	case "CreateEvent":
		return fmt.Sprintf("Created a %s in %s", activity.Payload.RefType, activity.Repo.Name)
	case "DeleteEvent":
		return fmt.Sprintf("Deleted %s in %s", activity.Payload.RefType, activity.Repo.Name)
	case "ForkEvent":
		return fmt.Sprintf("Forked the repository %s", activity.Repo.Name)
	case "GollumEvent":
		return fmt.Sprintf("Created or updated a Wiki Page in %s", activity.Repo.Name)
	case "IssueCommentEvent":
		action := CapitalizeFirstLetter(activity.Payload.Action)
		return fmt.Sprintf("%s a issue comment in %s", action, activity.Repo.Name)
	case "IssuesEvent":
		action := CapitalizeFirstLetter(activity.Payload.Action)
		return fmt.Sprintf("%s a new issue in %s", action, activity.Repo.Name)
	case "MemberEvent":
		action := CapitalizeFirstLetter(activity.Payload.Action)
		return fmt.Sprintf("%s a new collaborator in %s", action, activity.Repo.Name)
	case "PublicEvent":
		return fmt.Sprintf("Made public the repository %s", activity.Repo.Name)
	case "PullRequestEvent":
		action := CapitalizeFirstLetter(activity.Payload.Action)
		return fmt.Sprintf("%s the Pull Request #%d of %s", action, activity.Payload.Number, activity.Repo.Name)
	case "PullRequestReviewEvent":
		action := CapitalizeFirstLetter(activity.Payload.Action)
		return fmt.Sprintf("%s a pull request review", action)
	case "PullRequestReviewCommentEvent":
		action := CapitalizeFirstLetter(activity.Payload.Action)
		return fmt.Sprintf("%s a review comment of Pull Request", action)
	case "PullRequestReviewThreadEvent":
		action := CapitalizeFirstLetter(activity.Payload.Action)
		return fmt.Sprintf("%s a Pull Request Thread", action)
	case "PushEvent":
		return fmt.Sprintf("Pushed %d commits to %s", activity.Payload.Size, activity.Repo.Name)
	case "ReleaseEvent":
		action := CapitalizeFirstLetter(activity.Payload.Action)
		return fmt.Sprintf("%s a release in %s", action, activity.Repo.Name)
	case "SponsorshipEvent":
		action := CapitalizeFirstLetter(activity.Payload.Action)
		return fmt.Sprintf("%s a sponsorship in %s", action, activity.Repo.Name)
	case "WatchEvent":
		return fmt.Sprintf("Starred %s", activity.Repo.Name)
	default:
		return fmt.Sprintf("Event not handled %s", activity.Type)
	}
}
