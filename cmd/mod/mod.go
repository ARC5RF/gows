/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package mod

import (
	"github.com/ARC5RF/gows/cmd/___"
	"github.com/spf13/cobra"
)

// tidyCmd represents the tidy command
var ModCmd = &cobra.Command{
	Use:   "mod",
	Short: "module maintenance",
	Long: `Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.`,
	// Run: ModRun,
}

func init() {
	___.AddCommand(ModCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tidyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tidyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
