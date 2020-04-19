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
	"fmt"
	"io/ioutil"
	"log"

	"github.com/h2non/filetype"
	"github.com/spf13/cobra"
)

var sourceFile, originalFile string

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Find difference between two files.",
	Long:  ``,
	Run:   runDiff,
}

func runDiff(cmd *cobra.Command, args []string) {
	for _, x := range args {
		fmt.Println(x)
	}
	imgageOperation()
}

func imgageOperation() {
	buf, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	if filetype.IsImage(buf) {
		fmt.Println("Image file")
	} else {
		fmt.Println("Not Image file")
	}

}

func init() {
	rootCmd.AddCommand(diffCmd)
	diffCmd.Flags().StringVarP(&sourceFile, "source", "s", "", "give the source file path and name")
	diffCmd.Flags().StringVarP(&originalFile, "destination", "d", "", "give the destination file path and name")
	diffCmd.MarkFlagRequired("source")
	diffCmd.MarkFlagRequired("destination")
}
