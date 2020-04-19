package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var BinCompareDirName string = ".bincompare"
var CurrentDirAbsPath string
var CurrentDirName string
var err error
var PathSeparator string

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

/**
 * Assign value to global variable
 */
func init() {
	CurrentDirAbsPath = GetPwd()
	CurrentDirName = GetCurrentDirName(CurrentDirAbsPath)
	PathSeparator = GetPathSeperator()
}

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

/**
 * Returns directory name as per path
 */
func GetDirName(dirpath string) string {
	return filepath.FromSlash(dirpath)
}

/**
 * Returns home directory path
 */
func GetHomeDirPath() string {
	homeDir, err := os.UserHomeDir()
	LogError(err)
	return homeDir
}

/**
 * Returns same output as pwd command in linux
 */
func GetPwd() string {
	d, err := os.Getwd()
	LogError(err)
	return d
}

/**
 * Returns current directory name
 */
func GetCurrentDirName(dirpath string) string {
	return filepath.Base(dirpath)
}

/**
 * Returns directory names which contains image files
 */
func GetDirectoryPathContainsImages() {
	//get current directory to initialize the project

	//storing all the directory path who contains images
	imageFiles := make([]string, 0)
	err = filepath.Walk(CurrentDirAbsPath, func(path string, info os.FileInfo, err error) error {
		LogError(err)
		if !(IsItDir(path)) && IsItImageFile(path) {
			imageFiles = append(imageFiles, path)
		}
		return nil
	})

	var liveFilePath = make([]string, 0)

	for _, imageFilePath := range imageFiles {
		s := strings.ReplaceAll(imageFilePath, CurrentDirAbsPath, BinCompareDirName+"/"+CurrentDirName)
		imageRelPath := strings.ReplaceAll(imageFilePath, CurrentDirAbsPath, ".")
		liveFilePath = append(liveFilePath, imageRelPath)
		a := strings.Split(s, PathSeparator)
		var subDirLen int = len(a) - 1
		a[subDirLen] = ""
		var clonedPath string = strings.Join(a, PathSeparator)
		err = os.MkdirAll(clonedPath, os.ModePerm)
		LogError(err)
	}

	fileDiffLog := FileDiffLog{
		Signature: []string{},
		Delta:     []string{},
	}

	fileLog := FileLog{
		Alive:    liveFilePath,
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
	fmt.Println(string(jsonData))

	LogError(err)
}

/**
 * Returns path separator as per os
 */
func GetPathSeperator() string {
	if IsWindows() {
		return "\\"
	}

	return "/"
}

/**
 * Returns true if current os is windows else false
 */
func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}

	return false
}
