/*
Copyright © 2019 Zhecheng Li zhechel1@uci.edu

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
	"os"

	"github.com/lzhecheng/comments-collector/utils"
	"github.com/spf13/cobra"
)

// rootCmd is the root command.
var rootCmd = &cobra.Command{
	Use:   "cmtcltor",
	Short: "Collect comments",
	Long:  `Comments-collector is a tool to collect comments in code.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		files := utils.FindAllFiles(args[0])
		fmt.Println("===== Find all files finished =====")

		var results []string
		total := float64(len(files))
		m25 := false
		m50 := false
		m75 := false
		for i, file := range files {
			progress := float64(i) / total
			if progress > 0.25 && !m25 {
				fmt.Println("===== 25% =====")
				m25 = true
			}
			if progress > 0.5 && !m50 {
				fmt.Println("===== 50% =====")
				m50 = true
			}
			if progress > 0.75 && !m75 {
				fmt.Println("===== 75% =====")
				m75 = true
			}
			
			if result := utils.CheckFile(file); result != "" {
				results = append(results, result)
			}
		}
		utils.WriteToOutput(args[1], results)

		fmt.Println("===== Collecting finished =====")
	},
}

// Execute executes all commands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
