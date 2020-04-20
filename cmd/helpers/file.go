package helpers

import (
	"io/ioutil"
	"strings"

	"github.com/h2non/filetype"
)

func IsItImageFile(filepath string) bool {
	fi, err := ioutil.ReadFile(filepath)
	LogError(err)
	return filetype.IsImage(fi)
}

func GetRelativeImageFilePath(filepath *string) string {
	imageRelPath := strings.ReplaceAll(*filepath, CurrentDirAbsPath, ".")
	return imageRelPath
}
