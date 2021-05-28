package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Joshua152/jot/note"
	"github.com/Joshua152/jot/utils"
	"github.com/spf13/cobra"
)

var (
	all bool

	removeCmd = &cobra.Command{
		Use:     "remove",
		Short:   "Removes the note(s) corresponding to the specified number or range",
		Long:    `Removes the note(s) at the number or range given (ex. 10, 1:20, :20, 20:) (inclusive)`,
		Aliases: []string{"rem", "del"},
		Run: func(cmd *cobra.Command, args []string) {
			notes, err := note.GetNotes()
			if err != nil {
				log.Fatal(err)
				fmt.Println(utils.ErrUnableToRetrieveNotesFile)
			}

			if all {
				remove(1, len(notes))
				return
			}

			if len(args) == 0 {
				fmt.Println(utils.ErrRequiresAnArgument)
				return
			}

			for _, str := range args {
				split := strings.Split(str, "")
				start := 0
				end := 0

				if len(split) == 1 {
					var err error

					start, err = strconv.Atoi(split[0])
					checkErrInvalidRange(err)
					end, err = strconv.Atoi(split[0])
					checkErrInvalidRange(err)
				} else if len(split) == 2 {
					var err error

					if split[0] == ":" {
						start = 0
						end, err = strconv.Atoi(split[1])
					} else {
						start, err = strconv.Atoi(split[0])
						end = len(notes)
					}

					checkErrInvalidRange(err)
				} else if len(split) == 3 {
					var err error

					start, err = strconv.Atoi(split[0])
					checkErrInvalidRange(err)
					end, err = strconv.Atoi(split[2])
					checkErrInvalidRange(err)
				} else {
					fmt.Println(utils.ErrInvalidArgument)
				}

				remove(start, end)
			}
		},
	}
)

func init() {
	removeCmd.Flags().BoolVarP(&all, "all", "a", false, "Removes all notes")
}

func remove(start, end int) {
	for i := start; i <= end; i++ {
		note.Remove(start - 1)
		fmt.Printf("Removed note #%d\n", i)
	}
}

func checkErrInvalidRange(err error) {
	if err != nil {
		fmt.Println(utils.ErrInvalidArgument)
		os.Exit(1)
	}
}
