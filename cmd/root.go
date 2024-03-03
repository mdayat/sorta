package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sorta",
	Short: "Terminal-based sorting algorithms cheat sheet",
}

func Execute() error {
	return rootCmd.Execute()
}
