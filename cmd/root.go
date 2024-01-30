/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/huntermotko/git_a_job/internal/font"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
const logo = `
 ________  ___  _________        ________             ___  ________  ________     
|\   ____\|\  \|\___   ___\     |\   __  \           |\  \|\   __  \|\   __  \    
\ \  \___|\ \  \|___ \  \_|     \ \  \|\  \          \ \  \ \  \|\  \ \  \|\ /_   
 \ \  \  __\ \  \   \ \  \       \ \   __  \       __ \ \  \ \  \\\  \ \   __  \  
  \ \  \|\  \ \  \   \ \  \       \ \  \ \  \     |\  \\_\  \ \  \\\  \ \  \|\  \ 
   \ \_______\ \__\   \ \__\       \ \__\ \__\    \ \________\ \_______\ \_______\
    \|_______|\|__|    \|__|        \|__|\|__|     \|________|\|_______|\|_______|
`

// fmt.Printf("%s\n", logo)
var rootCmd = &cobra.Command{
	Use:   "git_a_job",
	Short: "My tool for finding a job", 
	Long:  font.GetFont() + logo,
	//		Run: func(cmd *cobra.Command, args []string) {
	//	  },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.json_sql.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
