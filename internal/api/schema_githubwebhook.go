package api

import (
	"fmt"
)

type GitHubWebhook struct {
	Commits	[]Commit	`json:"commits"`

	// The Git Reference For The Commit.
	Ref	string	`json:"ref"`
}

// Checks if all of the required fields for GitHubWebhook are set
// and validates all of the constraints for the object.
func (obj *GitHubWebhook) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"ref": obj.Ref,
		"commits": obj.Commits,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty{
			return fmt.Errorf("required field '%s' for object 'GitHubWebhook' is empty or unset", field)
		}
	}

	return nil
}

