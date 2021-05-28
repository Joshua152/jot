package cmd

import (
	"fmt"
	"log"

	"github.com/Joshua152/jot/note"
	"github.com/Joshua152/jot/utils"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the list of your notes",
	Run: func(cmd *cobra.Command, args []string) {
		notes, err := note.GetNotes()
		if err != nil {
			fmt.Println(utils.ErrUnableToRetrieveNotesFile, err)
			log.Fatal(err)
		}

		if len(notes) == 0 {
			fmt.Println("No notes. Add notes using 'jot add'.")
		}

		for i, n := range notes {
			fmt.Printf("%d) %s", i+1, n.Note)

			if Verbose {
				fmt.Print("\n" + n.TimeCreated)
			}

			fmt.Println()
		}
	},
}
