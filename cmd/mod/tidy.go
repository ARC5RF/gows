/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package mod

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ARC5RF/gows/cmd/___"
	"github.com/spf13/cobra"
)

// tidyCmd represents the tidy command
var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "add missing and remove unused modules",
	Long:  `add missing and remove unused modules`,
	Run:   TidyRun,
}

func TidyRun(cmd *cobra.Command, args []string) {
	wf, return_point, err := ___.FindWorkFile()
	if err != nil {
		fmt.Println("could not find go.work file...")
		fmt.Println("you are probably not in a go workspace")
		return
	}

	for _, use := range wf.Use {
		if use.DiskPath != "" {
			p := filepath.Join(return_point, filepath.FromSlash(use.DiskPath))
			fmt.Println("go mod tidy", p)
			os.Chdir(p)
			___.Exec("mod", "tidy")
		}
	}

	os.Chdir(return_point)
}

func init() {
	ModCmd.AddCommand(tidyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tidyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tidyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
