package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Joshua152/jot/logging"
	"github.com/spf13/cobra"
)

var (
	Verbose bool

	rootCmd = &cobra.Command{
		Use:   "jot",
		Short: "Jot is a CLI note jotting interface",
		Long:  `Jot is a CLI program quickly jot down notes and or ideas`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Output all info")

	// close file?
	log.SetOutput(logging.GetLogFile())
	log.SetPrefix("cmd: ")
	log.SetFlags(log.Lshortfile)

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(removeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
