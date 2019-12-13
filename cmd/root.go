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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		files := utils.FindAllFiles(args[0])
		for _, file := range files {
			if result := utils.CheckFile(file); result != "" {
				fmt.Println(result)
			}
		}

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
