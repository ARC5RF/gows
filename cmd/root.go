/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/ARC5RF/gows/cmd/___"
	_ "github.com/ARC5RF/gows/cmd/mod"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := ___.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gows.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//___.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
