package util

import (
	"os"
	"path/filepath"
)

func GetCurrentDirName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Base(dir), nil
}
