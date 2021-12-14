package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAppDataDir() string {
	appDataDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Print(MsgUnableToGetLocalAppDataDir)
	}

	appDataDir = filepath.Join(appDataDir, "jot")

	if _, err := os.Stat(appDataDir); os.IsNotExist(err) {
		os.Mkdir(appDataDir, 0755)
	}

	return appDataDir
}
