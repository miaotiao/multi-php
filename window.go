package main

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"os/exec"
	"strings"
)

type PowerShell struct {
	PowerShell string
}

func New() *PowerShell {
	ps, err := exec.LookPath("powershell.exe")
	if err != nil {
		panic(err)
		return nil
	}
	return &PowerShell{
		PowerShell: ps,
	}
}

func (ps *PowerShell) exec(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(ps.PowerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

func RefreshEnv() {
	fmt.Println(`$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")`)
	fmt.Println("copy,paste,run the above code in this command prompt")
}

func Registry() (phpPath string, sysMap []string) {

	if phpPath != "" {
		return phpPath, sysMap
	}

	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	s, _, _ := k.GetStringValue("Path")
	s = strings.TrimRight(s, ";")
	sysMap = strings.Split(s, ";")
	for _, val := range sysMap {

		if FileExists(val + "\\php.exe") {
			phpPath = val
		}
		//SysPath = append(SysPath, val)
	}
	return phpPath, sysMap
}

// SetEnv 设置环境变量
func SetEnv(newPhpPath string) bool {
	if newPhpPath == "" || !FileExists(newPhpPath) {
		return false
	}

	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	s, _, _ := k.GetStringValue("Path")
	s = strings.TrimRight(s, ";")
	sMap := strings.Split(s, ";")

	var newSString string
	for _, sPath := range sMap {

		if FileExists(sPath + "\\php.exe") {
			continue
		}
		newSString += sPath + ";"
	}
	newSString += newPhpPath + ";"

	err = k.SetStringValue("Path", newSString)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

// addToEnv add var to env
func addToEnv(path string) error {
	if path == "" || !FileExists(path) {
		return errors.New("path is incorrect")
	}

	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	s, _, _ := k.GetStringValue("Path")
	s = strings.TrimRight(s, ";")
	sMap := strings.Split(s, ";")

	var newSString string
	for _, sPath := range sMap {

		if sPath == path {
			continue
		}
		newSString += sPath + ";"
	}
	newSString += path + ";"

	err = k.SetStringValue("Path", newSString)
	if err != nil {
		return err
	}
	return nil
}

// delFromEnv remove var from env
func delFromEnv(path string) error {
	if path == "" || !FileExists(path) {
		return errors.New("path is incorrect")
	}

	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	s, _, _ := k.GetStringValue("Path")
	s = strings.TrimRight(s, ";")
	sMap := strings.Split(s, ";")

	var newSString string
	for _, sPath := range sMap {

		if sPath == path {
			continue
		}
		newSString += sPath + ";"
	}

	err = k.SetStringValue("Path", newSString)
	if err != nil {
		return err
	}
	return nil
}
