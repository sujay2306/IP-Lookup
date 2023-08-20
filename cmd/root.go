package cmd

import (
	"github.com/spf13/cobra"
)

var (

	rootCmd = &cobra.Command{
		Use:   "IPLookup-CLI",
		Short: "A CLI for tracking IP",
		Long: `IPLookup-cli is a CLI application which gives info about the user provided IP Address.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
