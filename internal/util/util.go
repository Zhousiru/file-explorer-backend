package util

import (
	"os"
	"path"

	"github.com/Zhousiru/file-explorer-backend/internal/config"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		// not exist
		return false
	}

	return true
}

func IsDir(relPath string) bool {
	fullPath := path.Join(config.Get(config.K_ROOT), relPath)

	info, err := os.Stat(fullPath)
	if err != nil {
		return false
	}

	return info.IsDir()
}
