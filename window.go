package main

import (
	"github.com/syyongx/php2go"
	"golang.org/x/sys/windows/registry"
	"log"
	"strings"
)

var (
	SysPath []string
	PhpPath string
)

func Init() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	s, _, _ := k.GetStringValue("Path")
	s = strings.TrimRight(s, ";")
	SysPath = strings.Split(s, ";")
	for _, val := range SysPath {
		if php2go.FileExists(val + "php.exe") {
			PhpPath = val
		}
		SysPath = append(SysPath, val)
	}
}

func GetEnv() {
}

func SetEnv() {

}
