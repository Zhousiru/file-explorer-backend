package util

import "os"

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		// not exist
		return false
	}

	return true
}
