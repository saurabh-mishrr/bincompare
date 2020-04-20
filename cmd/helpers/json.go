package helpers

import (
	"encoding/json"
	"io/ioutil"
)

type FileDiffLog struct {
	Signature []string
	Delta     []string
}
type FileLog struct {
	Alive    []string
	Dead     []string
	FilePath FileDiffLog
}
type MainDir struct {
	Project string
	Tree    FileLog
}

func GenerateJson(relativeFilePath []string) {
	fileDiffLog := FileDiffLog{
		Signature: []string{},
		Delta:     []string{},
	}

	fileLog := FileLog{
		Alive:    relativeFilePath,
		Dead:     []string{},
		FilePath: fileDiffLog,
	}

	mainJsonNode := MainDir{
		Project: CurrentDirName,
		Tree:    fileLog,
	}
	var jsonData []byte
	jsonData, err = json.MarshalIndent(mainJsonNode, "", "   ")
	LogError(err)

	_ = ioutil.WriteFile(BinCompareDirName+PathSeparator+"config.json", jsonData, 0644)
}
