package validators

import (
	"fmt"

	"github.com/Alceaa/SkyPhone/models"
)

func ValidateUser(user models.User) error {
	if user.Username == "" || user.Password == "" {
		return fmt.Errorf("username and password cannot be empty")
	}

	return nil
}
