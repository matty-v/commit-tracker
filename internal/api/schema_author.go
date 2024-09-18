package api

import (
	"fmt"
)

type Author struct {
	// The Email Of The Author.
	Email	string	`json:"email"`


	// The Name Of The Author.
	Name	string	`json:"name"`
}

// Checks if all of the required fields for Author are set
// and validates all of the constraints for the object.
func (obj *Author) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"name": obj.Name,
		"email": obj.Email,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty{
			return fmt.Errorf("required field '%s' for object 'Author' is empty or unset", field)
		}
	}

	return nil
}

