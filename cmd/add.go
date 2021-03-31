package cmd

import (
	"github.com/Joshua152/jot/note"
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add one or more notes",
		Run: func(cmd *cobra.Command, args []string) {
			for _, s := range args {
				note.Add(note.New(s))
			}
		},
	}
)
