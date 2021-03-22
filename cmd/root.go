package cmd

import (
	"fmt"
	"os"

	"github.com/Joshua152/jot/note"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jot",
	Short: "Jot is a CLI note jotting interface",
	Long:  `Jot is a CLI program quickly jot down notes and or ideas`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Jot!")

		note := note.New("This is a note")
		fmt.Println(note.Note, note.TimeCreated)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
