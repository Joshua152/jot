package logging

import (
	"fmt"
	"os"
	"strings"

	"github.com/Joshua152/jot/utils"
)

func GetLogFile() *os.File {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(utils.ErrUnableToGetExeDir)
	}

	f, err := os.OpenFile(strings.TrimSuffix(path, `\jot.exe`)+`\logging\log.txt`, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(utils.MsgUnableToOpenLogFile, err)
	}

	return f
}
