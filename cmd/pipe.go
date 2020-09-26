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

	"github.com/spf13/cobra"
  "jib/pkg/newPipe"
)

// pipeCmd represents the pipe command
var pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Create a new pipeline",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    // get the flag value, its default value is false
    pipeName, _ := cmd.Flags().GetBool("new")

    if pipeName {
      newPipeline(args)
    } else {
      printHello(args)
    }

	},
}

func printHello(args []string) {
  fmt.Println("Updateing Pipeline")
  for i, _ := range args {
    fmt.Println(args[i])
  }
}

func newPipeline(args []string) {
  //fmt.Println("Creating new pipeline")
  newPipe.Hello(args)
  //for i, _ := range args {
  //  fmt.Println(args[i])
  //}
}

func init() {
	rootCmd.AddCommand(pipeCmd)
  pipeCmd.Flags().BoolP("new", "n", false, "create a new pipeline")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
