package cmd

import (
	"fmt"
	"os"

	"github.com/Joshua152/jot/note"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jot [note]",
	Short: "Jot is a CLI note jotting interface",
	Long:  `Jot is a CLI program quickly jot down notes and or ideas`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}

		for _, s := range args {
			note.Add(note.New(s))
		}
	},
}

// func isValid(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

// }

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
