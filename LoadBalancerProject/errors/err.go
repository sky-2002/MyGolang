package errors

import (
	"fmt"
	"os"
)

// handleErr prints the error and exits the program
// Note: this is not how one would want to handle errors in production, but
// serves well for demonstration purposes.
func HandleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
