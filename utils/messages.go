package utils

import "fmt"

var (
	ErrInvalidArgument    = fmt.Errorf("invalid argument")
	ErrRequiresAnArgument = fmt.Errorf("requires an argument; add flag -h or --help to see more")

	ErrUnableToRetrieveNotesFile = fmt.Errorf("unable to retrieve notes file")

	MsgUnableToGetLocalAppDataDir = `unable to open path to C:\users\[user]\AppData\Local`
	MsgUnableToOpenLogFile        = "unable to open logging file; no logging will occur"
)
