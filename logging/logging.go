package logging

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Joshua152/jot/utils"
)

func GetLogFile() *os.File {
	f, err := os.OpenFile(filepath.Join(utils.GetAppDataDir(), "log.txt"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(utils.MsgUnableToOpenLogFile, err)
	}

	return f
}
