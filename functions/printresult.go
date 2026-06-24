package functions

import (
	"fmt"
)

func PrintResult(result string) error {
	if result == "" {
		return fmt.Errorf("nothing to print")
	}
	fmt.Print(result)
	return nil
}
