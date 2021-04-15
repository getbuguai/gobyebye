//+build linux

package main

import (
	"os"
	"path/filepath"
)

const (
	jbDir = "config"
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

	return clean(filepath.Join(homeDir, jbDir), EmptyDir)
}
