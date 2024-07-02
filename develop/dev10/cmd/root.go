/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dev10/internal/handlers"
	"os"

	"github.com/spf13/cobra"
)

var h = handlers.NewHandler()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "telnet",
	Short: "A simple TCP client",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: h.Telnet,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dev10.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().String("timeout", "10s", "Timeout in seconds")
}
