package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/syyongx/php2go"
	"io/ioutil"
	"os"
	"strings"
)

var PhpMap map[string]string

const (
	PvmVersion = "0.0.1"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
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
	case "in-path":
		addSelfPath()
	default:
		help()
	}
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
			color.Red("当前未有 php 变量环境")
			return nil
		}

		// 获取 php 版本信息
		currentPhpVersion, err := WinExec("php", "-r", "echo PHP_VERSION;")
		if err != nil {
			return nil
		}

		err = php2go.FilePutContents("./php.txt", currentPhpVersion+" "+phpPath, 0777)
		if err != nil {
			color.Red(err.Error())
			return nil
		}
	}

	raw, err := ioutil.ReadFile("./php.txt")
	if err != nil {
		color.Red("php.txt can't read")
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
		color.Red("本机未有 php 变量环境")
		return
	}

	phpMap := getPhpMap()
	if phpMap == nil {
		color.Red("本机未有 php 变量环境")
	}

	var newPhpString string
	for key, item := range phpMap {
		if !FileExists(item + "\\php.exe") {
			continue
		}

		console := key + " " + item
		if item == phpPath {
			color.Blue("--> " + console)
		} else {
			fmt.Println("    " + console)
		}
		newPhpString += console + "\r\n"
	}

	// 将新的 phpMap 写入 php.txt 中
	err := php2go.FilePutContents("./php.txt", newPhpString, 0777)
	if err != nil {
		color.Red(err.Error())
	}
}

func addPhp(pPath string) {
	phpExePath := pPath + "/php.exe"
	if !FileExists(phpExePath) {
		color.Red("file not existed")
		return
	}
	// 获取 php 版本信息
	phpVersion, err := WinExec(phpExePath, "-r", "echo PHP_VERSION;")
	if err != nil {
		color.Red(err.Error())
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
		color.Red(err.Error())
	}
	PhpMap = nil
}

// usePhp
// todo check this version is existed
func usePhp(key string) {
	phpMap := getPhpMap()
	path := phpMap[key]
	if path == "" {
		color.Red(key + " not existed")
		return
	}

	if !FileExists(path + "\\php.exe") {
		color.Red(path + "\\php.exe file not existed")
		return
	}

	if !SetEnv(path) {
		color.Red("Failed")
		return
	}

	color.Green("use " + key + " success")
	RefreshEnv()
}

func addSelfPath() {
	// todo 判断是否已经加入
	currentPath := CurrentPath()
	// 将当前目录加入到运行环境中
	err := addToEnv(currentPath)
	if err != nil {
		color.Red(err.Error())
	}
	RefreshEnv()
}

func help() {
	color.Green("\nRunning version v" + PvmVersion)
	color.Yellow("\nUsage:")
	fmt.Println("  command [arguments]")
	color.Yellow("\nCommands:")
	fmt.Println("  add <php_path>		: Add path of new php install dir.")
	fmt.Println("  ls				: List all php versions.")
	fmt.Println("  use <version>			: Switch to use the specified version.")
	fmt.Println("  in-path			: Add pvm to Environment Variables.")
	//fmt.Println("  current                  : Display active version.")
	//fmt.Println("  uninstall <version>      : The version must be a specific version.")
	//fmt.Println("  update                   : Automatically update pvm to the latest version.")
	fmt.Println(" ")
}
