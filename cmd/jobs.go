/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
// var api string

// jobCmd represents the find command
var jobCmd = &cobra.Command{
	Use:   "job [api]",
	Short: "List jobs from a given api",
	Long: `
    `,
	Run: Run,
}


func Run(cmd *cobra.Command, args []string) {
  if len(args) == 0 {
    fmt.Println("Please specify an api")
    fmt.Println("Available apis: hn(hacker news), gh(github), so(stack overflow)")
    return
  } 
  api := args[0]
  switch api {
  case "hn":
    fmt.Println("Hacker news")
  case "gh":
    fmt.Println("Github")
  case "so":
    fmt.Println("Stack overflow")
  default:
    fmt.Println("Unknown api")
  }
}

func init() {
	rootCmd.AddCommand(jobCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobCmd.Flags().StringVar(&api, "api", "", "Chose an api to use")
}
