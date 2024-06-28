/*
Copyright © 2024 NAME HERE work@prr133f.ru
*/
package cmd

import (
	"dev05/internal/handlers"
	"os"

	"github.com/spf13/cobra"
)

var h = handlers.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep [flags] pattern files...",
	Short: "Simple grep utility",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: h.Grep,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dev05.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntP("after", "A", 0, "Print N lines after")
	rootCmd.Flags().IntP("before", "B", 0, "Print N lines before")
	rootCmd.Flags().IntP("context", "C", 0, "Print N lines before and after")
	rootCmd.Flags().BoolP("count", "c", false, "Count lines")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "Ignore case")
	rootCmd.Flags().BoolP("invert", "v", false, "Invert match")
	rootCmd.Flags().BoolP("fixed", "F", false, "Fixed mode")
	rootCmd.Flags().BoolP("line-num", "n", false, "Line numbers")
}
