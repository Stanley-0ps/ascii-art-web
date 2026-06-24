package functions

import (
	"fmt"
)

func ValidateArgs(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("oops! this requires 2 input %d", len(args))
	}
	return nil
}
