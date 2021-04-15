//+build windows

package main

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	jbDir = "JetBrains"
)

// Clean 清除
func Clean() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	err = clean(homeDir, ConfigDir)
	if err != nil {
		return err
	}
	appData := os.Getenv("APPDATA")
	if appData == "" {
		return errors.New("windows no APPDATA")
	}

	return clean(filepath.Join(appData, jbDir), EmptyDir)
}
