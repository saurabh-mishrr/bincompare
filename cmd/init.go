/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bincompare/cmd/helpers"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var currentDir string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialisation of project.",
	Long:  `Initialisation of project.`,
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {

	fmt.Println("Initializing project.")
	currentDir, err := os.Getwd()

	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		helpers.LogError(err)

		if helpers.IsItDir(path) {
			fmt.Println(helpers.GetDirName(path))
		} else {
			if helpers.IsItImageFile(path) {
				fmt.Println(path)
			}
		}

		return nil
	})

	helpers.LogError(err)
	// os.Mkdir("./.bincompare", os.ModePerm)
}

func init() {
	rootCmd.AddCommand(initCmd)
}
