package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func WinExec(name string, arg ...string) (string, error) {
	c := exec.Command(name, arg...)
	output, err := c.CombinedOutput()
	return string(output), err
}

// InstallPath get install path
func InstallPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
