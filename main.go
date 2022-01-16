package main

import (
	"fmt"
	"github.com/syyongx/php2go"
	"io/ioutil"
	"os"
	"strings"
)

var PhpMap map[string]string

func main() {
	args := os.Args
	if len(args) < 2 {
		// todo show usage == help
		fmt.Println("pvm add|use|ls [args]")
		return
	}

	switch args[1] {
	case "use":
		usePhp(args[2])
	case "add":
		addPhp(args[2])
		lsPhp()
	case "ls":
		lsPhp()
	}
	//execPhp()

}

// getPhpMap 获取 php 的版本信息
func getPhpMap() map[string]string {
	if PhpMap != nil {
		return PhpMap
	}

	PhpMap = make(map[string]string)
	//	获取文件内容
	if !FileExists("./php.txt") {
		err := php2go.FilePutContents("./php.txt", "", 0777)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		phpPath, _ := Registry()
		if phpPath == "" {
			fmt.Println("当前未有 php 变量环境")
			return nil
		}

		// 获取 php 版本信息
		currentPhpVersion, err := WinExec("php", "-r", "echo PHP_VERSION;")
		if err != nil {
			return nil
		}

		err = php2go.FilePutContents("./php.txt", currentPhpVersion+" "+phpPath, 0777)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}

	raw, err := ioutil.ReadFile("./php.txt")
	if err != nil {
		fmt.Println("php.json can't read")
		return nil
	}

	// 为 win && *unix 换行考虑
	rows := strings.Split(strings.ReplaceAll(string(raw), "\r\n", "\n"), "\n")
	for _, item := range rows {
		if item != "" {
			col := strings.Fields(item)
			PhpMap[col[0]] = col[1]
		}
	}
	return PhpMap
}

func lsPhp() {
	//	获取当前 php 环境
	phpPath, _ := Registry()

	if phpPath == "" {
		fmt.Println("本机未有 php 变量环境")
		return
	}

	phpMap := getPhpMap()
	if phpMap == nil {
		fmt.Println("本机未有 php 变量环境")
	}

	var newPhpString string
	for key, item := range phpMap {
		if !FileExists(item + "\\php.exe") {
			continue
		}

		console := key + " " + item
		if item == phpPath {
			fmt.Println("--> " + console)
		} else {
			fmt.Println("    " + console)
		}
		newPhpString += console + "\r\n"
	}

	// 将新的 phpMap 写入 php.txt 中
	err := php2go.FilePutContents("./php.txt", newPhpString, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func addPhp(pPath string) {
	phpExePath := pPath + "/php.exe"
	if !FileExists(phpExePath) {
		fmt.Println("file not existed")
		return
	}
	// 获取 php 版本信息
	phpVersion, err := WinExec(phpExePath, "-r", "echo PHP_VERSION;")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	phpMap := getPhpMap()
	var newPhpString string
	if phpMap != nil {
		for key, item := range PhpMap {
			if pPath == item {
				continue
			}
			phpStringRow := key + " " + item
			newPhpString += phpStringRow + "\r\n"
		}
	}
	newPhpString += phpVersion + " " + pPath + "\r\n"
	err = php2go.FilePutContents("./php.txt", newPhpString, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
	PhpMap = nil
}

// usePhp
// todo check this version is existed
func usePhp(key string) {
	phpMap := getPhpMap()
	path := phpMap[key]
	if path == "" {
		fmt.Println(key + " not existed")
		return
	}

	if !FileExists(path + "\\php.exe") {
		fmt.Println(path + "\\php.exe file not existed")
		return
	}

	if !SetEnv(path) {
		fmt.Println("Failed")
		return
	}

	fmt.Println("use " + key + " success")
	RefreshEnv()
}
