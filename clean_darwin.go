//+build darwin

package main

import (
	"os"
	"path/filepath"
)

const (
	jbDir = "JetBrains"

	library            = "Library"
	preferences        = "Preferences"
	applicationSupport = "Application Support"
)

// Clean 清除
func Clean() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	err = clean(filepath.Join(homeDir, library, preferences), EmptyDir)
	if err != nil {
		return err
	}
	return clean(filepath.Join(homeDir, library, applicationSupport, jbDir), EmptyDir)
}
