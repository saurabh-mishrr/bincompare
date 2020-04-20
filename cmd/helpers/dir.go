package helpers

import (
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
	imageFiles := GetAbsPathContainsImages()

	var relativeFilePath = make([]string, 0)

	for _, imageFilePath := range imageFiles {
		CloneToBinCompare(&imageFilePath)
		CloneToSignature(&imageFilePath)
		CloneToDelta(&imageFilePath)
		CloneToDead(&imageFilePath)
		relativeFilePath = append(relativeFilePath, GetRelativeImageFilePath(&imageFilePath))
		LogError(err)
	}

	GenerateJson(relativeFilePath)

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

func CloneToBinCompare(dirpath *string) {
	s := strings.ReplaceAll(*dirpath, CurrentDirAbsPath, BinCompareDirName+"/"+CurrentDirName)
	MakeRecursiveDirctories(&s)
}

func CloneToSignature(dirpath *string) {
	s := strings.ReplaceAll(*dirpath, CurrentDirAbsPath, BinCompareDirName+"/signatures/"+CurrentDirName)
	MakeRecursiveDirctories(&s)
}

func CloneToDelta(dirpath *string) {
	s := strings.ReplaceAll(*dirpath, CurrentDirAbsPath, BinCompareDirName+"/deltas/"+CurrentDirName)
	MakeRecursiveDirctories(&s)
}

func CloneToDead(dirpath *string) {
	s := strings.ReplaceAll(*dirpath, CurrentDirAbsPath, BinCompareDirName+"/deads/"+CurrentDirName)
	MakeRecursiveDirctories(&s)
}

func GetAbsPathContainsImages() []string {
	//storing all the directory path who contains images
	imageFiles := make([]string, 0)
	err = filepath.Walk(CurrentDirAbsPath, func(path string, info os.FileInfo, err error) error {
		LogError(err)
		if !(IsItDir(path)) && IsItImageFile(path) {
			imageFiles = append(imageFiles, path)
		}
		return nil
	})
	LogError(err)
	return imageFiles
}

func MakeRecursiveDirctories(dirpath *string) {
	a := strings.Split(*dirpath, PathSeparator)
	var subDirLen int = len(a) - 1
	a[subDirLen] = ""
	*dirpath = strings.Join(a, PathSeparator)
	err = os.MkdirAll(*dirpath, os.ModePerm)
	LogError(err)
}
