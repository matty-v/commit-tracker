package api

import (
	"fmt"
)

type Commit struct {
	Author	Author	`json:"author"`

	// The Unique Identifier For The Commit.
	Id	string	`json:"id"`


	// The Commit Message.
	Message	string	`json:"message"`


	// The Date And Time When The Commit Was Made.
	Timestamp	string	`json:"timestamp"`
}

// Checks if all of the required fields for Commit are set
// and validates all of the constraints for the object.
func (obj *Commit) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"id": obj.Id,
		"message": obj.Message,
		"timestamp": obj.Timestamp,
		"author": obj.Author,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty{
			return fmt.Errorf("required field '%s' for object 'Commit' is empty or unset", field)
		}
	}

	return nil
}

