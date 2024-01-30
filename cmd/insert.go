/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


func RunInsert(cmd *cobra.Command, args []string) {
  fmt.Printf("Inserting data from %s\n", args[0])
}

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long: ``,
  Args: cobra.MinimumNArgs(1),
	Run: RunInsert,

}

func init() {
	rootCmd.AddCommand(insertCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
