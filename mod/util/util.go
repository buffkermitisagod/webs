package util

import (
	"errors"
	"os"
)

func Verify_path(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		// does not exist
		return false
	}

	return true
}
