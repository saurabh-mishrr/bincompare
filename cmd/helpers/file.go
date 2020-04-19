package helpers

import (
	"io/ioutil"

	"github.com/h2non/filetype"
)

func IsItImageFile(filepath string) bool {
	fi, err := ioutil.ReadFile(filepath)
	LogError(err)
	return filetype.IsImage(fi)
}
