package main

import (
	"os"
	"os/exec"
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
