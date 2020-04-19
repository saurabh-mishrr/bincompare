package helpers

import (
	"os"
	"path/filepath"
)

var FileTypes = [...]string{"jpg", "jpeg", "png", "psd"}

/**
 * Returns bool if given path is directory
 */
func IsItDir(dirpath string) bool {
	fi, err := os.Stat(dirpath)
	if fi.IsDir() {
		return true
	}
	LogError(err)
	return false
}

func GetDirName(dirpath string) string {
	return filepath.FromSlash(dirpath)
}

func GetHomeDirPath() string {
	homeDir, err := os.UserHomeDir()
	LogError(err)
	return homeDir
}
