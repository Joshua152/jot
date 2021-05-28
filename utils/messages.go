package utils

import "fmt"

var (
	ErrInvalidArgument    = fmt.Errorf("invalid argument")
	ErrRequiresAnArgument = fmt.Errorf("requires an argument; add flag -h or --help to see more")

	ErrUnableToRetrieveNotesFile = fmt.Errorf("unable to retrieve notes file")

	ErrUnableToGetExeDir = fmt.Errorf("unable to get exe directory")

	MsgUnableToOpenLogFile = "unable to open logging file; no logging will occur"
)
